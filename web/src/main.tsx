import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { RouterProvider, createRouter } from '@tanstack/react-router'
import {QueryCache, QueryClient, QueryClientProvider} from '@tanstack/react-query'
import {AppLoading} from "@/components/app-loading";
import {AxiosError} from "axios";
import {HTTP_STATUS_CODE} from "@/lib/api";
import {useAuthStore} from "@/states/auth-state";

// Import the styles
import './index.css'

// Import the generated route tree
import { routeTree } from './routeTree.gen'

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: (failureCount, error) => {
        if (import.meta.env.DEV) console.log({ failureCount, error })
        if (failureCount >= 0 && import.meta.env.DEV) return false
        if (failureCount > 3 && import.meta.env.PROD) return false
        // noinspection SuspiciousTypeOfGuard
        return !(
          error instanceof AxiosError &&
          [
            HTTP_STATUS_CODE.INTERNAL_SERVER_ERROR,
            HTTP_STATUS_CODE.FORBIDDEN
          ].includes(error.response?.status ?? 0)
        )
      },
      refetchOnWindowFocus: import.meta.env.PROD,
      staleTime: 10 * 1000, // 10s
    },
    mutations: {
      onError: (error) => {
        // noinspection SuspiciousTypeOfGuard
        if (error instanceof AxiosError) {
          if (error.response?.status === HTTP_STATUS_CODE.UNAUTHORIZED) {
            useAuthStore.getState().auth.reset()
            const redirect = `${router.history.location.href}`
            router.navigate({ to: '/login', search: { redirect } })
          }
        }
      },
    },
  },
  queryCache: new QueryCache({
    onError: (error) => {
      // noinspection SuspiciousTypeOfGuard
      if (error instanceof AxiosError) {
        if (error.response?.status === HTTP_STATUS_CODE.UNAUTHORIZED) {
          useAuthStore.getState().auth.reset()
          const redirect = `${router.history.location.href}`
          router.navigate({ to: '/login', search: { redirect } })
        }
      }
    },
  }),
})

// Create a new router instance
const router = createRouter({
  context: { queryClient },
  routeTree,
  defaultPreload: 'intent',
  defaultPendingMinMs: 0,
  defaultPendingComponent: AppLoading,
})

// Register the router instance for type safety
declare module '@tanstack/react-router' {
  interface Register {
    router: typeof router
  }
}

// Render the app
const rootElement = document.getElementById('root')!
if (!rootElement.innerHTML) {
  const root = createRoot(rootElement)
  root.render(
    <StrictMode>
      <QueryClientProvider client={queryClient}>
        <RouterProvider router={router} />
      </QueryClientProvider>
    </StrictMode>,
  )
}

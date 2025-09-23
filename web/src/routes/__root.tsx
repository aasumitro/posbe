import { createRootRouteWithContext } from '@tanstack/react-router'
import MainLayout from '@/layouts/main'
import type {QueryClient} from "@tanstack/react-query";

export const Route = createRootRouteWithContext<{
  queryClient: QueryClient
}>()({
  component: MainLayout,
})
import { createFileRoute } from '@tanstack/react-router'
import AuthenticatedLayout from "@/layouts/auth";

export const Route = createFileRoute('/_authenticated')({
  beforeLoad: () => {
    // TODO:
    //  validate token
    // const authState = useAuthStore.getState().auth;
    // if (!authState.accessToken) {
    //   throw redirect({to: "/sign-in", replace: true})
    // }
  },
  component: AuthenticatedLayout,
})
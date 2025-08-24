import {createFileRoute, redirect} from '@tanstack/react-router'
import AuthenticatedLayout from "@/layouts/auth";
import {useAuthStore} from "@/states/auth-state";

export const Route = createFileRoute('/_authenticated')({
  beforeLoad: () => {
    const authState = useAuthStore.getState().auth;
    if (!authState.accessToken) {
      throw redirect({to: "/login", replace: true})
    }
  },
  component: AuthenticatedLayout,
})
import {createFileRoute, redirect} from "@tanstack/react-router";
import {useAuthStore} from "@/states/auth-state";
import {LoginPage} from "@/features/login";

export const Route = createFileRoute("/login")({
  beforeLoad: () => {
    const authState = useAuthStore.getState().auth;
    if (authState.accessToken) {
      const searchParams = new URLSearchParams(location.search);
      const redirectTo = searchParams.get("redirect") ?? "/";
      throw redirect({ to: redirectTo, replace: true });
    }
  },
  component: LoginPage,
})

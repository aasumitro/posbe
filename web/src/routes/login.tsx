import {createFileRoute, redirect} from "@tanstack/react-router";
import {useAuthStore} from "@/states/auth-state";
import {LoginPage} from "@/features/login";
import {getSafeRedirect} from "@/lib/route";

export const Route = createFileRoute("/login")({
  beforeLoad: () => {
    const authState = useAuthStore.getState().auth;
    if (authState.accessToken) {
      const redirectTo = getSafeRedirect(location.search);
      throw redirect({ to: redirectTo, replace: true });
    }
  },
  component: LoginPage,
})

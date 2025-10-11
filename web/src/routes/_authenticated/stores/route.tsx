import {createFileRoute, redirect} from '@tanstack/react-router'
import {useAuthStore} from "@/states/auth-state";
import {StoreLayout} from "@/features/stores/layout";

export const Route = createFileRoute('/_authenticated/stores')({
  beforeLoad: () => {
    const authState = useAuthStore.getState().auth;
    if (authState.user?.role?.name !== 'admin') throw redirect({to: "/", replace: true});
  },
  component: StoreLayout,
})

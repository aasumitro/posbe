import {createFileRoute, redirect} from '@tanstack/react-router'
import { MenuPage } from "@/features/orders/menus";
import {useAuthStore} from "@/states/auth-state";

export const Route = createFileRoute('/_authenticated/orders/menus')({
  beforeLoad: () => {
    const authState = useAuthStore.getState().auth;
    if (!authState.user?.role?.name) throw redirect({to: "/", replace: true});
    if (!['admin', 'cashier'].includes(authState.user?.role?.name)) throw redirect({to: "/orders/floors"});
  },
  component: MenuPage,
})

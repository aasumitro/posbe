import {Outlet, createFileRoute, useRouterState, redirect} from '@tanstack/react-router'
import {Navigation, NavigationInset} from "@/features/stores/components/navigation";
import {useAuthStore} from "@/states/auth-state";

export const Route = createFileRoute('/_authenticated/stores')({
  beforeLoad: () => {
    const authState = useAuthStore.getState().auth;
    if (authState.user?.role?.name !== 'admin') throw redirect({to: "/", replace: true});
  },
  component: StoreLayout,
})

function StoreLayout() {
  const { location } = useRouterState();
  const pathname = location.pathname;
  return (
    <div className="@container/main flex flex-1 flex-col xl:flex-row">
      {(pathname.startsWith('/stores')) && (
        <Navigation/>
      )}

      <NavigationInset>
        <Outlet />
      </NavigationInset>
    </div>
  )
}
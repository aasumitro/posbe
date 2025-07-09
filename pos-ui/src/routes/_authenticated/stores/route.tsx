import {Outlet, createFileRoute, useRouterState} from '@tanstack/react-router'
import {Navigation, NavigationInset} from "@/features/stores/components/navigation";

export const Route = createFileRoute('/_authenticated/stores')({
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
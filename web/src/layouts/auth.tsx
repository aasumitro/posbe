import {type CSSProperties, useEffect} from "react";
import {AppLeftSidebar} from "@/components/app-left-sidebar";
import {SidebarInset, SidebarProvider} from "@/components/ui/sidebar";
import {Outlet} from "@tanstack/react-router";
import {AppRightSidebar} from "@/components/app-right-sidebar";
import {UserProfileActionSheet} from "@/components/user-profile-action-sheet";
import {UserPasswordActionSheet} from "@/components/user-password-action-sheet";
import {LogoutAlertDialog} from "@/components/logout-alert-dialog";
import {useAuthStore} from "@/states/auth-state";
import {useUserDetail} from "@/hooks/use-user";

function AuthenticatedLayout() {
  const {auth} = useAuthStore()
  const {data: user} = useUserDetail(auth.userId ?? undefined)

  useEffect(() => {
    if (user?.data) {
      auth.setUser(user.data)
    }
  }, [user?.data])

  return (
    <SidebarProvider
      className="hidden lg:flex"
      style={{
        "--sidebar-width": "calc(var(--spacing) * 72)",
        "--header-height": "calc(var(--spacing) * 12)",
      } as CSSProperties}
      open={false}
    >
      <AppLeftSidebar />

      <SidebarInset>
        <Outlet />
      </SidebarInset>

      <AppRightSidebar />

      <UserProfileActionSheet />
      <UserPasswordActionSheet />
      <LogoutAlertDialog />
    </SidebarProvider>
  )
}

export default AuthenticatedLayout
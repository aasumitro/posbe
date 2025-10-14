import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupContent,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/components/ui/sidebar"
import {
  IconHome2,
  IconClipboardList,
  IconArmchair,
  IconBasketCog, IconClockRecord,
} from "@tabler/icons-react";
import {cn} from "@/lib/utils";
import {Link, useRouterState} from "@tanstack/react-router";
import {BorderBeam} from "@/components/border-beam";
import {UserMenu} from "@/components/user-menu";
import {useAuthStore} from "@/states/auth-state";
import {useActionState} from "@/states/action-state";
import {CloseShiftModalState} from "@/components/shift-action-close-alert-dialog";

const items = [
  {
    title: "Menu orders",
    url: "/orders/menus",
    icon: IconClipboardList,
    access: ["admin", "cashier"],
  },
  {
    title: "Table orders",
    url: "/orders/floors",
    icon: IconArmchair,
    access: ['admin', 'cashier', 'waiter'],
  },
  {
    title: "Close shift",
    url: "#close-shift",
    icon: IconClockRecord,
    access: ['admin', 'cashier'],
  }
]

export function AppLeftSidebar() {
  const { location } = useRouterState();
  const pathname = location.pathname;
  const { auth } = useAuthStore();
  const { setBoolState } = useActionState();

  const MidIco = (
    <span className="relative">
      O <span className="absolute top-1/2 left-0 w-full h-px bg-white rotate-60" />
    </span>
  )

  return (
    <Sidebar
      variant="inset"
      collapsible="icon"
    >
      <SidebarHeader>
        <div className={cn(
          "text-xs font-semibold tracking-wider",
          "bg-black text-white p-2 rounded-md",
          "flex items-center justify-center",
          "cursor-pointer select-none relative"
        )}>
          <BorderBeam />
          P{MidIco}S
        </div>

        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton
              className="cursor-pointer select-none mt-4"
              tooltip="Home"
              isActive={pathname === '/'}
              asChild
            >
              <Link to="/">
                <IconHome2 className={cn(
                  pathname === '/' && "text-gray-500"
                )} />
                <span>Home</span>
              </Link>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>

      <SidebarContent>
        <section className="my-auto">
          <SidebarGroup>
            <SidebarGroupContent>
              <SidebarMenu>
                {items.map((item) => {
                  if (auth.user?.role?.name &&
                    !item.access.includes(auth.user?.role?.name)) return;

                  // TODO: apply active shift validation
                  const noActiveShift = true;
                  if (item.url === "#close-shift" && noActiveShift) return;

                  return (
                    <SidebarMenuItem key={item.title}>
                      <SidebarMenuButton
                        tooltip={item.title}
                        isActive={pathname === item.url}
                        asChild
                      >
                        {item.url.includes("#") ? (
                            <button onClick={(e) => {
                              e.preventDefault();
                              const actions: Record<string, string> = {
                                "#close-shift": CloseShiftModalState,
                                // "#other-stuff": OtherModalState
                              };
                              const stateAction = actions[item.url];
                              if (!stateAction) return;
                              setBoolState(stateAction, true);
                            }} className="cursor-pointer">
                              <item.icon />
                              <span>{item.title}</span>
                            </button>
                          ) : (
                          <Link to={item.url}>
                            <item.icon className={cn(
                              pathname === item.url && "text-gray-500"
                            )}/>
                            <span>{item.title}</span>
                          </Link>
                        )}
                      </SidebarMenuButton>
                    </SidebarMenuItem>
                  )
                })}
              </SidebarMenu>
            </SidebarGroupContent>
          </SidebarGroup>
        </section>
      </SidebarContent>

      <SidebarFooter>
        <SidebarMenu>
          {auth.user?.role?.name === "admin" && (
            <SidebarMenuItem>
              <SidebarMenuButton
                className="cursor-pointer select-none mb-4"
                tooltip="Store settings - Manage teams, products, and more."
                isActive={pathname.startsWith('/stores')}
                asChild
              >
                <Link to="/stores">
                  <IconBasketCog className={cn(
                    pathname.startsWith('/stores') && "text-gray-500"
                  )} />
                  <span>Store settings</span>
                </Link>
              </SidebarMenuButton>
            </SidebarMenuItem>
          )}

          <SidebarMenuItem>
            <UserMenu />
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarFooter>
    </Sidebar>
  )
}
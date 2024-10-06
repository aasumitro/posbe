import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem, DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import {Avatar, AvatarFallback} from "@radix-ui/react-avatar";
import {LogOut, UserIcon} from "lucide-react";
import {useSessionStateStore} from "@/stores/session-state.ts";
import {useGlobalStateStore} from "@/stores/global-state.ts";
import {KEYS} from "@/lib/keys.ts";
import {useStorePageState} from "@/stores/store-state.ts";

export const UserNavigationMenu = () => {
  const {user} = useSessionStateStore();
  const {setBoolState} = useGlobalStateStore();
  const {setUserDetail} = useStorePageState();

  const openLogoutDialog = () => setBoolState(KEYS.DISPLAY_LOGOUT_MODAL, true)

  const openProfileSheet = () => {
    setUserDetail(user);
    setBoolState(KEYS.DISPLAY_USER_DETAIL_SHEET, true)
  }

  return (
    <div className="hidden w-full md:flex md:w-auto ml-auto">
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <button className="flex items-center gap-4 py-2 px-5 hover:bg-gray-50 rounded-lg">
            <div className="flex items-center justify-center h-12 w-12 bg-gray-50 rounded-full">
              <Avatar className="rounded-full">
                <AvatarFallback>
                  {user?.name.split(" ")[0][0].toUpperCase() ?? ""}
                  {user?.name.split(" ")[1][0].toUpperCase() ?? ""}
                </AvatarFallback>
              </Avatar>
            </div>
            <div className="text-left">
              <p className="text-xs font-semibold">{user?.name}</p>
              <p className="text-xs">{user?.email}</p>
            </div>
          </button>
        </DropdownMenuTrigger>
        <DropdownMenuContent className="w-44 gap-2" align="end" forceMount>
          <DropdownMenuItem onClick={openProfileSheet}>
            <UserIcon className="mr-2 h-4 w-4" />
            <span>Profile</span>
          </DropdownMenuItem>
          <DropdownMenuSeparator />
          <DropdownMenuItem onClick={openLogoutDialog}>
            <LogOut className="mr-2 h-4 w-4" />
            <span>Log out</span>
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
  )
}
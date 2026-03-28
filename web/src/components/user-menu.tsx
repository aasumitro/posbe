import {useState} from "react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger
} from "@/components/ui/dropdown-menu";
import {Avatar, AvatarFallback, AvatarImage} from "@/components/ui/avatar";
import {ProfileSheetState} from "@/components/user-profile-action-sheet";
import {PasswordSheetState} from "@/components/user-password-action-sheet";
import {LogoutModalState} from "@/components/logout-alert-dialog";
import { useActionState } from "@/states/action-state";
import {SidebarMenuButton} from "@/components/ui/sidebar";

export function UserMenu() {
  const [openMenu, setOpenMenu] = useState(false);
  const { setBoolState } = useActionState();

  return (
    <DropdownMenu open={openMenu} onOpenChange={setOpenMenu}>
      <DropdownMenuTrigger asChild>
        <SidebarMenuButton
          size="lg"
          className="data-[state=open]:bg-sidebar-accent data-[state=open]:text-sidebar-accent-foreground cursor-pointer rounded-lg"
          tooltip="My Account"
        >
          <Avatar className="rounded-lg">
            <AvatarImage
              src="https://github.com/evilrabbit.png"
              alt="@evilrabbit"
            />
            <AvatarFallback>U</AvatarFallback>
          </Avatar>
        </SidebarMenuButton>
      </DropdownMenuTrigger>

      <DropdownMenuContent className="w-56 select-none" align="start" side="right">
        <DropdownMenuLabel>Account</DropdownMenuLabel>
        <DropdownMenuGroup>
          <DropdownMenuItem
            className="cursor-pointer"
            onClick={(e) =>{
              e.preventDefault();
              setOpenMenu(false);
              setBoolState(ProfileSheetState, true);
            }}
          >
            My Profile
          </DropdownMenuItem>
          <DropdownMenuItem
            className="cursor-pointer"
            onClick={(e) =>{
              e.preventDefault();
              setOpenMenu(false);
              setBoolState(PasswordSheetState, true);
            }}
          >
            Update password
          </DropdownMenuItem>
        </DropdownMenuGroup>

        <DropdownMenuSeparator />
        <DropdownMenuItem
          className="cursor-pointer"
          onClick={(e) => {
            e.preventDefault();
            setOpenMenu(false);
            setBoolState(LogoutModalState, true);
          }}
        >
          Log out
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  )
}
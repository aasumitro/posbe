import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger
} from "@/components/ui/dropdown-menu.tsx";
import {Button} from "@/components/ui/button.tsx";
import {CheckIcon, Contact, File, FilterX, PlusCircle, SquarePen, UserPlus} from "lucide-react";
import {UserType} from "@/pages/store/components/user-tab-trigger.tsx";
import {useStorePageState} from "@/stores/store-state.ts";
import {capitalize} from "@/lib/str.ts";
import {useState} from "react";
import {KEYS} from "@/lib/keys.ts";
import {useGlobalStateStore} from "@/stores/global-state.ts";

interface UserActionFilterProps {
  tab: UserType;
  filterByRole(role: string): void;
}

export const UserActionFilter = ({tab, filterByRole}: UserActionFilterProps) => {
  const {roles} = useStorePageState();
  const [selectedRole, setSelectedRole] = useState("");
  const {setBoolState} = useGlobalStateStore();

  const onRoleSelected = (role: string) => {
    setSelectedRole(role);
    filterByRole(role);
  }

  const openNewEmployeeModal = () => {
    setBoolState(KEYS.DISPLAY_NEW_EMPLOYEE_MODAL, true)
  }

  return (
    <div className="ml-auto flex items-center gap-2">
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button size="sm" variant="outline" className="h-8 gap-2">
            <File className="h-3.5 w-3.5"/>
            <span className="sr-only sm:not-sr-only sm:whitespace-nowrap">Export</span>
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="end">
          <DropdownMenuLabel>
            Data Type
          </DropdownMenuLabel>
          <DropdownMenuSeparator/>
          <DropdownMenuItem>
            Employee
          </DropdownMenuItem>
          <DropdownMenuItem disabled>
            Customer
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
      {tab === UserType.Employee && <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button size="sm" variant="outline" className="h-8 gap-1">
            <Contact className="h-3.5 w-3.5"/>
            <span className="sr-only sm:not-sr-only sm:whitespace-nowrap">
              Roles {selectedRole !== "" && `: ${capitalize(selectedRole)}`}
            </span>
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="end">
          <DropdownMenuLabel>
            Filter by
          </DropdownMenuLabel>
          <DropdownMenuSeparator/>
          {roles && roles.map((role) => (
            <DropdownMenuItem
              key={role.id}
              onClick={() => onRoleSelected(role.name)}
              className="flex justify-between items-center"
            >
              {capitalize(role.name)}
              {selectedRole == role.name &&  <CheckIcon className="h-3.5 w-3.5" />}
            </DropdownMenuItem>
          ))}
          <DropdownMenuSeparator/>
          <DropdownMenuLabel>
            Actions
          </DropdownMenuLabel>
          <DropdownMenuSeparator/>
          <DropdownMenuItem
            className="gap-2"
            onClick={() => onRoleSelected("")}
            disabled={selectedRole === ""}
          >
            <FilterX className="h-3.5 w-3.5"/>
            Clear Filter
          </DropdownMenuItem>
          <DropdownMenuItem className="gap-2" disabled>
            <SquarePen className="h-3.5 w-3.5"/>
            Modify List
          </DropdownMenuItem>
          <DropdownMenuItem className="gap-2" disabled>
            <PlusCircle className="h-3.5 w-3.5"/>
            Add new Role
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>}
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button size="sm" className="h-8 gap-2">
            <UserPlus className="h-3.5 w-3.5"/>
            <span className="sr-only sm:not-sr-only sm:whitespace-nowrap">New User</span>
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="end">
          <DropdownMenuLabel>
            User type
          </DropdownMenuLabel>
          <DropdownMenuSeparator/>
          <DropdownMenuItem onClick={openNewEmployeeModal}>
            Employee
          </DropdownMenuItem>
          <DropdownMenuItem disabled>
            Customer
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
  )
}
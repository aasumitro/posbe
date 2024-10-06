import {TabsContent} from "@/components/ui/tabs.tsx";
import {Card, CardContent, CardDescription, CardHeader, CardTitle} from "@/components/ui/card.tsx";
import {Table, TableBody, TableCell, TableHead, TableHeader, TableRow} from "@/components/ui/table.tsx";
import {Badge} from "@/components/ui/badge.tsx";
import {
  DropdownMenu,
  DropdownMenuContent, DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuTrigger
} from "@/components/ui/dropdown-menu.tsx";
import {Button} from "@/components/ui/button.tsx";
import {MoreHorizontal} from "lucide-react";
import {Avatar, AvatarFallback} from "@radix-ui/react-avatar";
import {useStorePageState} from "@/stores/store-state.ts";
import {formatTimestamp} from "@/lib/time.ts";
import {capitalize} from "@/lib/str.ts";
import {KEYS} from "@/lib/keys.ts";
import {User} from "@/lib/types/user.ts";
import {useGlobalStateStore} from "@/stores/global-state.ts";
import {ToastMessageContainer} from "@/components/toast-message-container.tsx";
import {useSessionStateStore} from "@/stores/session-state.ts";
import {useEffect, useState} from "react";

interface UserEmployeeTabProps {
  role: string;
}

export const UserEmployeeTab = ({role}: UserEmployeeTabProps) => {
  const {users} = useStorePageState();
  const {setBoolState} = useGlobalStateStore();
  const {setUserDetail} = useStorePageState();
  const {user: access} = useSessionStateStore();
  const [userList, setUserList] = useState(users);

  const onUpdateAction = (user: User) => {
    setUserDetail(user);
    setBoolState(KEYS.DISPLAY_USER_DETAIL_SHEET, true)
  }

  const onDeleteAction = (user: User) => {
    if (access?.role.name !== "admin") {
      ToastMessageContainer("Not Authorized",
        "Sorry, you don't have authorization to access this menu!");
      return;
    }
    if (users?.filter((user) => user.role.name === "admin")
      .length === 1 && user.role.name === "admin") {
      ToastMessageContainer("Action Not Allowed",
        "Sorry, you can't remove this admin user.");
      return;
    }
    setUserDetail(user);
    setBoolState(KEYS.DISPLAY_REMOVE_EMPLOYEE_CONFIRMATION_MODAL, true)
  }

  useEffect(() => {
    if (role != "") {
      const filtered  = users?.filter(user =>
        user.role.name == role)
      setUserList(filtered ?? null)
      return;
    }

    if (role == "" || users) {
      setUserList(users)
    }
  }, [role, users]);

  return (
    <TabsContent value="employees">
      <Card>
        <CardHeader>
          <CardTitle>Employees</CardTitle>
          <CardDescription>
            Manage your employee and view their sales performance.
          </CardDescription>
        </CardHeader>
        <CardContent>
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead className="hidden w-[100px] sm:table-cell">
                  <span className="sr-only">Image</span>
                </TableHead>
                <TableHead>Name</TableHead>
                <TableHead>Role</TableHead>
                <TableHead className="hidden md:table-cell">
                  Created at
                </TableHead>
                <TableHead>
                  <span className="sr-only">Actions</span>
                </TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {userList && userList?.map((user) => (
                <TableRow key={user.id}>
                  <TableCell className="hidden sm:table-cell">
                    <div className="rounded-full w-10 h-10 p-2 bg-gray-50 flex justify-center items-center">
                      <Avatar>
                        <AvatarFallback>
                          {user.name.split(" ")[0][0].toUpperCase() ?? ""}
                          {user.name.split(" ")[1][0].toUpperCase() ?? ""}
                        </AvatarFallback>
                      </Avatar>
                    </div>
                  </TableCell>
                  <TableCell className="font-medium">
                    {user.name}
                  </TableCell>
                  <TableCell>
                    <Badge variant="outline">
                      {capitalize(user.role.name)}
                    </Badge>
                  </TableCell>
                  <TableCell className="hidden md:table-cell">
                    {formatTimestamp(user.created_at.Int64)}
                  </TableCell>
                  <TableCell>
                    <DropdownMenu>
                      <DropdownMenuTrigger asChild>
                        <Button
                          aria-haspopup="true"
                          size="icon"
                          variant="ghost"
                        >
                          <MoreHorizontal className="h-4 w-4"/>
                          <span className="sr-only">Toggle menu</span>
                        </Button>
                      </DropdownMenuTrigger>
                      <DropdownMenuContent align="end">
                        <DropdownMenuLabel>Actions</DropdownMenuLabel>
                        <DropdownMenuItem
                          onClick={() => onUpdateAction(user)}
                        >Edit</DropdownMenuItem>
                        <DropdownMenuItem
                          onClick={() => onDeleteAction(user)}
                        >Delete</DropdownMenuItem>
                      </DropdownMenuContent>
                    </DropdownMenu>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </CardContent>
      </Card>
    </TabsContent>
  )
}
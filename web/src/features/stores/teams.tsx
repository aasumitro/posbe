import {TeamContainer} from "@/features/stores/components/team/team-container";
import {useUserState} from "@/states/user-state";
import {useRoleList, useUserList} from "@/hooks/use-user";
import {useEffect} from "react";

export function TeamMemberPage() {
  const {data: roles, isPending: isLoadRoles} = useRoleList()
  const {data: users, isPending: isLoadUsers} = useUserList()
  const {setRoles, setUsers} =  useUserState();

  useEffect(() => {
    if (roles?.data) {
      setRoles(roles.data)
    }
    if (users?.data) {
      setUsers(users?.data ?? null)
    }
  }, [roles?.data, users?.data]);

  if (isLoadRoles || isLoadUsers) return <>Loading . . .</>

  return (
    <div className="w-full pt-4 xl:pt-6 space-y-6">
      <aside className="flex items-center justify-between gap-4 px-8">
        <div className="block">
          <h3 className="text-lg font-semibold mb-2">Teams</h3>
          <p className="text-sm text-muted-foreground">
            Manage your organization's teams. Add, view, and organize team structures to match your operational needs.
          </p>
        </div>
      </aside>

      <TeamContainer />
    </div>
  )
}
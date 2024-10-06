import {Tabs } from "@/components/ui/tabs.tsx";
import {UserTabTrigger, UserType} from "@/pages/store/components/user-tab-trigger.tsx";
import {UserActionFilter} from "@/pages/store/components/user-action-filter.tsx";
import {UserEmployeeTab} from "@/pages/store/components/user-employee-tab.tsx";
import {UserCustomerTab} from "@/pages/store/components/user-customer.tab.tsx";
import {useState} from "react";
import {useStoreUser} from "@/hooks/use-store-user.tsx";
import {useStorePageState} from "@/stores/store-state.ts";
import {useQuery} from "react-query";
import {HttpResponse} from "@/lib/types/http-response.ts";
import {Role, User as UType} from "@/lib/types/user.ts";
import {HandleRequestLoading} from "@/components/request-loading.tsx";
import {RemoveEmployeeModal} from "@/pages/store/components/remove-employee-modal.tsx";
import {NewEmployeeModal} from "@/pages/store/components/new-employee-modal.tsx";

export const User = () => {
  const [tab, setTab] = useState(UserType.Employee)
  const {getUsers, getRoles} = useStoreUser();
  const {setUserList, setUserRoleList} = useStorePageState();
  const [roleFilter, setRoleFilter] = useState("");

  const userList = useQuery<HttpResponse<UType[]>>(
    "store.users", getUsers, {
      onSuccess: (resp) => setUserList(resp.data),
      onError: (error) => console.log(error),
      retry: false,
    });

  const roleList = useQuery<HttpResponse<Role[]>>(
    "user.roles", getRoles, {
      onSuccess: (resp) => setUserRoleList(resp.data),
      onError: (error) => console.log(error),
      retry: false,
    });

  if (userList.isLoading || roleList.isLoading) {
    return <HandleRequestLoading />
  }

  const onTabChange = (tab: UserType) => setTab(tab)

  const onRoleFilterChange = (role: string) => setRoleFilter(role);

  return (
    <>
      <Tabs defaultValue={tab}>
        <div className="flex items-center">
          <UserTabTrigger action={onTabChange} />
          <UserActionFilter tab={tab} filterByRole={onRoleFilterChange}/>
        </div>
        <UserCustomerTab />
        <UserEmployeeTab role={roleFilter} />
      </Tabs>
      <NewEmployeeModal />
      <RemoveEmployeeModal />
    </>
  )
}
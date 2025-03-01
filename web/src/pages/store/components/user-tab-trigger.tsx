import {TabsList, TabsTrigger} from "@/components/ui/tabs.tsx";

export enum UserType {
  Employee = "employees",
  Customer = "customers"
}

interface UserTabTrigger {
  action(tab: UserType): void;
}

export const UserTabTrigger = ({action}: UserTabTrigger) => {
  return (
    <TabsList>
      <TabsTrigger
        value={UserType.Employee}
        onClick={() => action(UserType.Employee)}
      >Employees</TabsTrigger>
      <TabsTrigger
        value={UserType.Customer}
        onClick={() => action(UserType.Customer)}
      >Customers</TabsTrigger>
    </TabsList>
  )
}
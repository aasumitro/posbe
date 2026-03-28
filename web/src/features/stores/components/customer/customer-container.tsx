import {CustomerActionAdd} from "@/features/stores/components/customer/customer-action-add";
import {CustomerCard} from "@/features/stores/components/customer/customer-card";
import {CustomerDetailSheet} from "@/features/stores/components/customer/customer-detail-sheet";
import {useCustomerState} from "@/states/customer-state";

export function CustomerContainer() {
  const {customers} = useCustomerState();

  return (
    <aside className="flex flex-wrap gap-4 px-8 py-4">
      <CustomerActionAdd />

      {customers?.map((customer, index) => (
        <CustomerCard key={index} customer={customer} />
      ))}

      <CustomerDetailSheet />
    </aside>
  )
}
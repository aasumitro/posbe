import {CustomerActionAdd} from "@/features/stores/components/customer/customer-action-add";
import {CustomerCard} from "@/features/stores/components/customer/customer-card";
import {CustomerDetailSheet} from "@/features/stores/components/customer/customer-detail-sheet";

export function CustomerContainer() {
  const customers = [1]

  return (
    <aside className="flex flex-wrap gap-4 px-8 py-4">
      <CustomerActionAdd />

      {customers.map((index) => (
        <CustomerCard key={index} />
      ))}

      <CustomerDetailSheet />
    </aside>
  )
}
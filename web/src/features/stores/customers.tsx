import {useEffect} from "react";
import {CustomerContainer} from "@/features/stores/components/customer/customer-container";
import {useCustomerList} from "@/hooks/use-customer";
import {useCustomerState} from "@/states/customer-state";

export function CustomersPage() {
  const {data: customers, isPending} = useCustomerList()
  const {setCustomers} = useCustomerState();

  useEffect(() => {
    if (customers?.data) {
      setCustomers(customers?.data ?? null)
    }
  }, [customers?.data]);

  if (isPending) return <>Loading . . .</>

  return (
    <div className="w-full pt-4 xl:pt-6 space-y-6">
      <aside className="flex items-center justify-between gap-4 px-8">
        <div className="block">
          <h3 className="text-lg font-semibold mb-2">Customers</h3>
          <p className="text-sm text-muted-foreground">
            Manage your customer records and track their interactions in one place.
          </p>
        </div>
      </aside>

      <CustomerContainer />
    </div>
  )
}
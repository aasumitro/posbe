import {CustomerContainer} from "@/features/stores/components/customer/customer-container";

export function CustomersPage() {
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
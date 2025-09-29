import {useState} from "react";
import {TransactionDetail} from "@/features/stores/components/transaction/transaction-detail";
import {TransactionList} from "@/features/stores/components/transaction/transaction-list";
import {AppComingSoon} from "@/components/app-coming-soon";

export function TransactionPage() {
  const [soon] = useState(false);

  if (!soon) {
    return  <AppComingSoon feature="transactions" type="page" />
  }

  return (
    <div className="w-full pt-4 xl:pt-6 space-y-6">
      <aside className="flex items-center justify-between gap-4 px-8">
        <div className="block">
          <h3 className="text-lg font-semibold mb-2">Transactions</h3>
          <p className="text-sm text-muted-foreground">
            View and manage your recent sales, refunds, and transaction history.
          </p>
        </div>
      </aside>

      <aside className="grid grid-cols-3 gap-4 px-8">
        <div className="col-span-2">
          <TransactionList />
        </div>

        <div>
          <TransactionDetail />
        </div>
      </aside>
    </div>
  )
}
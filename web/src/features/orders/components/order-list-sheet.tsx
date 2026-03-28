import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
} from "@/components/ui/sheet"
import {useActionState} from "@/states/action-state";
import {useEffect, useState} from "react";
import {IconReceiptOff} from "@tabler/icons-react";
import {Badge} from "@/components/ui/badge";
import {useOrderState} from "@/states/order-state";
import {HorizontalScrollItems} from "@/components/horizontal-scroll-items";
import {OrderCard} from "@/features/orders/components/order-card";

export const OrderListSheetState = "order_list_sheet_state"

type FilterOption =
  | "All"
  | "Active"
  | "Cancelled"
  | "Completed";

const filterOptions: FilterOption[] = [
  "All",
  "Active",
  "Cancelled",
  "Completed",
];

export function OrderListSheet() {
  const [filter, setFilter] = useState<FilterOption>("All")
  const [openOrderListSheet, setOrderListSheetOpen] = useState(false);
  const { bool, setBoolState } = useActionState();
  const { orders } =  useOrderState();

  useEffect(() => {
    if (bool[OrderListSheetState]){
      setOrderListSheetOpen(bool[OrderListSheetState])
    }
  }, [bool]);

  const applyFilter = (filter: FilterOption) => {
    setFilter(filter);

    // TODO: apply for the data query
  }

  function onOpenChange(newOpen: boolean) {
    setOrderListSheetOpen(newOpen);
    if (!newOpen) {
      setBoolState(OrderListSheetState, false);
    }
  }

  return (
    <Sheet open={openOrderListSheet} onOpenChange={onOpenChange}>
      <SheetContent>
        <SheetHeader>
          <SheetTitle>Order List ({orders?.length ?? 0})</SheetTitle>
          <SheetDescription>
            View all current orders in this session. You can review, update, or manage them from here.
          </SheetDescription>
        </SheetHeader>

        <div className="px-4 space-y-2 overflow-y-auto mb-6">
          {!orders && (
            <div className="text-center py-12">
              <IconReceiptOff className="my-4 w-14 h-14 mx-auto"/>
              <h4 className="text-primary text-md font-bold tracking-tight">
                No Orders Found
              </h4>
              <p className="text-secondary-foreground text-xs font-normal">
                There are no active orders at the moment.
              </p>
            </div>
          )}

          {orders && orders.length > 0 && (
            <>
              <HorizontalScrollItems>
                <div className="flex gap-2 w-max select-none">
                  {filterOptions.map((fl) => (
                    <Badge
                      key={fl}
                      onClick={() => applyFilter(fl)}
                      variant={(fl === filter) ? "default" : "outline"}
                      className="hover:cursor-pointer h-8"
                    >{fl}</Badge>
                  ))}
                </div>
              </HorizontalScrollItems>

              {orders?.map((order) => (
                  <OrderCard order={order} className=" w-full" />
              ))}
            </>
          )}
        </div>
      </SheetContent>
    </Sheet>
  )
}
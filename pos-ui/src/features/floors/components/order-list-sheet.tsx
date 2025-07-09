import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
} from "@/components/ui/sheet"
import {useActionState} from "@/states/action-state";
import {useEffect, useState} from "react";
import {IconPointFilled} from "@tabler/icons-react";
import {Tooltip, TooltipContent, TooltipTrigger} from "@/components/ui/tooltip";
import {Badge} from "@/components/ui/badge";

export const OrderListSheetState = "order_list_sheet_state"

export function OrderListSheet() {
  const [openOrderListSheet, setOrderListSheetOpen] = useState(false);
  const { bool, setBoolState } = useActionState();
  const orders = [1]

  useEffect(() => {
    if (bool[OrderListSheetState]){
      setOrderListSheetOpen(bool[OrderListSheetState])
    }
  }, [bool]);

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
          <SheetTitle>Order List (20)</SheetTitle>
          <SheetDescription>
            View all current orders in this session. You can review, update, or manage them from here.
          </SheetDescription>
        </SheetHeader>

        <div className="px-4 space-y-2  overflow-y-auto mb-6">
          {orders.map((index) => (
            <div
              key={index}
              className={`flex flex-col p-4 rounded-xl w-full border cursor-pointer hover:bg-green-50`}
            >
              <div className="flex justify-between items-center w-full">
                <h5 className="text-lg">Zeros Mardigu</h5>
                <p className="text-muted-foreground text-sm">#0{index < 9 && "0"}{index}</p>
              </div>
              <div className="flex items-center w-full gap-1">
                <p className="text-muted-foreground text-xs">3 items</p>
                <IconPointFilled className="w-2 h-2 text-muted-foreground" />
                <p className="text-muted-foreground text-xs">Table TA1</p>
              </div>
              <div className="mt-2">
                <p className="text-sm">Order:</p>
                <Tooltip>
                  <TooltipTrigger asChild>
                    <p className="text-xs text-muted-foreground truncate whitespace-nowrap overflow-hidden text-ellipsis">
                      1x Daging Enak Banget - (normal), 1x Daging Enak Banget - (half), 2x Teh Manis Banget - (xl/25)
                    </p>
                  </TooltipTrigger>
                  <TooltipContent>
                    <p>
                      1x Daging Enak Banget - (normal), 1x Daging Enak Banget - (half), 2x Teh Manis Banget - (xl/25)
                    </p>
                  </TooltipContent>
                </Tooltip>

              </div>
              <Badge variant="secondary" className="mt-4">Ready to serve</Badge>
            </div>
          ))}
        </div>
      </SheetContent>
    </Sheet>
  )
}
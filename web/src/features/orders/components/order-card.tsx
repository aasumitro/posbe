import type {Order} from "@/types/order";
import {IconPointFilled} from "@tabler/icons-react";
import {Tooltip, TooltipContent, TooltipTrigger} from "@/components/ui/tooltip";
import {Badge} from "@/components/ui/badge";
import {cn} from "@/lib/utils";

export function OrderCard({
  order, className
}: { order: Order, className?: string }) {
  return (
    <div
      key={order.id}
      className={cn('flex flex-col p-4 rounded-xl border cursor-pointer hover:bg-green-50', className)}
    >
      <div className="flex justify-between items-center w-full">
        <h5 className="text-lg">Zeros Mardigu</h5>
        <p className="text-muted-foreground text-sm">
          #{String(order.id).padStart(4, "0")}
        </p>
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
  )
}
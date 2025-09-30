import {IconPointFilled} from "@tabler/icons-react";
import {Badge} from "@/components/ui/badge";
import {Tooltip, TooltipContent, TooltipTrigger} from "@/components/ui/tooltip";

export function OrderList() {
  const orders = [1]

  return (
    <div className="flex gap-3 overflow-x-auto pb-4 px-5 ml-[-12px] mr-[-12px]">
      {orders.map((index) => (
        <div
          key={index}
          className={`flex flex-col p-4 rounded-xl w-[300px] border cursor-pointer hover:bg-green-50`}
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
  )
}

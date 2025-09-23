import {useActionState} from "@/states/action-state";
import {Card, CardDescription, CardFooter, CardHeader, CardTitle} from "@/components/ui/card";
import {Separator} from "@/components/ui/separator";
import {CustomerDetailActionSheetState} from "@/features/stores/components/customer/customer-detail-sheet";

export function CustomerCard() {
  const { setBoolState } = useActionState();

  return (
    <Card
      className="w-96 max-h-56 hover:bg-gray-50/50 select-none cursor-pointer"
      onClick={(e) => {
        e.preventDefault();
        setBoolState(CustomerDetailActionSheetState, true);
      }}
    >
      <CardHeader className="pb-0 flex flex-row gap-4">
        <div className="w-[80px] h-fit">
          <img
            className="w-full h-full rounded-md"
            alt="placeholder"
            src="https://placehold.jp/150x150.png"
          />
        </div>
        <div className="w-fit space-y-2">
          <CardTitle className="mt-4">
            Customer Good
          </CardTitle>
          <CardDescription>
            +6282275558899
          </CardDescription>
        </div>
      </CardHeader>
      <CardFooter className="flex flex-row border-t px-4 [.border-t]:pt-4">
        <div className="flex w-full items-center gap-2">
          <div className="grid flex-1 auto-rows-min gap-0.5 pl-4">
            <div className="text-xs text-muted-foreground">Orders</div>
            <div className="flex items-baseline gap-1 text-2xl font-bold tabular-nums leading-none">
              250
              <span className="text-sm font-normal text-muted-foreground">
                items
              </span>
            </div>
          </div>
          <Separator orientation="vertical" className="mx-2 h-10 w-px" />
          <div className="grid flex-1 auto-rows-min gap-0.5">
            <div className="text-xs text-muted-foreground">Spents</div>
            <div className="flex items-baseline gap-1 text-2xl font-bold tabular-nums leading-none">
              IDR 2,67
              <span className="text-sm font-normal text-muted-foreground">
                 Mio
              </span>
            </div>
          </div>
        </div>
      </CardFooter>
    </Card>
  )
}
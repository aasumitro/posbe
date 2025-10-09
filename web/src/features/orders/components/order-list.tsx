import {Button} from "@/components/ui/button";
import {useActionState} from "@/states/action-state";
import {OrderListSheetState} from "@/features/orders/components/order-list-sheet";
import {useOrderState} from "@/states/order-state";
import {OrderCard} from "@/features/orders/components/order-card";

export function OrderList() {
  const { setBoolState } = useActionState();
  const { orders } =  useOrderState();

  return (
    <div className="flex gap-3 overflow-x-auto pb-4 px-5 ml-[-12px] mr-[-12px]">
      {orders?.map((order) => (
        <OrderCard order={order} className="w-[300px]" />)
      )}

      <Button
        variant="link"
        className="my-auto cursor-pointer"
        onClick={(e) => {
          e.preventDefault();
          setBoolState(OrderListSheetState, true);
        }}
      >
        View More Orders
      </Button>
    </div>
  )
}

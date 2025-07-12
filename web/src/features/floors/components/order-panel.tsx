import {useActionState} from "@/states/action-state";
import {OrderListSheetState} from "@/features/floors/components/order-list-sheet";

export function OrderPanel() {
  const { setBoolState } =useActionState();

  return (
    <div className="fixed top-0 bottom-0 right-2 lg:flex items-center z-50">
      <div
        className="w-8 h-32 bg-gray-100 shadow-sm rounded-l-xl cursor-pointer flex items-center justify-center"
        onClick={(e) => {
          e.preventDefault();
          setBoolState(OrderListSheetState, true);
        }}
      >
        <p className="-rotate-90 text-sm font-semibold">Orders</p>
      </div>
    </div>
  )
}
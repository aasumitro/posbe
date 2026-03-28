import {useEffect} from "react";
import {ActionVisibleRightSidebarKey} from "@/components/app-right-sidebar";
import { useActionState } from "@/states/action-state";
import {MenuContainer} from "@/features/orders/components/menu-container";
import {useOrderState} from "@/states/order-state";
import {OrderList} from "@/features/orders/components/order-list";

export function MenuPage() {
  const { setBoolState } = useActionState();
  const { orders } =  useOrderState();

  useEffect(() => {
    setBoolState(ActionVisibleRightSidebarKey, true)
  }, [setBoolState]);

  return (
    <div className="@container/main flex flex-1 flex-col gap-2 p-4">
      {orders && orders.length > 0 &&  (
        <>
          <h5 className="text-lg font-bold">
            Order List
          </h5>
          
          <OrderList />
        </>
      )}

      <MenuContainer />
    </div>
  )
}
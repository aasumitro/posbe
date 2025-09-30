import {useEffect} from "react";
import {ActionVisibleRightSidebarKey} from "@/components/app-right-sidebar";
import { useActionState } from "@/states/action-state";
import { OrderContainer } from "@/features/orders/components/order-container";
import {MenuContainer} from "@/features/orders/components/menu-container";

export function MenuPage() {
  const { setBoolState } = useActionState();

  useEffect(() => {
    setBoolState(ActionVisibleRightSidebarKey, true)
  }, [setBoolState]);

  return (
    <div className="@container/main flex flex-1 flex-col gap-2 p-4">
      <OrderContainer />

      <MenuContainer />
    </div>
  )
}
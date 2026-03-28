import {
  Drawer,
  DrawerContent,
  DrawerHeader,
  DrawerTitle,
  DrawerDescription,
} from "@/components/ui/drawer"
import {useEffect, useState} from "react";
import {useActionState} from "@/states/action-state";
import {MenuContainer} from "@/features/orders/components/menu-container";
import {Cart} from "@/features/orders/components/cart";

export const MenuOrderDrawerState = "menu_order_drawer_state"

export function MenuOrderDrawer() {
  const [isOpen, setIsOpen] = useState(false);
  const { bool, setBoolState } = useActionState();

  useEffect(() => {
    if (!bool[MenuOrderDrawerState]) return;
    setIsOpen(bool[MenuOrderDrawerState]);
  }, [bool]);

  const onClose = () => {
    setBoolState(MenuOrderDrawerState, false)
    setIsOpen(false);
  }

  return (
    <Drawer
      open={isOpen}
      onOpenChange={onClose}
      repositionInputs={false}
      autoFocus={isOpen}
    >
      <DrawerContent
        className="min-h-[93.5%] h-full"
        aria-description="make-orders"
        aria-describedby="make-orders"
      >
        <DrawerHeader>
          <DrawerTitle />
          <DrawerDescription />
        </DrawerHeader>

        <div className="px-6 pb-6 grid grid-cols-4 h-full">
          <MenuContainer className="col-span-3 pr-8" />
          <Cart className="col-span-1 bg-gray-100/50 rounded-lg" />
        </div>
      </DrawerContent>
    </Drawer>
  )
}
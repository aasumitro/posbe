import {Sidebar} from "@/components/ui/sidebar";
import * as React from "react";
import {Cart} from "@/features/menus/components/cart";

export function MenuSidebarSection({...props}: React.ComponentProps<typeof Sidebar>) {
  return (
    <Sidebar
      className="sticky top-0 h-svh w-80 xl:w-96"
      collapsible="none"
      side="right"
      variant="inset"
      {...props}
    >
      <Cart />
    </Sidebar>
  )
}
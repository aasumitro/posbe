import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuPortal,
  DropdownMenuSeparator,
  DropdownMenuSub,
  DropdownMenuSubContent,
  DropdownMenuSubTrigger,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { cn } from "@/lib/utils"
import type {DraggableTableItem, TableStatus} from "@/components/table";
import { getLabelBgColorByStatus } from "@/lib/table";
import {useAuthStore} from "@/states/auth-state";
import {useState} from "react";

export function TableMenu(
  {table, actions}: {
    table: DraggableTableItem,
    actions?: {
      handlePlaceOrder: (tableId: number) => void,
      handleEditOrder: (tableId: number) => void,
      handlePrintBill: (tableId: number) => void,
      handleProcessPayment: (tableId: number) => void,
      handlerSetStatus: (tableId: number, status: string) => void,
      handleViewHistory: (tableId: number) => void,
    }
  }
) {
  const [open, setOpen] = useState(false);
  const { auth } = useAuthStore();

  const canManageTable = ["admin", "cashier", "waiter"].includes(auth?.user?.role?.name ?? "");
  const canViewHistory = ["admin", "cashier"].includes(auth?.user?.role?.name ?? "");

  const tableMenu = [
    {
      label: "Place an Order",
      onClick: () =>  actions?.handlePlaceOrder(table.id),
      display: canManageTable && !table.hasOrder
    },
    {
      label: "View / Edit Order",
      onClick: () =>  actions?.handleEditOrder(table.id),
      display: canManageTable && table.hasOrder
    },
    {
      label: "View / Print Bill",
      onClick:() => actions?.handlePrintBill(table.id),
      display: canViewHistory && table.hasOrder
    },
    {
      label: "Process Payment",
      onClick:() => actions?.handleProcessPayment(table.id),
      display: canViewHistory && table.hasOrder
    },
    { label: "separator", display: true},
    {
      label: "Set status",
      submenu: [
        {
          label: "Available",
          onClick: () => actions?.handlerSetStatus(table.id,"AVAILABLE"),
          display: canManageTable && table.hasOrder
        },
        { label: "separator", display: true},
        {
          label: "Occupied",
          onClick: () => actions?.handlerSetStatus(table.id,"OCCUPIED"),
          display: canManageTable && !table.hasOrder
        },
        {
          label: "Reserved",
          onClick: () => actions?.handlerSetStatus(table.id,"RESERVED"),
          display: canManageTable && !table.hasOrder
        },
        { label: "separator", display: true },
        {
          label: "Needs Cleaning",
          onClick: () => actions?.handlerSetStatus(table.id,"NEEDS_CLEANING"),
          display: canManageTable
        },
      ],
      display: canManageTable
    },
    { label: "separator", display: true},
    {
      label: "View History",
      onClick: () => actions?.handleViewHistory(table.id),
      display: canViewHistory && table.hasHistory
    },
  ];

  const bgColor = getLabelBgColorByStatus(table.status as TableStatus)

  const visibleMenu = tableMenu.filter(item => item.display);

  const handleMenuAction = (fn?: () => void) => {
    console.log("handleMenuAction", fn);
    if (fn) fn();
    setOpen(false);
  };

  return (
    <DropdownMenu open={open} onOpenChange={setOpen}>
      <DropdownMenuTrigger asChild>
        <div className={cn(
          "w-full h-full",
          table.config.shape === "circle" && "p-[2px]"
        )}>
          <div
            className={cn(
              "w-full h-full flex justify-center items-center text-xs",
              "cursor-pointer transition-colors hover:brightness-90",
              table.status === "available" && "text-white",
              table.config.shape === "circle" ? "rounded-full": "rounded-sm",
              table.config.shape === "rectangle" && table.config.chairs <= 2 && "-rotate-90"
            )}
            style={{ backgroundColor: bgColor }}
          >{table.name}</div>
        </div>
      </DropdownMenuTrigger>

      <DropdownMenuContent className="w-56" align="start">
        <DropdownMenuGroup>
          {visibleMenu.map((menu, index) => {
            if (menu.label === "separator") {
              // only render if previous or next item exists
              const prev = visibleMenu[index - 1];
              const next = visibleMenu[index + 1];
              if (!prev || !next || prev.label === "separator" || next.label === "separator") return null;
              return <DropdownMenuSeparator key={index} />;
            }

            if (menu?.submenu) {
              const visibleSubmenu = menu.submenu.filter(sub => sub.display);

              return (
                <DropdownMenuSub key={index}>
                  <DropdownMenuSubTrigger>{menu.label}</DropdownMenuSubTrigger>
                  <DropdownMenuPortal>
                    <DropdownMenuSubContent>
                      {visibleSubmenu.map((submenu, subIndex) => {
                        if (submenu.label === "separator") {
                          const prev = visibleSubmenu[subIndex - 1];
                          const next = visibleSubmenu[subIndex + 1];
                          if (!prev || !next || prev.label === "separator" || next.label === "separator") return null;
                          return <DropdownMenuSeparator key={subIndex} />;
                        }

                        return (
                          <DropdownMenuItem
                            key={subIndex}
                            onClick={() => handleMenuAction(submenu.onClick)}
                          >{submenu.label}</DropdownMenuItem>
                        )
                      })}
                    </DropdownMenuSubContent>
                  </DropdownMenuPortal>
                </DropdownMenuSub>
              )
            }

            return (
              <DropdownMenuItem
                key={index}
                onClick={() => handleMenuAction(menu.onClick)}
              >{menu.label}</DropdownMenuItem>
            )
          })}
        </DropdownMenuGroup>
      </DropdownMenuContent>
    </DropdownMenu>
  )
}

"use client"

import { type DraggableTableItem } from "@/components/table"
import { ZoomableContainer } from "@/components/zoomable-container";
import { Table } from "@/components/table";
import TableStatusPanel from "@/features/orders/components/status-panel";
import {FloorPanel} from "@/features/orders/components/floor-panel";
import {TableMenu} from "@/features/orders/components/table-menu";
import {OrderPanel} from "@/features/orders/components/order-panel";
import {getTableScale} from "@/lib/table";
import {cn} from "@/lib/utils";

interface FloorManagementProps {
  tables: DraggableTableItem[]
}

export function FloorManagement({
  tables,
}: FloorManagementProps) {
  function handlePlaceOrder(tableId: number) {
    // TODO:
    //  1. Display a modal for creating a new order.
    //  2. Allow employees to select menu items, set quantities, and confirm the order.
    //  3. Save and link the order to the selected table once confirmed.
   alert(`Place order ${tableId}`)
  }

  function handleEditOrder(tableId: number) {
    // TODO:
    //  1. Display the current order details in a modal.
    //  2. Allow employees to manage the order:
    //     - Add new items.
    //     - Modify or remove items that have not yet been processed.
    //  3. Save and update the order once changes are confirmed.
    alert(`Edit order ${tableId}`)
  }

  function handlePrintBill(tableId: number) {
    // TODO:
    //  1. Generate and format the bill or receipt.
    //  2. If compatible printing hardware is available, integrate direct printing in the future.
    //  3. For now, export the bill as a PDF for manual printing.
    alert(`Print bill ${tableId}`)
  }

  function handleProcessPayment(tableId: number) {
    // TODO:
    //  1. Open a confirmation modal for payment processing.
    //  2. In the modal, display the ordered items, bill details, and total amount.
    //  3. Allow the user to select a payment method (e.g., cash, card, e-wallet).
    //  4. Confirm and finalize the payment once approved.
    alert(`Process payment ${tableId}`)
  }

  function handlerSetStatus(tableId: number, status: string) {
    // TODO:
    //  direct call api endpoint
    alert(`set status ${tableId}: ${status}`)
  }

  function handleViewHistory(tableId: number) {
    // TODO:
    //  open order panel with table id filter
    alert(`view history ${tableId}`)
  }

  return (
    <div
      className="relative h-full w-full overflow-hidden"
      onPointerDown={(e) => e.stopPropagation()}
    >
      <div className="absolute inset-0 overflow-hidden">
        <ZoomableContainer>
          {tables.map((table, index) => {
            return (
              <div
                key={index}
                className={cn(
                  "flex items-center justify-center",
                  table.config.shape === "rectangle" &&
                  table.config.chairs <= 2 && "rotate-90"
                )}
                style={{
                  position: "absolute",
                  left: table.xPos,
                  top: table.yPos,
                  transform: getTableScale(
                    table.config.shape,
                    table.config.chairs
                  ),
                  transformOrigin: "center",
                }}
              >
                <Table
                  id={table.id}
                  name={table.name}
                  status={table.status}
                  customers={table.customers}
                  config={table.config}
                  menu={<TableMenu
                    table={table}
                    actions={{
                      handlePlaceOrder,
                      handleEditOrder,
                      handlePrintBill,
                      handleProcessPayment,
                      handlerSetStatus,
                      handleViewHistory,
                    }}
                  />}
                />
              </div>
            )
          })}
        </ZoomableContainer>
      </div>

      <FloorPanel />
      <OrderPanel />
      <TableStatusPanel />
    </div>
  )
}

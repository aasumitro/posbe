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
                  transform: getTableScale(table.config.shape, table.config.chairs),
                  transformOrigin: "center",
                }}
              >
                <Table
                  id={table.id}
                  name={table.name}
                  status={table.status}
                  customers={table.customers}
                  config={table.config}
                  menu={<TableMenu table={table} />}
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

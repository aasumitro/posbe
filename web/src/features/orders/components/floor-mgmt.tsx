"use client"

import { useState } from "react"
import { type DraggableTableItem } from "@/components/table"
import { ZoomableContainer } from "@/components/zoomable-container";
import { Table } from "@/components/table";
import TableStatusPanel from "@/features/orders/components/status-panel";
import {FloorPanel} from "@/features/orders/components/floor-panel";
import {TableMenu} from "@/features/orders/components/table-menu";
import {OrderPanel} from "@/features/orders/components/order-panel";
import {OrderListSheet} from "@/features/orders/components/order-list-sheet";

interface FloorManagementProps {
  tables: DraggableTableItem[]
}

export function FloorManagement({
  tables,
}: FloorManagementProps) {
  const [scale] = useState(1)

  return (
    <div
      className="relative h-full w-full overflow-hidden"
      onPointerDown={(e) => e.stopPropagation()}
    >
      <div className="absolute inset-0 overflow-hidden">
        <ZoomableContainer>
          {tables.map((table, index) => {
            return (
              <div key={index} style={{
                position: "absolute",
                left: table.xPos * scale,
                top: table.yPos * scale,
              }}>
                <Table
                  id={table.id}
                  name={table.name}
                  status={table.status}
                  customers={table.customers}
                  config={table.config}
                  menu={
                    <TableMenu
                      name={table.name}
                      status={table.status}
                      shape={table.config.shape}
                    />
                  }
                />
              </div>
            )
          })}
        </ZoomableContainer>
      </div>

      <FloorPanel />
      <OrderPanel />
      <OrderListSheet />
      <TableStatusPanel />
    </div>
  )
}

"use client"

import type {DraggableTableItem} from "@/components/table";
import {useEffect, useState} from "react";
import {ZoomableContainer} from "@/components/zoomable-container";
import {cn} from "@/lib/utils";
import {DraggableTable} from "@/features/stores/components/floor/draggable-table";
import {FloorPanel} from "@/features/stores/components/floor/floor-panel";
import {UpdateActionDock} from "@/features/stores/components/floor/update-action";
import {TableActionNew} from "@/features/stores/components/floor/table-action-new";
import {TableActionDelete} from "@/features/stores/components/floor/table-action-delete";

interface FloorDesignerProps {
  initialTables?: DraggableTableItem[]
}

export function FloorDesigner({
  initialTables = [],
}: FloorDesignerProps) {
  const [tables, setTables] = useState<DraggableTableItem[]>(initialTables)
  const [selectedItem, setSelectedItem] = useState<number | null>(null)
  const [isDraggingTable, setIsDraggingTable] = useState(false)

  useEffect(() => {
    setTables(initialTables)
  }, [initialTables])

  const moveItem = (id: number, x: number, y: number) => {
    setTables(tables.map((table) => (table.id === id ? { ...table, xPos: x, yPos: y } : table)))
  }

  return (
    <div  className="relative h-full w-full overflow-hidden">
      {/* Floor Plan Canvas */}
      <p className="hidden">{selectedItem}</p>
      <div className="absolute inset-0 overflow-hidden p-2">
        <ZoomableContainer className={cn(
          isDraggingTable ? "no-pan" : ""
        )}>
          {tables.map((table) => (
            <DraggableTable
              key={table.id}
              item={table}
              onMove={moveItem}
              onSelect={setSelectedItem}
              isSelected={false}
              scale={0.75}
              setDragging={setIsDraggingTable}
            />
          ))}
        </ZoomableContainer>
      </div>

      <TableActionNew />

      <FloorPanel />

      <UpdateActionDock />

      <TableActionDelete />
    </div>
  )
}
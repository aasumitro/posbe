"use client"

import type {DraggableTableItem} from "@/components/table";
import {useEffect, useState} from "react";
import {ZoomableContainer} from "@/components/zoomable-container";
import {cn} from "@/lib/utils";
import {DraggableTable} from "@/features/stores/components/floor/draggable-table";
import {FloorPanel} from "@/features/stores/components/floor/floor-panel";
import {UpdateActionDock} from "@/features/stores/components/floor/update-action";
import {NewTableAction} from "@/features/stores/components/floor/new-table-action";

interface FloorDesignerProps {
  initialTables?: DraggableTableItem[]
}

export function FloorDesigner({
  initialTables = [],
}: FloorDesignerProps) {
  const [tables, setTables] = useState<DraggableTableItem[]>(initialTables)
  const [selectedItem, setSelectedItem] = useState<string | null>(null)
  const [isDraggingTable, setIsDraggingTable] = useState(false)

  useEffect(() => {
    setTables(initialTables)
  }, [initialTables])

  const moveItem = (id: string, x: number, y: number) => {
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
              scale={0.8}
              setDragging={setIsDraggingTable}
            />
          ))}
        </ZoomableContainer>
      </div>

      <NewTableAction />

      <FloorPanel />

      <UpdateActionDock />
    </div>
  )
}
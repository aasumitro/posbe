"use client"

import type {DraggableTableItem} from "@/components/table";
import {useEffect, useState} from "react";
import {ZoomableContainer} from "@/components/zoomable-container";
import {cn} from "@/lib/utils";
import {DraggableTable} from "@/features/stores/components/floor/draggable-table";
import {FloorPanel} from "@/features/stores/components/floor/floor-panel";
import {UpdateActionDock} from "@/features/stores/components/floor/dock-action-update";
import {TableActionNew} from "@/features/stores/components/floor/table-action-new";
import {TableActionDelete} from "@/features/stores/components/floor/table-action-delete";
import {useUpdateTable} from "@/hooks/use-seating";
import {AppLoading} from "@/components/app-loading";
import {toast} from "sonner";
import {isHTTPResponse} from "@/lib/api";
import {useSeatingState} from "@/states/seating-state";
import {useQueryClient} from "@tanstack/react-query";

interface FloorDesignerProps {
  initialTables?: DraggableTableItem[]
}

type tableErrorResponse = {
  floor_id?: string[]
  name?: string[]
  x_pos?: string[]
  y_pos?: string[]
  w_size?: string[]
  h_size?: string[]
  d_size?: string[]
  capacity?: string[]
  type?: string[]
}

export function FloorDesigner({
  initialTables = [],
}: FloorDesignerProps) {
  const [originalTables, setOriginalTables] = useState<DraggableTableItem[]>(initialTables)
  const [tables, setTables] = useState<DraggableTableItem[]>(initialTables)
  const [selectedItem, setSelectedItem] = useState<number | null>(null)
  const [isDraggingTable, setIsDraggingTable] = useState(false)
  const { mutate: updateTable} = useUpdateTable();
  const [isPendingUpdate, setIsPendingUpdate] = useState(false)

  const { selectedFloor } =  useSeatingState();
  const queryClient = useQueryClient();

  useEffect(() => {
    setOriginalTables(initialTables)
    setTables(initialTables)
  }, [initialTables])

  const moveItem = (id: number, x: number, y: number) => {
    setTables(tables.map((table) => (table.id === id ? { ...table, xPos: x, yPos: y } : table)))
  }

  function isEqualTables(a: DraggableTableItem[], b: DraggableTableItem[]) {
    if (a.length !== b.length) return false;
    return a.every(table => {
      const other = b.find(t => t.id === table.id);
      return other && JSON.stringify({ xPos: table.xPos, yPos: table.yPos }) ===
        JSON.stringify({ xPos: other.xPos, yPos: other.yPos });
    });
  }

  const revertChanges = () => setTables(initialTables)

  const saveChanges = async () => {
    // find tables that have changed compared to original
    const changedTables = tables.filter((table) => {
      const original = originalTables.find(
        (t) => t.id === table.id);
      if (!original) return true;
      return table.xPos !== original.xPos ||
        table.yPos !== original.yPos;
    });

    setIsPendingUpdate(true);

    try {
      await Promise.all(
        changedTables.map(async (t) => {
          updateTable({
            id: t.id,
            body: JSON.stringify({ x_pos: t.xPos, y_pos: t.yPos }),
          });
          toast.success(`Table [${t.id} | ${t.name}] updated successfully`);
        })
      );
    } catch (error) {
      if (error && isHTTPResponse<null>(error)) {
        if (typeof error.data === "string") {
          toast.error(error.data);
          return;
        }

        if (typeof error.data === "object" && error.data !== null) {
          const data = error.data as tableErrorResponse;
          if (data.name && data.name.length > 0) {
            toast.error(data.name[0]);
          }
        }
      }
      if (error instanceof Error) {
        toast.error(error.message);
      }
    } finally {
      setIsPendingUpdate(false);
      await queryClient.invalidateQueries({ queryKey: ["floors"] });
      await queryClient.invalidateQueries({ queryKey: ["floor.tables", selectedFloor?.id] });
    }
  }

  if (isPendingUpdate) return <AppLoading />

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

      {!isEqualTables(originalTables, tables) && (
        <UpdateActionDock
          revert={revertChanges}
          save={saveChanges}
        />
      )}

      <TableActionDelete />
    </div>
  )
}
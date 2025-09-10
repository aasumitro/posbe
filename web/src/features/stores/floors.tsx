import {FloorDesigner} from "@/features/stores/components/floor/designer";
import {useEffect, useState} from "react";
import type {DraggableTableItem} from "@/components/table";
import {useSeatingState} from "@/states/seating-state";
import {useFloorList, useFloorTableList} from "@/hooks/use-seating";
import type {Table} from "@/types/seating";

export function FloorPlanPage() {
  const {data: floors, isPending: isLoadFloors} = useFloorList()
  const {setFloors, selectedFloor, setSelectedFloor, setTables} =  useSeatingState();
  const {data: tables, isPending: isLoadTable} = useFloorTableList(selectedFloor?.id)
  const [draggableTable, setDraggableTable] = useState<DraggableTableItem[]>([])

  useEffect(() => {
    if (floors?.data) {
      setFloors(floors.data)

      if (!selectedFloor) {
        const selected = floors.data.reduce((min, f) =>
          f.id < min.id ? f : min)
        setSelectedFloor(selected)
      }
    }

    if (tables?.data) {
      setTables(tables.data)
      setDefaultLayout(tables.data);
    } else {
      setDefaultLayout([])
    }
  }, [floors?.data, tables?.data]);

  const setDefaultLayout = (tables: Table[] | null) => {
    const defaultTables: DraggableTableItem[] = tables?.map((table) => ({
      id: table.id.toString(),
      xPos: table.x_pos,
      yPos: table.y_pos,
      config: table.type === "circle"
        ? {
          shape: "circle",
          diameter: table.d_size,
          chairs: table.capacity
        }
        : {
          shape: "rectangle",
          width: table.w_size,
          height: table.h_size,
          chairs: table.capacity
        },
      customers: table.capacity,
      status: "available",
    })) ?? []

    setDraggableTable(defaultTables)
  }

  if (isLoadFloors || isLoadTable) return <>Loading . . .</>

  return <FloorDesigner initialTables={draggableTable}/>
}
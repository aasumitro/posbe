import {useEffect, useState} from "react";
import type {DraggableTableItem} from "@/components/table";
import {FloorManagement} from "@/features/orders/components/floor-mgmt";
import {useOrderState} from "@/states/order-state";
import {useFloorDetail} from "@/hooks/use-seating";
import {AppLoading} from "@/components/app-loading";
import type {Table} from "@/types/seating";

export function FloorPage() {
  const {defaultFloorId} =  useOrderState();
  const [draggableTable, setDraggableTable] = useState<DraggableTableItem[]>([])
  const {data: floor, isFetching, isSuccess} = useFloorDetail(defaultFloorId);

  useEffect(() => {
    if (floor?.data) setDefaultLayout(floor?.data?.tables || []);
  }, [floor?.data])

  const setDefaultLayout = (tables: Table[]) => {
    const defaultTables: DraggableTableItem[] = tables.map((table) => ({
      id: table.id,
      floor_id: table.floor_id,
      name: table.name,
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

  if (isFetching && !isSuccess) return <AppLoading />;

  return (
    <div className="@container/main flex flex-1 flex-col gap-2 p-2">
      <FloorManagement tables={draggableTable}  />
    </div>
  )
}
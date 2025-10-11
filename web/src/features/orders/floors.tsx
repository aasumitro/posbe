import {useCallback, useEffect, useState} from "react";
import type {DraggableTableItem} from "@/components/table";
import {FloorManagement} from "@/features/orders/components/floor-mgmt";
import {useOrderState} from "@/states/order-state";
import {useFloorDetail} from "@/hooks/use-seating";
import {AppLoading} from "@/components/app-loading";
import type {Table} from "@/types/seating";
import {type EventSourceData, useEventSource, useEventSourceListener} from "@/hooks/use-sse";
import {API_URL} from "@/lib/api";

export function FloorPage() {
  const {defaultFloorId} =  useOrderState();
  const [tables, setTables] = useState<DraggableTableItem[]>([])
  const {data: floor, isFetching, isSuccess} = useFloorDetail(defaultFloorId);

  useEffect(() => {
    if (floor?.data) setDefaultLayout(floor?.data?.tables || []);
  }, [floor?.data])

  const setDefaultLayout = useCallback((tables: Table[]) => {
    const defaultTables: DraggableTableItem[] = tables.map((table) => ({
      id: table.id,
      floor_id: table.floor_id,
      name: table.name,
      xPos: table.x_pos,
      yPos: table.y_pos,
      config: table.type === "circle"
        ? {shape: "circle", diameter: table.d_size, chairs: table.capacity}
        : {shape: "rectangle", width: table.w_size, height: table.h_size, chairs: table.capacity},
      customers: 0,
      status: "available",
    })) ?? []

    setTables(defaultTables)
  }, []);

  const [eventSource] = useEventSource(`${API_URL}/orders/events`, true);
  useEventSourceListener(eventSource, ["update"], (evt) =>
      handleEventSourceChange(evt), [handleEventSourceChange]);
  function handleEventSourceChange(evt: EventSourceData) {
    try {
      const data = JSON.parse(evt.data);
      if (data.type === "table") {
        // do this if it's a single data
        // e.g:
        // {"field": "status", "status": "occupied"}
        // setDraggableTable((prev) => prev.map((t) =>
        //     t.id === data.id ? { ...t, [data.field]: data[data.field] } : t
        // ));

        // use this for multiple data
        setTables((prev) =>
          prev.map((t) => {
            if (t.id !== data.id) return t;

            const updated = { ...t };

            data.field.forEach((f: keyof DraggableTableItem) => {
              if (f in updated && data[f] !== undefined) {
                (updated as any)[f] = data[f];
              }
            });

            return updated;
          })
        );
      }

      if (data.type === "reload") {
        window.location.reload();
      }
    } catch (error) {
      console.log("unexpected json format", error)
    }
  }

  if (isFetching && !isSuccess) return <AppLoading />;

  return (
    <div className="@container/main flex flex-1 flex-col gap-2 p-2">
      <FloorManagement tables={tables}  />
    </div>
  )
}
import {FloorDesigner} from "@/features/stores/components/floor/designer";
import {useEffect, useState} from "react";
import type {DraggableTableItem} from "@/components/table";

export function FloorPlanPage() {
  const [tables, setTables] = useState<DraggableTableItem[]>([])

  // Load saved floor plan on component mount
  useEffect(() => {
    const saved = localStorage.getItem("restaurantFloorPlan")
    if (saved) {
      try {
        const data = JSON.parse(saved)
        setTables(data.tables || [])
      } catch (error) {
        console.error("Failed to load saved floor plan:", error)
        // Set default layout if loading fails
        setDefaultLayout()
      }
    } else {
      // Set default layout if no saved data
      setDefaultLayout()
    }
  }, [])

  const setDefaultLayout = () => {
    const defaultTables: DraggableTableItem[] = [
      {id: "TA1", xPos: 0, yPos: 0, config: { shape: "rectangle", width: 100, height: 60, chairs: 6}, status: "occupied", customers: 6},
      {id: "TA2", xPos: 400, yPos: 0, config: { shape: "rectangle", width: 140, height: 60, chairs: 8}, status: "occupied", customers: 6},
      {id: "TB1", xPos: 0, yPos: 700, config: { shape: "circle", diameter: 60, chairs: 2}, status: "reserved", customers: 2},
      {id: "TB2", xPos: 0, yPos: 300, config: { shape: "circle", diameter: 60, chairs: 4}, status: "reserved", customers: 4},
      {id: "TC1", xPos: 400, yPos: 300, config: { shape: "rectangle", width: 60, height: 60, chairs: 4}, status: "needs-cleaning", customers: 3},
      {id: "TC2", xPos: 400, yPos: 650, config: { shape: "rectangle", width: 60, height: 60, chairs: 4}, status: "available", customers: 0}
    ]
    setTables(defaultTables)
  }

  return (
    <FloorDesigner initialTables={tables}/>
  )
}
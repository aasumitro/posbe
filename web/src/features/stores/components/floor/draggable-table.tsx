"use client"

import type React from "react"
import { useState, useEffect, useCallback } from "react"
import type {DraggableTableItem} from "@/components/table"
import { Table } from "@/components/table";
import {TableMenu} from "@/features/stores/components/floor/table-menu";
import {cn} from "@/lib/utils";

interface DraggableTableProps {
  item: DraggableTableItem
  onMove: (id: number, x: number, y: number) => void
  onSelect: (id: number) => void
  isSelected: boolean
  scale: number
  setDragging: (dragging: boolean) => void
}

export function DraggableTable({
  item, onMove, onSelect, isSelected, setDragging
}: DraggableTableProps) {
  const [isDragging, setIsDragging] = useState(false)
  const [dragStart, setDragStart] = useState({ x: 0, y: 0 })

  useEffect(() => {
    setDragging(isDragging)
  }, [isDragging, setDragging])

  const handleMouseMove = useCallback(
    (e: MouseEvent) => {
      if (isDragging) {
        e.preventDefault()
        const newX = e.clientX - dragStart.x
        const newY = e.clientY - dragStart.y
        onMove(item.id, newX, newY)
      }
    }, [isDragging, dragStart, onMove, item.id],
  )

  const handleMouseUp = useCallback(() => {
    setIsDragging(false)
  }, [])

  useEffect(() => {
    if (isDragging) {
      document.addEventListener("mousemove", handleMouseMove)
      document.addEventListener("mouseup", handleMouseUp)
      document.body.style.cursor = "grabbing"
      document.body.style.userSelect = "none"

      return () => {
        document.removeEventListener("mousemove", handleMouseMove)
        document.removeEventListener("mouseup", handleMouseUp)
        document.body.style.cursor = ""
        document.body.style.userSelect = ""
      }
    }
  }, [isDragging, handleMouseMove, handleMouseUp])

  const handleMouseDown = (e: React.MouseEvent) => {
    e.preventDefault()
    setIsDragging(true)
    setDragStart({
      x: e.clientX - item.xPos,
      y: e.clientY - item.yPos,
    })
    onSelect(item.id)
  }

  function getScale(shape: string, chairs: number): string {
    const circleScales: Record<number, string> = {
      1: "scale(0.8)",
      2: "scale(1.1)",
      3: "scale(0.85)",
    };
    const rectangleScales: Record<number, string> = {
      1: "scale(0.5)",
      2: "scale(0.5)",
      3: "scale(0.8)",
      4: "scale(1)",
      5: "scale(0.9)",
      6: "scale(1.2)",
      7: "scale(1.1)",
      8: "scale(1.4)",
      9: "scale(1.3)",
      12: "scale(1.8)",
    };

    if (shape === "circle") {
      return circleScales[chairs] ?? "scale(1)";
    } else {
      return rectangleScales[chairs] ?? "scale(1.5)";
    }
  }

  return (
    <div
      className={cn(
        "flex items-center justify-center",
        item.config.shape ===  "rectangle" && item.config.chairs <= 2 && "rotate-90"
      )}
      style={{
        position: "absolute",
        left: item.xPos,
        top: item.yPos,
        zIndex: isDragging ? 1000 : isSelected ? 100 : 1,
        transform: getScale(item.config.shape, item.config.chairs),
        transformOrigin: "center",
      }}
      onMouseDown={handleMouseDown}
    >
      <Table
        id={item.id}
        name={item.name}
        status={item.status}
        customers={item.customers}
        config={item.config}
        menu={<TableMenu id={item.id} shape={item.config.shape}/>}
      />
    </div>
  )
}

"use client"

import type React from "react"
import { useState, useEffect, useCallback } from "react"
import type {DraggableTableItem} from "@/components/table"
import { Table } from "@/components/table";
import {TableMenu} from "@/features/stores/components/floor/table-menu";
import {cn} from "@/lib/utils";
import {getTableScale} from "@/lib/table";

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
        transform: getTableScale(item.config.shape, item.config.chairs),
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
        menu={<TableMenu
          id={item.id}
          fid={item.floor_id}
          shape={item.config.shape}
          chairs={item.config.chairs}
        />}
      />
    </div>
  )
}

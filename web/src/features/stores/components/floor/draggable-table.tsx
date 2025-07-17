"use client"

import type React from "react"
import { useState, useEffect, useCallback } from "react"
import type {DraggableTableItem} from "@/components/table"
import { Table } from "@/components/table";
import {TableMenu} from "@/features/stores/components/floor/table-menu";

interface DraggableTableProps {
  item: DraggableTableItem
  onMove: (id: string, x: number, y: number) => void
  onSelect: (id: string) => void
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
      style={{
        position: "absolute",
        left: item.xPos,
        top: item.yPos,
        zIndex: isDragging ? 1000 : isSelected ? 100 : 1,
      }}
      onMouseDown={handleMouseDown}
    >
      <Table
        name={item.id}
        status={item.status}
        customers={item.customers}
        config={item.config}
        menu={
          <TableMenu shape={item.config.shape} />
        }
      />
    </div>
  )
}

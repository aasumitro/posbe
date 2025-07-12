"use client"

import { useRef, useState, useCallback } from "react"
import type { ReactNode, MouseEvent, WheelEvent } from "react"
import { Button } from "@/components/ui/button"
import { ZoomIn, ZoomOut, RotateCcw, InfoIcon } from "lucide-react"
import { cn } from "@/lib/utils";

interface ZoomableContainerProps {
  children: ReactNode
  className?: string
}

export function ZoomableContainer({ children, className }: ZoomableContainerProps) {
  const [zoom, setZoom] = useState(0.8)
  const [pan, setPan] = useState({ x: 0, y: 0 })
  const [isDragging, setIsDragging] = useState(false)
  const [dragStart, setDragStart] = useState({ x: 0, y: 0 })
  const containerRef = useRef<HTMLDivElement>(null)
  const contentRef = useRef<HTMLDivElement>(null)

  const handleZoomIn = useCallback(() => {
    setZoom((prev) => Math.min(prev * 1.2, 3))
  }, [])

  const handleZoomOut = useCallback(() => {
    setZoom((prev) => Math.max(prev / 1.2, 0.3))
  }, [])

  const handleReset = useCallback(() => {
    setZoom(1)
    setPan({ x: 0, y: 0 })
  }, [])

  const handleMouseDown = useCallback(
    (e: MouseEvent) => {
      // Only start dragging if clicking on the container background, not on cards or buttons
      const target = e.target as HTMLElement
      if (
        target === containerRef.current ||
        target === contentRef.current ||
        target.closest('[data-draggable="true"]')
      ) {
        e.preventDefault()
        setIsDragging(true)
        setDragStart({ x: e.clientX - pan.x, y: e.clientY - pan.y })
      }
    }, [pan],
  )

  const handleMouseMove = useCallback(
    (e: MouseEvent) => {
      if (isDragging) {
        if (e.currentTarget.classList.contains("no-pan")) return;
        e.preventDefault()
        setPan({
          x: e.clientX - dragStart.x,
          y: e.clientY - dragStart.y,
        })
      }
    }, [isDragging, dragStart],
  )

  const handleMouseUp = useCallback(() => {
    setIsDragging(false)
  }, [])

  const handleWheel = useCallback((e: WheelEvent) => {
    const delta = e.deltaY > 0 ? 0.9 : 1.1
    setZoom((prev) => Math.max(0.3, Math.min(3, prev * delta)))
  }, [])

  return (
    <div className="relative w-full h-full overflow-hidden bg-card">
      {/* Controls */}
      <div className="absolute bottom-4 left-4 z-20 flex gap-2 opacity-65 hover:opacity-100">
        <Button onClick={handleZoomIn} size="sm" variant="outline" className="cursor-pointer">
          <ZoomIn className="w-4 h-4" />
        </Button>
        <Button onClick={handleZoomOut} size="sm" variant="outline" className="cursor-pointer">
          <ZoomOut className="w-4 h-4" />
        </Button>
        <Button onClick={handleReset} size="sm" variant="outline" className="cursor-pointer">
          <RotateCcw className="w-4 h-4" />
        </Button>
      </div>

      {/* Zoom indicator */}
      <div className="absolute top-4 right-4 z-20 flex gap-2 select-none">
        <div className="flex items-center gap-2 px-3 py-1 rounded-sm border text-sm cursor-auto group bg-gray-100">
          <InfoIcon className="w-4 h-4" />
          <span className="hidden group-hover:block">
            Drag to move & scroll to zoom in/out
          </span>
        </div>

        <div className="flex items-center gap-2 px-3 py-1 rounded-sm border text-sm cursor-auto bg-gray-100">
          {Math.round(zoom * 100)}%
        </div>
      </div>

      {/* Zoomable content */}
      <div
        ref={containerRef}
        onMouseDown={handleMouseDown}
        onMouseMove={handleMouseMove}
        onMouseUp={handleMouseUp}
        onMouseLeave={handleMouseUp}
        onWheel={handleWheel}
        data-draggable="true"
        className={cn(
          "w-full h-full overflow-hidden",
          isDragging ? "cursor-grabbing" : "cursor-grab",
          className
        )}
        style={{
          backgroundImage: "radial-gradient(circle, rgba(148, 163, 184, 0.3) 1px, transparent 1px)",
          backgroundSize: `${20 * zoom}px ${20 * zoom}px`,
          backgroundPosition: `${pan.x}px ${pan.y}px`,
        }}
      >
        <div
          ref={contentRef}
          data-draggable="true"
          className="origin-center transition-transform duration-100 ease-out"
          style={{
            transform: `translate(${pan.x}px, ${pan.y}px) scale(${zoom})`,
            minWidth: "100%",
            minHeight: "100%",
          }}
        >
          <div
            data-draggable="true"
            className="flex justify-center h-screen items-start pt-8"
          >
            {children}
          </div>
        </div>
      </div>
    </div>
  )
}
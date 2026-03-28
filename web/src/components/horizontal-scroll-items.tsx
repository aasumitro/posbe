"use client"

import type {MouseEvent, ReactNode} from "react";
import { useRef} from "react";

export function HorizontalScrollItems({children}: { children: ReactNode }) {
  const scrollRef = useRef<HTMLDivElement>(null);
  const isDragging = useRef(false);
  const startX = useRef(0);
  const scrollLeft = useRef(0);

  const handleMouseDown = (e: MouseEvent) => {
    isDragging.current = true;
    startX.current = e.pageX - (scrollRef.current?.offsetLeft || 0);
    scrollLeft.current = scrollRef.current?.scrollLeft || 0;
  };

  const handleMouseLeaveOrUp = () => {
    isDragging.current = false;
  };

  const handleMouseMove = (e: MouseEvent) => {
    if (!isDragging.current || !scrollRef.current) return;
    e.preventDefault();
    const x = e.pageX - scrollRef.current.offsetLeft;
    const walk = (x - startX.current); // drag speed
    scrollRef.current.scrollLeft = scrollLeft.current - walk;
  };

  return (
    <div
      className="max-w-full overflow-x-auto cursor-grab active:cursor-grabbing min-h-14 hide-scrollbar"
      ref={scrollRef}
      onMouseDown={handleMouseDown}
      onMouseUp={handleMouseLeaveOrUp}
      onMouseLeave={handleMouseLeaveOrUp}
      onMouseMove={handleMouseMove}
    >
      {children}
    </div>
  )
}
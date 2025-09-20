"use client"

import type { JSX, ReactNode} from "react"

export type TableStatus = "available" | "occupied" | "reserved" | "needs-cleaning"

export interface BaseTableConfig {
  chairs: number
}

export interface RectangleTableConfig extends BaseTableConfig {
  shape: "rectangle"
  width: number
  height: number
}

export interface CircleTableConfig extends BaseTableConfig {
  shape: "circle"
  diameter: number
}

export interface DraggableTableItem {
  id: string
  xPos: number
  yPos: number
  config: RectangleTableConfig | CircleTableConfig
  status: TableStatus
  customers?: number
}

export interface ChairPosition {
  x: number
  y: number
  rotation: number
}

/**
 * Component for displaying a visual preview of the table
 */
export interface TableProps {
  name: string
  status: TableStatus
  customers?: number
  config: CircleTableConfig | RectangleTableConfig
  menu?: ReactNode
}

function getChairPositions(
  config: RectangleTableConfig | CircleTableConfig
): ChairPosition[] {
  const positions: ChairPosition[] = []

  if (config.chairs === 0) return positions

  if (config.shape === "rectangle") {
    const halfWidth = config.width / 2
    const halfHeight = config.height / 2

    if (config.chairs === 1) {
      // Single chair - place on the longer side, centered
      if (config.width >= config.height) {
        positions.push({ x: 0, y: -halfHeight - 25, rotation: 0 })
      } else {
        positions.push({ x: halfWidth + 25, y: 0, rotation: 90 })
      }
    } else if (config.chairs === 2) {
      // Two chairs - place on opposite sides
      if (config.width >= config.height) {
        // Place on top and bottom
        positions.push({ x: 0, y: -halfHeight - 25, rotation: 0 })
        positions.push({ x: 0, y: halfHeight + 25, rotation: 180 })
      } else {
        // Place on left and right
        positions.push({ x: -halfWidth - 25, y: 0, rotation: 270 })
        positions.push({ x: halfWidth + 25, y: 0, rotation: 90 })
      }
    } else {
      // 3+ chairs - use proportional distribution
      const totalSides = 2 * (config.width + config.height)
      const topChairs = Math.max(0, Math.round((config.width / totalSides) * config.chairs))
      const rightChairs = Math.max(0, Math.round((config.height / totalSides) * config.chairs))
      const bottomChairs = Math.max(0, Math.round((config.width / totalSides) * config.chairs))
      const leftChairs = Math.max(0, config.chairs - topChairs - rightChairs - bottomChairs)

      // Ensure at least one chair if we have remaining chairs
      if (topChairs === 0 && rightChairs === 0 && bottomChairs === 0 && leftChairs === 0 && config.chairs > 0) {
        // Fallback: distribute evenly starting with longer sides
        const chairsPerSide = Math.floor(config.chairs / 4)
        const remainder = config.chairs % 4

        const adjustedTopChairs = chairsPerSide + (remainder > 0 ? 1 : 0)
        const adjustedRightChairs = chairsPerSide + (remainder > 1 ? 1 : 0)
        const adjustedBottomChairs = chairsPerSide + (remainder > 2 ? 1 : 0)
        const adjustedLeftChairs = chairsPerSide

        // Use the adjusted values
        if (adjustedTopChairs > 0) {
          const spacing = config.width / adjustedTopChairs
          for (let i = 0; i < adjustedTopChairs; i++) {
            const x = -halfWidth + spacing * (i + 0.5)
            const y = -halfHeight - 25
            positions.push({ x, y, rotation: 0 })
          }
        }

        if (adjustedRightChairs > 0) {
          const spacing = config.height / adjustedRightChairs
          for (let i = 0; i < adjustedRightChairs; i++) {
            const x = halfWidth + 25
            const y = -halfHeight + spacing * (i + 0.5)
            positions.push({ x, y, rotation: 90 })
          }
        }

        if (adjustedBottomChairs > 0) {
          const spacing = config.width / adjustedBottomChairs
          for (let i = 0; i < adjustedBottomChairs; i++) {
            const x = halfWidth - spacing * (i + 0.5)
            const y = halfHeight + 25
            positions.push({ x, y, rotation: 180 })
          }
        }

        if (adjustedLeftChairs > 0) {
          const spacing = config.height / adjustedLeftChairs
          for (let i = 0; i < adjustedLeftChairs; i++) {
            const x = -halfWidth - 25
            const y = halfHeight - spacing * (i + 0.5)
            positions.push({ x, y, rotation: 270 })
          }
        }
      } else {
        // Original proportional logic
        // Top edge
        if (topChairs > 0) {
          const spacing = config.width / topChairs
          for (let i = 0; i < topChairs; i++) {
            const x = -halfWidth + spacing * (i + 0.5)
            const y = -halfHeight - 25
            positions.push({ x, y, rotation: 0 })
          }
        }

        // Right edge
        if (rightChairs > 0) {
          const spacing = config.height / rightChairs
          for (let i = 0; i < rightChairs; i++) {
            const x = halfWidth + 25
            const y = -halfHeight + spacing * (i + 0.5)
            positions.push({ x, y, rotation: 90 })
          }
        }

        // Bottom edge
        if (bottomChairs > 0) {
          const spacing = config.width / bottomChairs
          for (let i = 0; i < bottomChairs; i++) {
            const x = halfWidth - spacing * (i + 0.5)
            const y = halfHeight + 25
            positions.push({ x, y, rotation: 180 })
          }
        }

        // Left edge
        if (leftChairs > 0) {
          const spacing = config.height / leftChairs
          for (let i = 0; i < leftChairs; i++) {
            const x = -halfWidth - 25
            const y = halfHeight - spacing * (i + 0.5)
            positions.push({ x, y, rotation: 270 })
          }
        }
      }
    }
  } else {
    // Round table
    const radius = config.diameter / 2
    const chairRadius = radius + 25

    for (let i = 0; i < config.chairs; i++) {
      const angle = (i * 2 * Math.PI) / config.chairs
      const x = Math.cos(angle) * chairRadius
      const y = Math.sin(angle) * chairRadius
      const rotation = (angle * 180) / Math.PI + 90

      positions.push({ x, y, rotation })
    }
  }

  return positions
}

/**
 * Calculates the SVG viewBox based on table and chair positions
 * @returns SVG viewBox string
 */
function calculateViewBox(
  config: RectangleTableConfig | CircleTableConfig,
  chairPositions: ChairPosition[]
): string {
  let minX: number
  let minY: number
  let maxX: number
  let maxY: number

  const CHAIR_WIDTH = 20
  const CHAIR_HEIGHT = 30
  const CHAIR_PADDING = 6
  const BOUND_PADDING = 0

  if (config.shape === "rectangle") {
    const width = config.width ?? 0
    const height = config.height ?? 0
    minX = -width / 2 - CHAIR_PADDING
    minY = -height / 2 - CHAIR_PADDING
    maxX = width / 2 + CHAIR_PADDING
    maxY = height / 2 + CHAIR_PADDING
  } else {
    const diameter = config.diameter ?? 0
    const radius = diameter / 2 + CHAIR_PADDING
    minX = -radius
    minY = -radius
    maxX = radius
    maxY = radius
  }

  const chairOffset = Math.max(CHAIR_WIDTH, CHAIR_HEIGHT) + CHAIR_PADDING

  for (const chair of chairPositions) {
    minX = Math.min(minX, chair.x - chairOffset)
    minY = Math.min(minY, chair.y - chairOffset)
    maxX = Math.max(maxX, chair.x + chairOffset)
    maxY = Math.max(maxY, chair.y + chairOffset)
  }

  minX -= BOUND_PADDING
  minY -= BOUND_PADDING
  maxX += BOUND_PADDING
  maxY += BOUND_PADDING

  const width: number = maxX - minX
  const height: number = maxY - minY
  return `${minX} ${minY} ${width} ${height}`
}

export function Table({
  customers, config, menu
}: TableProps) {
  const chairPositions = getChairPositions(config)

  const viewBox: string = calculateViewBox(config, chairPositions)

  /**
   * Renders a rectangular table
   */
  const renderRectangleTable = (config: RectangleTableConfig): JSX.Element => {
    return (
      <>
        {/* Table shadow */}
        <rect
          x={-config.width / 2 - 5}
          y={-config.height / 2 - 5}
          width={config.width + 10}
          height={config.height + 10}
          fill="rgba(0,0,0,0.2)"
          rx="10"
        />
        {/* Table top */}
        <rect
          x={-config.width / 2}
          y={-config.height / 2}
          width={config.width}
          height={config.height}
          fill="url(#woodPattern)"
          stroke="#654321"
          strokeWidth="3"
          rx="8"
        />
        {/* Table edge highlight */}
        <rect
          x={-config.width / 2 + 5}
          y={-config.height / 2 + 5}
          width={config.width - 10}
          height={config.height - 10}
          fill="none"
          stroke="#D2B48C"
          strokeWidth="1"
          rx="4"
          opacity="0.5"
        />
      </>
    )
  }

  /**
   * Renders a round table
   */
  const renderRoundTable = (config: CircleTableConfig): JSX.Element => {
    return (
      <>
        {/* Round table shadow */}
        <circle cx="0" cy="0" r={config.diameter / 2 + 5} fill="rgba(0,0,0,0.2)" />
        {/* Round table top */}
        <circle cx="0" cy="0" r={config.diameter / 2} fill="url(#woodPattern)" stroke="#654321" strokeWidth="3" />
        {/* Round table edge highlight */}
        <circle cx="0" cy="0" r={config.diameter / 2 - 10} fill="none" stroke="#8D6E63" strokeWidth="1" opacity="0.5" />
      </>
    )
  }

  const renderMenu = (config: CircleTableConfig | RectangleTableConfig,  menu?: ReactNode): JSX.Element => {
    if (config.shape === "circle") {
      return (
        <>
          <defs>
            <mask id="circleMenuMask">
              <circle
                cx="0"
                cy="0"
                r={config.diameter / 2 - 10}
                fill="white"
              />
            </mask>
          </defs>

          <foreignObject
            x={-(config.diameter / 2 - 10)}
            y={-(config.diameter / 2 - 10)}
            width={config.diameter - 20}
            height={config.diameter - 20}
            mask="url(#circleMenuMask)"
          >
            {menu}
          </foreignObject>
        </>
      )
    }
    return  (
      <foreignObject
        x={-config.width / 4}
        y={-config.height / 4}
        width={config.width / 2}
        height={config.height / 2}
      >
        {menu}
      </foreignObject>
    )
  }

  /**
   * Renders a chair at the specified position
   */
  const renderChair = (pos: ChairPosition, index: number, fillColor: string): JSX.Element => {
    return (
      <g key={index} transform={`translate(${pos.x}, ${pos.y}) rotate(${pos.rotation})`}>
        {/* Chair shadow */}
        <rect x={-12} y={-18} width="24" height="36" fill="rgba(0,0,0,0.2)" rx="4" />
        {/* Chair seat */}
        <rect x={-10} y={-15} width="20" height="30" fill={fillColor} stroke="#3E2723" strokeWidth="2" rx="3" />
        {/* Chair back */}
        <rect x={-8} y={-18} width="16" height="6" fill={fillColor} stroke="#3E2723" strokeWidth="2" rx="2" />
        {/* Chair highlight */}
        <rect x={-6} y={-10} width="12" height="20" fill="none" stroke="#8D6E63" strokeWidth="1" rx="1" opacity="0.5" />
      </g>
    )
  }

  const getChairColor = (isOccupied: boolean) => {
    return isOccupied ? "#593d35" : "#7a411c"
  }

  return (
    <div className="relative min-w-72">
      <svg width="100%" height="100%" viewBox={viewBox} className="max-w-full max-h-full">
        {/* Table */}
        <defs>
          <pattern
            id="woodPattern"
            patternUnits="userSpaceOnUse"
            width="30"
            height="30"
            patternTransform="rotate(45)"
          >
            <rect width="30" height="30" fill="#A0522D" />
            <rect width="15" height="30" fill="#8B4513" />
          </pattern>
        </defs>

        <g transform="translate(0,0)">
          {config.shape === "rectangle"
            ? renderRectangleTable(config)
            : renderRoundTable(config)}

          {/* Define a circular clipPath */}
          {renderMenu(config, menu)}
        </g>

        {/* Chairs */}
        {chairPositions.map((pos, index) =>
          renderChair(pos, index, getChairColor(index < (customers ?? 0))))}
      </svg>
    </div>
  )
}

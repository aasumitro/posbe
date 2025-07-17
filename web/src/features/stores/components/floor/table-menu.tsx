import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { cn } from "@/lib/utils"
import {BorderBeam} from "@/components/border-beam";
import {IconAdjustmentsCog} from "@tabler/icons-react";

export function TableMenu({
  shape
}: {
  shape: string
}) {
  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <div
          className={cn(
            "relative",
            "absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2",
            "w-11 h-11 cursor-pointer"
          )}
        >
          {/* Outer border beam */}
          <div className={cn(
            "absolute -inset-1 z-0",
            shape === "circle" ? "rounded-full": "rounded-lg "
          )}>
            <BorderBeam
              duration={5}
              size={25}
              delay={8}
              className={cn(
                "w-full h-full rounded-full pointer-events-none",
                "from-transparent via-white to-transparent"
              )}
            />
          </div>

          {/* Badge content (above beam) */}
          <div
            className={cn(
              "relative z-10 text-xs font-bold", // 👈 ensures it's above the beam
              "flex items-center justify-center text-center p-2 w-11 h-11",
              "bg-gray-50/10 hover:bg-gray-100/50",
              shape === "circle" ? "rounded-full": "rounded-lg "
            )}
          >
            <IconAdjustmentsCog className="text-white" />
          </div>
        </div>
      </DropdownMenuTrigger>
      <DropdownMenuContent className="w-56" align="start">
        <DropdownMenuGroup>
          <DropdownMenuItem>
            Edit
          </DropdownMenuItem>
          <DropdownMenuItem variant="destructive">
            Delete
          </DropdownMenuItem>
        </DropdownMenuGroup>
      </DropdownMenuContent>
    </DropdownMenu>
  )
}

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { cn } from "@/lib/utils"
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
            "relative cursor-pointer",
            "absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2",
          )}
        >
          <div
            className={cn(
              "relative z-10 text-xs font-bold",
              "flex items-center justify-center text-center p-2 w-11 h-11",
              "bg-gray-50/10 hover:bg-gray-100/50",
              shape === "circle" ? "rounded-full": "rounded-lg "
            )}
          >
            <IconAdjustmentsCog className="w-4 h-4 text-white" />
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

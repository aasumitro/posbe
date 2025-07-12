import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuPortal,
  DropdownMenuSeparator,
  DropdownMenuSub,
  DropdownMenuSubContent,
  DropdownMenuSubTrigger,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { cn } from "@/lib/utils"
import {BorderBeam} from "@/components/border-beam";
import type { TableStatus } from "@/components/table";
import { getLabelBgColorByStatus } from "@/lib/table";

export function TableMenu({
  name, status, shape
}: {
  name:string,
  status: string,
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
              status === "available" && "text-white",
              shape === "circle" ? "rounded-full": "rounded-lg "
            )}
            style={{ background: getLabelBgColorByStatus(status as TableStatus) }}
          >
            {name}
          </div>
        </div>
      </DropdownMenuTrigger>
      <DropdownMenuContent className="w-56" align="start">
        <DropdownMenuGroup>
          {status === "available" && (
            <DropdownMenuItem>
              Place an order
            </DropdownMenuItem>
          )}
          {status !== "available" && status !== "needs-cleaning" && (
            <>
              <DropdownMenuItem>
                Update order
              </DropdownMenuItem>
              <DropdownMenuItem>
                Complete
              </DropdownMenuItem>
            </>
          )}
          {status !== "needs-cleaning" && (
            <DropdownMenuSeparator />
          )}
          <DropdownMenuSub>
            <DropdownMenuSubTrigger>Set status</DropdownMenuSubTrigger>
            <DropdownMenuPortal>
              <DropdownMenuSubContent>
                <DropdownMenuItem>Occupied</DropdownMenuItem>
                <DropdownMenuItem>Reserved</DropdownMenuItem>
                {status !== "needs-cleaning" && (
                  <DropdownMenuItem>Need Cleaning</DropdownMenuItem>
                )}
                <DropdownMenuItem>Available</DropdownMenuItem>
              </DropdownMenuSubContent>
            </DropdownMenuPortal>
          </DropdownMenuSub>
        </DropdownMenuGroup>
      </DropdownMenuContent>
    </DropdownMenu>
  )
}

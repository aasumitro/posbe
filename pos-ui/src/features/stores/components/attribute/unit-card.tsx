import {
  DropdownMenu,
  DropdownMenuContent, DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger
} from "@/components/ui/dropdown-menu";
import {Button} from "@/components/ui/button";
import {IconDotsVertical} from "@tabler/icons-react";
import type {Unit} from "@/types/unit";

export function UnitCard({
  magnitude, symbol, name
}: Unit) {
  return (
    <div className="bg-muted/50 aspect-video rounded-xl relative flex items-center justify-center">
      <div className="flex items-center justify-center select-none">
        <span className="text-sm font-normal text-muted-foreground mr-2">
          #{magnitude}
        </span>
        <div className="flex items-baseline gap-2 font-bold tabular-nums leading-none">
          <span className="text-8xl">{symbol}</span>
          <span className="text-sm font-normal text-muted-foreground">/ {name}</span>
        </div>
      </div>

      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button
            aria-haspopup="true"
            size="icon"
            variant="ghost"
            className="absolute top-5 right-5"
          >
            <IconDotsVertical className="h-4 w-4"/>
            <span className="sr-only">Toggle menu</span>
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="end">
          <DropdownMenuLabel>Actions</DropdownMenuLabel>
          <DropdownMenuSeparator/>
          <DropdownMenuItem>Edit</DropdownMenuItem>
          <DropdownMenuItem variant="destructive">Delete</DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
  )
}
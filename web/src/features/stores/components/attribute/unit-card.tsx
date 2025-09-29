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
import {useActionState} from "@/states/action-state";
import {UnitActionDeleteModalState} from "@/features/stores/components/attribute/unit-action-delete";
import {useAttributeState} from "@/states/attribute-state";
import {UnitActionEditModalState} from "@/features/stores/components/attribute/unit-action-edit";

export function UnitCard({
  id, magnitude, symbol, name
}: Unit) {
  const { setBoolState } = useActionState();
  const { setSelectedUnit } =  useAttributeState();

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
            className="absolute top-5 right-5 cursor-pointer"
          >
            <IconDotsVertical className="h-4 w-4"/>
            <span className="sr-only">Toggle menu</span>
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="end">
          <DropdownMenuLabel>Actions</DropdownMenuLabel>
          <DropdownMenuSeparator/>
          <DropdownMenuItem
            className="cursor-pointer"
            onClick={(e) => {
              e.preventDefault();
              setSelectedUnit({id, name, magnitude, symbol});
              setBoolState(UnitActionEditModalState, true);
            }}
          >Edit</DropdownMenuItem>
          <DropdownMenuItem
            className="cursor-pointer"
            variant="destructive"
            onClick={(e) => {
              e.preventDefault();
              setSelectedUnit({id, name, magnitude, symbol});
              setBoolState(UnitActionDeleteModalState, true);
            }}
          >Delete</DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
  )
}
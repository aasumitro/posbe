import {Button} from "@/components/ui/button";
import {formatShortNumber} from "@/lib/numbers";
import {Tooltip, TooltipContent, TooltipTrigger} from "@/components/ui/tooltip";
import {IconDotsVertical, IconPuzzle} from "@tabler/icons-react";
import {
  DropdownMenu,
  DropdownMenuContent, DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger
} from "@/components/ui/dropdown-menu";
import {Badge} from "@/components/ui/badge";
import {useActionState} from "@/states/action-state";
import {AddonActionEditModalState} from "@/features/stores/components/catalog/addon-action-edit";
import {AddonActionDeleteModalState} from "@/features/stores/components/catalog/addon-action-delete";

const addons = [
  { title: "oat milk", description: "replace normal milk with oat milk", price: 12000},
  { title: "almond milk", description: "replace normal milk with almond milk", price: 8000},
  { title: "cheese", description: "add extra cheese", price: 7500},
  { title: "chocolate", description: "add extra chocolate", price: 17500},
  { title: "rice", description: "add extra rice", price: 5000},
]

export function AddonContainer() {
  const { setBoolState } = useActionState();

  return (
    <div className="grid gap-4 grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4 auto-rows-fr mt-4">
      {addons.map((addon, index) => (
        <div
          key={index}
          className="border-1 rounded-xl p-4 space-y-6 select-none flex flex-col h-full"
        >
          <div className="flex-grow space-y-6">
            <div className="flex gap-4">
              <div className="rounded-lg w-20 h-20 bg-gray-50/50 flex items-center justify-center">
                <IconPuzzle />
              </div>

              <div className="space-y-1">
                <div className="flex items-center gap-2">
                  <h5 className="text-xl font-semibold capitalize">
                    {addon.title}
                  </h5>
                  <Tooltip>
                    <TooltipTrigger asChild>
                      <Badge
                        className="h-5 min-w-5 rounded-full px-1 font-mono tabular-nums"
                        variant="outline"
                      >
                        20+
                      </Badge>
                    </TooltipTrigger>
                    <TooltipContent>
                      <p>20 orders are linked.</p>
                    </TooltipContent>
                  </Tooltip>
                </div>

                <Tooltip>
                  <TooltipTrigger asChild>
                    <p className="text-xs text-muted-foreground truncate whitespace-nowrap overflow-hidden text-ellipsis">
                      {(() => {
                        const desc = addon.description.trim();
                        const formatted = desc.charAt(0).toUpperCase() + desc.slice(1);
                        return formatted.endsWith('.') ? formatted : formatted + '.';
                      })()}
                    </p>
                  </TooltipTrigger>
                  <TooltipContent>
                    <p>
                      {addon.description}
                    </p>
                  </TooltipContent>
                </Tooltip>
                <h5 className="text-lg font-semibold text-green-500">
                  IDR {formatShortNumber(addon.price)}
                </h5>
              </div>

              <DropdownMenu>
                <DropdownMenuTrigger asChild>
                  <Button className="ml-auto my-auto" variant="ghost">
                    <IconDotsVertical />
                  </Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent align="end">
                  <DropdownMenuLabel>Actions</DropdownMenuLabel>
                  <DropdownMenuSeparator/>
                  <DropdownMenuItem onClick={(e) => {
                    e.preventDefault();
                    setBoolState(AddonActionEditModalState, true);
                  }}>Edit</DropdownMenuItem>
                  <DropdownMenuItem variant="destructive" onClick={(e) => {
                    e.preventDefault();
                    setBoolState(AddonActionDeleteModalState, true);
                  }}>Delete</DropdownMenuItem>
                </DropdownMenuContent>
              </DropdownMenu>
            </div>
          </div>
        </div>
      ))}
    </div>
  )
}
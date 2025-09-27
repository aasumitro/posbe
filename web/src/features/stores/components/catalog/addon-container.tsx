import {Button} from "@/components/ui/button";
import {formatShortNumber} from "@/lib/numbers";
import {Tooltip, TooltipContent, TooltipTrigger} from "@/components/ui/tooltip";
import {IconBoxOff, IconDotsVertical, IconPuzzle} from "@tabler/icons-react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger
} from "@/components/ui/dropdown-menu";
import {Badge} from "@/components/ui/badge";
import {useActionState} from "@/states/action-state";
import {AddonActionEditModalState} from "@/features/stores/components/catalog/addon-action-edit";
import {AddonActionDeleteModalState} from "@/features/stores/components/catalog/addon-action-delete";
import {useProductState} from "@/states/product-state";
import {useStoreState} from "@/states/store-state";

interface AddonContainerProps {
  sort?: Record<string, "asc" | "desc">;
}

export function AddonContainer({sort}: AddonContainerProps) {
  const { setBoolState } = useActionState();
  const {settings} = useStoreState();
  const {addons, setSelectedAddon} = useProductState();

  let sortedAddons = [...(addons ?? [])];

  if (sort) {
    const [field, order] = Object.entries(sort)[0] ?? [];
    if (field && order) {
      sortedAddons.sort((a, b) => {
        const aValue = a[field as keyof typeof a];
        const bValue = b[field as keyof typeof b];

        if (typeof aValue === "string" && typeof bValue === "string") {
          return order === "asc"
            ? aValue.localeCompare(bValue)
            : bValue.localeCompare(aValue);
        }

        if (typeof aValue === "number" && typeof bValue === "number") {
          return order === "asc" ? aValue - bValue : bValue - aValue;
        }

        return 0;
      });
    }
  }

  if (sortedAddons.length < 1) {
    return (
      <div className="text-center py-12">
        <IconBoxOff className="my-8 w-24 h-24 mx-auto"/>
        <h4 className="text-primary text-xl font-bold tracking-tight">
          No Addons Found
        </h4>
        <p className="text-secondary-foreground text-md font-normal">
          No addons exist yet. Add your first addon!
        </p>
      </div>
    )
  }


  return (
    <div className="grid gap-4 grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4 auto-rows-fr mt-4">
      {sortedAddons?.map((addon, index) => (
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
                    {addon?.name}
                  </h5>
                  <Tooltip>
                    <TooltipTrigger asChild>
                      <Badge
                        className="h-5 min-w-5 rounded-full px-1 font-mono tabular-nums"
                        variant="outline"
                      >
                        {addon?.usage}
                      </Badge>
                    </TooltipTrigger>
                    <TooltipContent>
                      <p>{addon?.usage} orders are linked.</p>
                    </TooltipContent>
                  </Tooltip>
                </div>

                <Tooltip>
                  <TooltipTrigger asChild>
                    <p className="text-xs text-muted-foreground truncate whitespace-nowrap overflow-hidden text-ellipsis">
                      {(() => {
                        const desc = addon?.description?.trim();
                        const formatted = desc?.charAt(0).toUpperCase() + desc?.slice(1);
                        return formatted?.endsWith('.') ? formatted : formatted + '.';
                      })()}
                    </p>
                  </TooltipTrigger>
                  <TooltipContent>
                    <p>
                      {addon?.description}
                    </p>
                  </TooltipContent>
                </Tooltip>
                <h5 className="text-lg font-semibold text-green-500">
                  {settings?.currency} {formatShortNumber(addon?.price)}
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
                    setSelectedAddon(addon);
                    setBoolState(AddonActionEditModalState, true);
                  }}>Edit</DropdownMenuItem>
                  <DropdownMenuItem variant="destructive" onClick={(e) => {
                    e.preventDefault();
                    setSelectedAddon(addon);
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
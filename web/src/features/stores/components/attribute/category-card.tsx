import type {Category} from "@/types/category";
import {
  DropdownMenu,
  DropdownMenuContent, DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger
} from "@/components/ui/dropdown-menu";
import {Button} from "@/components/ui/button";
import {IconDotsVertical} from "@tabler/icons-react";
import { Badge } from "@/components/ui/badge";
import {Tooltip, TooltipContent, TooltipTrigger} from "@/components/ui/tooltip";
import {cn} from "@/lib/utils";
import {useActionState} from "@/states/action-state";
import {useAttributeState} from "@/states/attribute-state";
import {CategoryActionDeleteModalState} from "@/features/stores/components/attribute/category-action-delete";
import {CategoryActionEditModalState} from "@/features/stores/components/attribute/category-action-edit";

export function CategoryCard({
 id, name, usage, subcategories, className
}: Category & {className?: string}) {
  const { setBoolState } = useActionState();
  const { setSelectedCategory } =  useAttributeState();

  return (
    <div className={cn("bg-muted/50 aspect-video rounded-xl relative px-6 py-4 select-none", className)}>
      <div className="flex items-center justify-between">
       <div className="flex items-center gap-2">
         <h5 className="text-lg font-mono uppercase">{name}</h5>
         <Tooltip>
           <TooltipTrigger asChild>
             <Badge
               className="h-5 min-w-5 rounded-full px-1 font-mono tabular-nums"
               variant="outline"
             >
               {usage ?? "0"}
             </Badge>
           </TooltipTrigger>
           <TooltipContent>
             <p>{usage ?? "no"} products are linked.</p>
           </TooltipContent>
         </Tooltip>
       </div>

        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button
              aria-haspopup="true"
              size="icon"
              variant="ghost"
              className="cursor-pointer"
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
                setSelectedCategory({id, name, usage, subcategories});
                setBoolState(CategoryActionEditModalState, true);
              }}
            >Edit</DropdownMenuItem>
            <DropdownMenuItem
              className="cursor-pointer"
              variant="destructive"
              onClick={(e) => {
                e.preventDefault();
                setSelectedCategory({id, name, usage});
                setBoolState(CategoryActionDeleteModalState, true);
              }}
            >Delete</DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </div>

      <section className="p-4 border-1 border-dashed mt-2 rounded-lg space-y-4 h-full">
        <p className="text-xs">Subcategories</p>

        {subcategories?.slice(0, 2).map((subcategory) => (
          <div className="flex items-center gap-2" key={subcategory.id}>
            <p className="text-sm text-muted-foreground">– {subcategory.name}</p>
            <Tooltip>
              <TooltipTrigger asChild>
                <Badge
                  className="h-5 min-w-5 rounded-full px-1 font-mono tabular-nums"
                  variant="outline"
                >
                  {subcategory.usage ?? "0"}
                </Badge>
              </TooltipTrigger>
              <TooltipContent>
                <p>{subcategory.usage ?? "no"} products are linked.</p>
              </TooltipContent>
            </Tooltip>
          </div>
        ))}

        {subcategories && subcategories.length > 2 && (
          <p className="text-xs text-muted-foreground">
            +{subcategories.length - 2} more
          </p>
        )}

        {(!subcategories || subcategories.length === 0) && (
          <p className="text-xs text-muted-foreground">
            No subcategory
          </p>
        )}
      </section>
    </div>
  )
}
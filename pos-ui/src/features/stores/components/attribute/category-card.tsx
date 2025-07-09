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

export function CategoryCard({
 name, subcategories
}: Category) {
  return (
    <div className="bg-muted/50 aspect-video rounded-xl relative px-6 py-4 select-none">
      <div className="flex items-center justify-between">
       <div className="flex items-center gap-2">
         <h5 className="text-lg font-mono uppercase">{name}</h5>
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
             <p>90 products are linked.</p>
           </TooltipContent>
         </Tooltip>
       </div>

        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button
              aria-haspopup="true"
              size="icon"
              variant="ghost"
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

      <section className="p-4 border-1 border-dashed mt-2 rounded-lg space-y-4">
        <p className="text-xs">Subcategories</p>
        {subcategories?.map((subcategory) => (
          <div className="flex items-center gap-2" key={subcategory.id}>
            <p className="text-sm text-muted-foreground">– {subcategory.name}</p>
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
                <p>90 products are linked.</p>
              </TooltipContent>
            </Tooltip>
          </div>
        ))}
      </section>
    </div>
  )
}
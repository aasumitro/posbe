import {ProductCard} from "@/pages/catalog/components/product-card.tsx";
import {
  DropdownMenu,
  DropdownMenuContent, DropdownMenuItem,
  DropdownMenuLabel, DropdownMenuSeparator,
  DropdownMenuTrigger
} from "@/components/ui/dropdown-menu.tsx";
import {PackageMinusIcon, PackageOpenIcon, PackageXIcon} from "lucide-react";

export const ProductItem = () => {
  return (
    <DropdownMenu >
      <DropdownMenuTrigger asChild>
        <a href="#">
          <ProductCard/>
        </a>
      </DropdownMenuTrigger>
      <DropdownMenuContent side="right">
        <DropdownMenuLabel>Actions</DropdownMenuLabel>
        <DropdownMenuSeparator />
        <DropdownMenuItem>
          <PackageOpenIcon className="w-4 h-4 mr-2"/>
          Update
        </DropdownMenuItem>
        <DropdownMenuItem>
          <PackageXIcon className="w-4 h-4 mr-2"/>
          Deactivate
        </DropdownMenuItem>
        <DropdownMenuSeparator />
        <DropdownMenuItem className="text-destructive">
          <PackageMinusIcon className="w-4 h-4 mr-2" />
          Delete
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  )
}
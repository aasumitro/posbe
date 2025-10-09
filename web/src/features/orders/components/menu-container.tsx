import {CategoryFilter} from "@/features/orders/components/category-filter";
import {Separator} from "@/components/ui/separator";
import {SubcategoryFilter} from "@/features/orders/components/subcategory-filter";
import {Input} from "@/components/ui/input";
import {IconSearch} from "@tabler/icons-react";
import {MenuList} from "@/features/orders/components/menu-list";
import {useState} from "react";

export function MenuContainer() {
  const [categoryId, setCategoryId] = useState<number>(0)
  const [subcategoryId, setSubcategoryId] = useState<number>(0)

  return (
    <>
      <h5 className="text-lg font-bold">Menu</h5>

      <div className="flex flex-row gap-6 items-center">
        <CategoryFilter
          selectedCategoryId={categoryId}
          setSelectedCategoryId={setCategoryId}
        />
        <Separator orientation="vertical" className="mb-4"/>
        <SubcategoryFilter
          selectedCategoryId={categoryId}
          selectedSubcategoryId={subcategoryId}
          setSelectedSubcategoryId={setSubcategoryId}
        />
      </div>

      <div className="relative mb-4">
        <Input className="h-10 pr-12" placeholder="Search products . . ." />
        <div className="pointer-events-none absolute inset-y-0 right-6 flex items-center">
          <IconSearch  className="w-4 h-4 text-muted-foreground"/>
        </div>
      </div>

      <MenuList />
    </>
  )
}
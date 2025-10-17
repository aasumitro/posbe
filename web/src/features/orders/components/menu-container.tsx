import {CategoryFilter} from "@/features/orders/components/category-filter";
import {Separator} from "@/components/ui/separator";
import {SubcategoryFilter} from "@/features/orders/components/subcategory-filter";
import {Input} from "@/components/ui/input";
import {IconSearch} from "@tabler/icons-react";
import {MenuList} from "@/features/orders/components/menu-list";
import {useState} from "react";
import {cn} from "@/lib/utils";

export function MenuContainer({className}: {className?: string}) {
  const [name, setName] = useState<string>("")
  const [categoryId, setCategoryId] = useState<number>(0)
  const [subcategoryId, setSubcategoryId] = useState<number>(0)

  const onCategoryChange = (categoryId: number) => {
    setCategoryId(categoryId);
    setSubcategoryId(0);
  }

  return (
    <div className={cn(className)}>
      <h5 className="text-lg font-bold">Menu</h5>

      <div className="flex flex-row gap-6 items-center">
        <CategoryFilter
          selectedCategoryId={categoryId}
          setSelectedCategoryId={onCategoryChange}
        />
        <Separator orientation="vertical" className="mb-4"/>
        <SubcategoryFilter
          selectedCategoryId={categoryId}
          selectedSubcategoryId={subcategoryId}
          setSelectedSubcategoryId={setSubcategoryId}
        />
      </div>

      <div className="relative mb-4">
        <Input
          className="h-10 pr-12"
          placeholder="Search products . . ."
          value={name}
          onChange={(e) => setName(e.target.value)}
        />

        {/* Search icon (hidden when typing) */}
        {name === "" && (
          <div className="pointer-events-none absolute inset-y-0 right-6 flex items-center">
            <IconSearch className="w-4 h-4 text-muted-foreground" />
          </div>
        )}

        {/* Clear (x) button when text exists */}
        {name !== "" && (
          <button
            type="button"
            onClick={() => setName("")}
            className="absolute inset-y-0 right-3 flex items-center text-muted-foreground hover:text-foreground"
          >×</button>
        )}
      </div>

      <MenuList name={name} categoryId={categoryId} subcategoryId={subcategoryId} />
    </div>
  )
}
import { Grid } from "lucide-react"
import {useAttributeState} from "@/states/attribute-state";
import {cn} from "@/lib/utils";
import {useMemo} from "react";

export function SubcategoryFilter(
  {
    selectedCategoryId,
    selectedSubcategoryId,
    setSelectedSubcategoryId,
  }: {
    selectedCategoryId: number,
    selectedSubcategoryId: number,
    setSelectedSubcategoryId: (id: number) => void;
  }
) {
  const {subcategories} =  useAttributeState();
  const totalAllProduct = subcategories?.reduce((sum, c) =>
    sum + (c.usage ?? 0), 0) ?? 0;

  const filteredSubcategory = useMemo(() => {
    let list = subcategories?.slice().sort(
      (a, b) => (b.usage ?? 0) - (a.usage ?? 0));
    if (selectedCategoryId !== 0) {
      list = list?.filter(fs => fs.category_id === selectedCategoryId);
    }
    return list;
  }, [subcategories, selectedCategoryId]);

  return (
    <div className="flex gap-3 overflow-x-auto pb-4 px-5 ml-[-20px] mr-[-12px]">
      <div
        onClick={() => setSelectedSubcategoryId(0)}
        className={cn(
          "flex flex-col items-center justify-center min-w-[100px] cursor-pointer",
          " p-3 rounded-xl border hover:bg-green-50 transition-colors",
          {
            "bg-green-50 text-green-600": selectedSubcategoryId === 0,
            "bg-white": selectedSubcategoryId !== 0,
          }
        )}
      >
        <Grid className="h-6 w-6 mb-1" />
        <span className="text-sm font-medium">All</span>
        <span className="text-xs text-gray-500">{totalAllProduct} items</span>
      </div>

      {filteredSubcategory?.map((subcategory) => (
        <div
          key={subcategory.id}
          onClick={() => subcategory.usage && subcategory.usage > 0 && setSelectedSubcategoryId(subcategory.id)}
          className={cn(
            "flex flex-col items-center justify-center p-3 rounded-xl min-w-[100px] border transition-colors",
            {
              "bg-gray-50 text-gray-400 cursor-not-allowed opacity-50 pointer-events-none":
                !subcategory.usage || subcategory.usage === 0,
              "bg-green-50 text-green-600 cursor-pointer hover:bg-green-50":
                subcategory.id === selectedSubcategoryId && subcategory.usage !== 0,
              "bg-white cursor-pointer hover:bg-green-50":
                subcategory.id !== selectedSubcategoryId && subcategory.usage !== 0,
            }
          )}
          aria-disabled={!subcategory.usage || subcategory.usage === 0}
        >
          <span className="text-sm font-medium">
            {subcategory.name.charAt(0).toUpperCase() + subcategory.name.slice(1)}
          </span>
          <span className="text-xs text-gray-500">{subcategory.usage ?? 0} items</span>
        </div>
      ))}
    </div>
  )
}

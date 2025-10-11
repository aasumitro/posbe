import { Grid} from "lucide-react"
import {useAttributeState} from "@/states/attribute-state";
import {cn} from "@/lib/utils";

export function CategoryFilter(
  {
    selectedCategoryId,
    setSelectedCategoryId,
  }: {
    selectedCategoryId: number,
    setSelectedCategoryId: (id: number) => void;
  }
) {
  const {categories} =  useAttributeState();
  const totalAllProduct = categories?.reduce((sum, c) =>
    sum + (c.usage ?? 0), 0) ?? 0;

  return (
    <div className="flex gap-3 overflow-x-auto pb-4 min-w-[325px] max-w-1/2">
      <div
        onClick={() => setSelectedCategoryId(0)}
        className={cn(
          "flex flex-col items-center justify-center min-w-[100px] cursor-pointer",
          " p-3 rounded-xl border hover:bg-green-50 transition-colors",
          {
            "bg-green-50 text-green-600": selectedCategoryId === 0,
            "bg-white": selectedCategoryId !== 0,
          }
        )}
      >
        <Grid className="h-6 w-6 mb-1" />
        <span className="text-sm font-medium">All</span>
        <span className="text-xs text-gray-500">{totalAllProduct} items</span>
      </div>

      {categories?.map((category) => (
        <div
          key={category.id}
          onClick={() => category.usage && category.usage > 0 && setSelectedCategoryId(category.id)}
          className={cn(
            "flex flex-col items-center justify-center p-3 rounded-xl min-w-[100px] border transition-colors",
            {
              "bg-gray-50 text-gray-400 cursor-not-allowed opacity-50 pointer-events-none":
                !category.usage || category.usage === 0,
              "bg-green-50 text-green-600 cursor-pointer hover:bg-green-50":
                category.id === selectedCategoryId && category.usage !== 0,
              "bg-white cursor-pointer hover:bg-green-50":
                category.id !== selectedCategoryId && category.usage !== 0,
            }
          )}
          aria-disabled={!category.usage || category.usage === 0}
        >
          <span className="text-sm font-medium">
            {category.name.charAt(0).toUpperCase() + category.name.slice(1)}
          </span>
          <span className="text-xs text-gray-500">{category.usage} items</span>
        </div>
      ))}
    </div>
  )
}

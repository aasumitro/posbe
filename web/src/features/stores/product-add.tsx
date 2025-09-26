import {useCategoryList, useUnitList} from "@/hooks/use-attribute";
import {ProductActionAdd} from "@/features/stores/components/catalog/product-action-add";

export function NewProductPage() {
  const {data: categories, isLoading: isLoadCategory} = useCategoryList();
  const {data: units, isLoading: isLoadUnit} = useUnitList();

  return (
    <div className="w-full pt-4 xl:pt-6 space-y-4">
      {(isLoadCategory || isLoadUnit) && (
        // TODO: add skeleton
        <>loading . . .</>
      )}

      <ProductActionAdd
        categories={categories.data}
        units={units.data}
      />
    </div>
  )
}

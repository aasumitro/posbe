import {useCategoryList, useUnitList} from "@/hooks/use-attribute";
import {ProductActionAdd} from "@/features/stores/components/catalog/product-action-add";
import {AppLoading} from "@/components/app-loading";

export function NewProductPage() {
  const {data: categories, isLoading: isLoadCategory} = useCategoryList();
  const {data: units, isLoading: isLoadUnit} = useUnitList();

  if (isLoadCategory && isLoadUnit) {
    return <AppLoading />
  }

  return (
    <div className="w-full pt-4 xl:pt-6 space-y-4">
      <ProductActionAdd
        categories={categories.data}
        units={units.data}
      />
    </div>
  )
}

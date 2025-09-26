import {useCategoryList, useUnitList} from "@/hooks/use-attribute";
import {ProductActionEdit} from "@/features/stores/components/catalog/product-action-edit";
import {useProductDetail} from "@/hooks/use-product";
import {useParams} from "@tanstack/react-router";

export function EditProductPage() {
  const { id } = useParams({ from: '/_authenticated/stores/products/$id' })
  if (!id || isNaN(Number(id))) {
    alert("No product ID found. Redirecting back...");
    window.location.href = "/stores/catalogs";
    return;
  }

  const {data: product, isLoading: isLoadProduct} = useProductDetail(Number(id));
  const {data: categories, isLoading: isLoadCategory} = useCategoryList();
  const {data: units, isLoading: isLoadUnit} = useUnitList();

  return (
    <div className="w-full pt-4 xl:pt-6 space-y-4">
      {(isLoadProduct || isLoadCategory || isLoadUnit) && (
        // TODO: add skeleton
        <>loading . . .</>
      )}

      {!product?.data && (
        <>Product not found get back</>
      )}

      <ProductActionEdit
        product={product?.data}
        categories={categories.data}
        units={units.data}
      />
    </div>
  )
}
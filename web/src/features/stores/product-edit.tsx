import {useCategoryList, useUnitList} from "@/hooks/use-attribute";
import {ProductActionEdit} from "@/features/stores/components/catalog/product-action-edit";
import {useProductDetail} from "@/hooks/use-product";
import {useParams} from "@tanstack/react-router";
import {AppLoading} from "@/components/app-loading";
import {useEffect} from "react";
import {useProductState} from "@/states/product-state";
import {useAttributeState} from "@/states/attribute-state";

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

  const {setSelectedProduct} = useProductState();
  const {setCategories, setUnits} = useAttributeState();

  useEffect(() => {
    if (product?.data) setSelectedProduct(product.data);
    if (categories?.data) setCategories(categories.data);
    if (units?.data) setUnits(units.data);
  }, [product?.data, categories?.data, units.data]);

  if (isLoadProduct && isLoadCategory && isLoadUnit) {
    return <AppLoading />
  }

  return (
    <div className="w-full pt-4 xl:pt-6 space-y-4">
      {!product?.data && <>Product not found get back</>}

      <ProductActionEdit />
    </div>
  )
}
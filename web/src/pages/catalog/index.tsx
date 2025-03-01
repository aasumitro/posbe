import {AddonList} from "@/pages/catalog/components/addon-list.tsx";
import {ProductList} from "@/pages/catalog/components/product-list.tsx";

export function CatalogPage() {
  return (
    <div className="flex flex-col xl:flex-row gap-4 p-4 md:gap-8 md:p-10">
      <ProductList />
      <AddonList />
    </div>
  )
}

import {NewProductCard} from "@/pages/catalog/components/new-product-card.tsx";
import {cn} from "@/lib/utils.ts";
import {ProductItem} from "@/pages/catalog/components/product-item.tsx";
import {ProductFilter} from "@/pages/catalog/components/product-filter.tsx";
export const ProductList = () => {
  const products = [1]

  return (
    <div className="w-full xl:w-2/3 border-2 rounded-lg border-dashed p-6">
      <div className="flex items-center justify-between">
        <section>
          <h3 className="text-2xl font-semibold leading-none tracking-tight">Products</h3>
          <p className="text-sm text-muted-foreground">lorem ipsum dolor si amet marque test molor</p>
        </section>
        <ProductFilter />
      </div>
      <section className={cn(
        "flex flex-row flex-wrap mt-6 gap-4",
        // {"justify-between": products.length > 2}
      )}>
        {/*display new product card only if store doesnt have any product*/}
        {/*if store has product then hide it move the add button into product fitler*/}
        <NewProductCard/>
        {products.map((index) => (
          <ProductItem key={index} />
        ))}
      </section>
    </div>
  )
}
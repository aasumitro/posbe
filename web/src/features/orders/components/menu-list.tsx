import {IconArchiveOff} from "@tabler/icons-react";
import {Box, Coffee, Soup} from "lucide-react";
import {useOrderState} from "@/states/order-state";
import {useMemo} from "react";
import type {Product} from "@/types/product";
import {ASSET_URL} from "@/lib/api";
import {formatShortNumber} from "@/lib/numbers";
import {useStoreState} from "@/states/store-state";
import {cn} from "@/lib/utils";
import {useActionState} from "@/states/action-state";
import {AddToCartModalState} from "@/features/orders/components/add-to-cart-modal";

export function MenuList({
  name,
  categoryId,
  subcategoryId
}: {
  name: string,
  categoryId: number,
  subcategoryId: number,
}) {
  const { settings } = useStoreState();
  const { products, setSelectedProduct } =  useOrderState();
  const { setBoolState } = useActionState();

  const filteredProducts = useMemo(() => {
    if (!products) return [];

    return products.filter((product) => {
      const matchName = !name || product.name.toLowerCase().includes(name.toLowerCase());
      const matchCategory = categoryId === 0 || product.category_id === categoryId;
      const matchSubcategory = subcategoryId === 0 || product.subcategory_id === subcategoryId;
      return matchName && matchCategory && matchSubcategory;
    });
  }, [products, name, categoryId, subcategoryId]);

  if (!filteredProducts || filteredProducts.length === 0) {
    return (
      <div className="text-center py-12">
        <IconArchiveOff className="my-4 w-14 h-14 mx-auto"/>
        <h4 className="text-primary text-md font-bold tracking-tight">
          No Matching Products
        </h4>
        <p className="text-secondary-foreground text-xs font-normal">
          Try adjusting your filters or search keywords.
        </p>
      </div>
    )
  }

  function renderActiveProductImage(product: Product) {
    if (!product.image && product.category) {
      if (product.category.name === "foods") {
        return <Soup />
      }

      if (product.category.name === "beverages") {
        return <Coffee />
      }

      return <Box />
    }

    return <img src={`${ASSET_URL}/${product.image}`} alt={product.name} />
  }

  function renderActiveProductName(product: Product) {
    const productName = product?.name?.trim() ?? "";

    // Capitalize every word
    const formatted = productName.split(" ").filter(Boolean)
      .map(word => word.charAt(0).toUpperCase() + word.slice(1))
      .join(" ");

    // Ensure it ends with a period
    const formattedName = formatted.endsWith(".") ? formatted : formatted + ".";

    return (
      <h5 className="text-xl font-semibold">
        {formattedName}
      </h5>
    );
  }

  function renderActiveProductTags(product: Product) {
    return (
      <div className="text-muted-foreground text-xs flex gap-2">
        {product.category && (<p>#{product.category.name}</p>)}
        {product.subcategory && (<p>#{product.subcategory.name}</p>)}
      </div>
    )
  }

  function renderActiveProductPrice(product: Product, currency: string | undefined) {
    if (!product?.variants || product.variants.length === 0) {
      return (
        <h5 className="text-3xl font-semibold text-green-500">
          N/A
        </h5>
      );
    }

    // If more than one variant → find lowest
    if (product.variants.length > 1) {
      const nonZero = product.variants.filter(v => v.price > 0);
      const lowestVariant = nonZero.length > 0
        ? nonZero.reduce((min, v) => v.price < min.price ? v : min)
        : product.variants[0];

      return (
        <section className="flex">
        <span className="text-sm font-normal text-muted-foreground mr-1">
          from
        </span>
          <div className="flex items-baseline gap-1 text-3xl font-bold tabular-nums leading-none text-green-500">
            {currency} {formatShortNumber(lowestVariant.price)}
            {lowestVariant.unit && (
              <span className="text-sm font-normal text-muted-foreground">
                / {lowestVariant.unit_size}{lowestVariant.unit?.symbol}
              </span>
            )}
          </div>
        </section>
      );
    }

    // Exactly one variant
    const variant = product.variants[0];
    return (
      <h5 className="text-3xl font-semibold text-green-500 flex items-baseline gap-1">
        {currency} {formatShortNumber(variant.price)}
        {variant.unit && (
          <span className="text-sm font-normal text-muted-foreground">
            / {variant.unit.symbol}
          </span>
        )}
      </h5>
    );
  }

  function renderActiveProductVariants(product: Product) {
    const variants = product?.variants ?? [];

    if (variants.length <= 1) return null;

    return (
      <p className="text-xs text-muted-foreground flex items-center gap-1">
        <span>🧩</span> Customizable options available
      </p>
    );
  }

  function onProductSelected(product: Product) {
    setSelectedProduct(product);
    setBoolState(AddToCartModalState, true)
  }

  return (
    <div className="flex flex-wrap gap-4 justify-start items-stretch">
      {filteredProducts?.map((product, index) => (
        <div key={index} className="flex-[1_1_400px] sm:flex-[1_1_50%] lg:flex-[1_1_33%] 2xl:flex-[1_1_25%]">
          <div
            className={cn(
              "border-1 rounded-xl p-4 space-y-6 select-none flex flex-col h-full w-full",
              "hover:bg-green-50/50 cursor-pointer hover:shadow-md"
            )}
            onClick={(e) => {
              e.preventDefault();
              onProductSelected(product);
            }}
          >
            <div className="flex-grow space-y-6">
              <div className="flex gap-4">
                <div className="rounded-lg w-20 h-20 bg-gray-50/50 flex items-center justify-center">
                  {renderActiveProductImage(product)}
                </div>

                <div className="space-y-1">
                  {renderActiveProductName(product)}

                  {renderActiveProductTags(product)}

                  {renderActiveProductPrice(product, settings?.currency)}

                  {renderActiveProductVariants(product)}
                </div>
              </div>
            </div>
          </div>
        </div>
      ))}
    </div>
  )
}
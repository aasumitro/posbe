import {Coffee, Soup} from "lucide-react";
import {Badge} from "@/components/ui/badge";
import {Separator} from "@/components/ui/separator";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger
} from "@/components/ui/dropdown-menu";
import {Button} from "@/components/ui/button";
import {IconDotsVertical} from "@tabler/icons-react";
import {useActionState} from "@/states/action-state";
import {ProductActionDeleteModalState} from "@/features/stores/components/catalog/product-action-delete";
import {useNavigate} from "@tanstack/react-router";
import {useStoreState} from "@/states/store-state";
import {useProductState} from "@/states/product-state";
import {formatShortNumber} from "@/lib/numbers";
import type {Product, ProductVariant} from "@/types/product";

export function  ProductContainer() {
  const {settings} = useStoreState();
  const {products} = useProductState();
  const {setBoolState} = useActionState();
  const navigate = useNavigate();

  function renderImage(product: Product) {
    if (!product.image && product.category) {
      if (product.category.name === "foods") {
        return <Soup />
      }
      return <Coffee />
    }

    return <img src={product.image} alt={product.name} />
  }

  function renderName(product: Product) {
    const name = product?.name?.trim() ?? "";

    // Capitalize every word
    const formatted = name
      .split(" ")
      .filter(Boolean) // remove extra spaces
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

  function renderCategoryTags(product: Product) {
    return (
      <div className="text-muted-foreground text-xs flex gap-2">
        {product.category && (<p>#{product.category.name}</p>)}
        {product.subcategory && (<p>#{product.subcategory.name}</p>)}
      </div>
    )
  }

  function renderPrice(product: Product, currency: string | undefined) {
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
            <span className="text-sm font-normal text-muted-foreground">
            / {lowestVariant.unit_size}{lowestVariant.unit?.symbol}
          </span>
          </div>
        </section>
      );
    }

    // Exactly one variant
    const variant = product.variants[0];
    return (
      <h5 className="text-3xl font-semibold text-green-500">
        {currency} {formatShortNumber(variant.price)}
      </h5>
    );
  }

  function renderActionMenu(product: Product) {
    return (
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button className="ml-auto my-auto" variant="ghost">
            <IconDotsVertical />
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="end">
          <DropdownMenuLabel>Actions</DropdownMenuLabel>
          <DropdownMenuSeparator/>
          <DropdownMenuItem onClick={async (e) => {
            e.preventDefault();
            await navigate({to: `/stores/products/${product.id}`})
          }}>Edit</DropdownMenuItem>
          {/*<DropdownMenuItem>set status</DropdownMenuItem>*/}
          <DropdownMenuItem variant="destructive" onClick={(e) => {
            e.preventDefault();
            setBoolState(ProductActionDeleteModalState, true)
          }}>Delete</DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    )
  }

  function renderSales(product: Product) {
    return (
      <div className="flex w-full items-center gap-2">
        <div className="grid flex-1 auto-rows-min gap-0.5">
          <div className="text-xs text-muted-foreground">This Year</div>
          <div className="flex items-baseline gap-1 text-2xl font-bold tabular-nums leading-none">
            {product.sales_this_year}
            <span className="text-sm font-normal text-muted-foreground">
              sales
            </span>
          </div>
        </div>

        <Separator orientation="vertical" className="mx-2 h-10 w-px" />

        <div className="grid flex-1 auto-rows-min gap-0.5">
          <div className="text-xs text-muted-foreground">This Week</div>
          <div className="flex items-baseline gap-1 text-2xl font-bold tabular-nums leading-none">
            {product.sales_this_week}
            <span className="text-sm font-normal text-muted-foreground">
              sales
            </span>
          </div>
        </div>

        <Separator orientation="vertical" className="mx-2 h-10 w-px" />

        <div className="grid flex-1 auto-rows-min gap-0.5">
          <div className="text-xs text-muted-foreground">Today</div>
          <div className="flex items-baseline gap-1 text-2xl font-bold tabular-nums leading-none">
            {product.sales_today}
            <span className="text-sm font-normal text-muted-foreground">
              sales
            </span>
          </div>
        </div>
      </div>
    )
  }

  function renderVariants(product: Product) {
    if (!product.variants || product.variants.length === 0) return null;

    // Group variants by type
    const grouped = product.variants.reduce<
      Record<string, ProductVariant[]>
    >((acc, v) => {
      if (!acc[v.type]) acc[v.type] = [];
      acc[v.type].push(v);
      return acc;
    }, {});

    return (
      <div className="grid grid-cols-2 gap-4">
        {Object.entries(grouped).map(([type, variants]) => (
          <div key={type} className="flex flex-col gap-2">
            <p className="text-sm font-medium">
              {type === "size"
                ? "Size Options"
                : type === "none"
                  ? (variants[0]?.description
                    ? variants[0].description.charAt(0).toUpperCase() + variants[0].description.slice(1)
                    : "Other Options")
                  : type.charAt(0).toUpperCase() + type.slice(1)}
            </p>
            <div className="flex flex-row gap-2 flex-wrap">
              {variants.map((v) => (
                <Badge key={v.id} variant="outline">
                  {v.name.toUpperCase()}
                </Badge>
              ))}
            </div>
          </div>
        ))}
      </div>
    );
  }

  return (
    <div className="grid gap-4 grid-cols-2 2xl:grid-cols-3 auto-rows-fr mt-4">
      {products?.map((product) => (
        <div
          key={product.id}
          className="border-1 rounded-xl p-4 space-y-6 select-none flex flex-col h-full"
        >
          <div className="flex-grow space-y-6">
            <div className="flex gap-4">
              <div className="rounded-lg w-20 h-20 bg-gray-50/50 flex items-center justify-center">
                {renderImage(product)}
              </div>

              <div className="space-y-1">
                {renderName(product)}

                {renderCategoryTags(product)}

                {renderPrice(product, settings?.currency)}
              </div>

              {renderActionMenu(product)}
            </div>

            {renderVariants(product)}
          </div>

          <div className="flex flex-row border-t p-4">
            {renderSales(product)}
          </div>
        </div>
      ))}
    </div>
  )
}
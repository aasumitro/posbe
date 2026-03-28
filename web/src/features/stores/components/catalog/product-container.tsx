import {Box, Coffee, Soup} from "lucide-react";
import {Badge} from "@/components/ui/badge";
import {Separator} from "@/components/ui/separator";
import {
  DropdownMenu,
  DropdownMenuContent, DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel, DropdownMenuPortal,
  DropdownMenuSeparator, DropdownMenuSub, DropdownMenuSubContent, DropdownMenuSubTrigger,
  DropdownMenuTrigger
} from "@/components/ui/dropdown-menu";
import {Button} from "@/components/ui/button";
import {IconBoxOff, IconDotsVertical} from "@tabler/icons-react";
import {useActionState} from "@/states/action-state";
import {ProductActionDeleteModalState} from "@/features/stores/components/catalog/product-action-delete";
import {useNavigate} from "@tanstack/react-router";
import {useStoreState} from "@/states/store-state";
import {useProductState} from "@/states/product-state";
import {formatShortNumber} from "@/lib/numbers";
import type {Product, ProductVariant} from "@/types/product";
import {Tooltip, TooltipContent, TooltipTrigger} from "@/components/ui/tooltip";
import {ASSET_URL, isHTTPResponse} from "@/lib/api";
import {cn} from "@/lib/utils";
import {useUpdateProduct} from "@/hooks/use-product";
import {useQueryClient} from "@tanstack/react-query";
import {toast} from "sonner";

interface ProductContainerProps {
  sort?: Record<string, "asc" | "desc">;
  status?: "draft" | "active" | "inactive";
  category?: string;
}

export function  ProductContainer({status, sort, category}: ProductContainerProps) {
  const {settings} = useStoreState();
  const {products, setSelectedProduct} = useProductState();
  const {setBoolState} = useActionState();
  const navigate = useNavigate();
  const {mutate: editProduct} = useUpdateProduct();
  const queryClient = useQueryClient();

  let sortedProducts = [...(products ?? [])];

  if (sort) {
    const [field, order] = Object.entries(sort)[0] ?? [];
    if (field && order) {
      sortedProducts.sort((a, b) => {
        let aValue: any;
        let bValue: any;

        if (field === "price") {
          const getLowestPrice = (product: Product) => {
            if (!product.variants || product.variants.length === 0) return 0;
            const nonZero = product.variants.filter(v => v.price > 0);
            const lowestVariant = nonZero.length > 0
              ? nonZero.reduce((min, v) => v.price < min.price ? v : min)
              : product.variants[0];
            return lowestVariant.price;
          };
          aValue = getLowestPrice(a);
          bValue = getLowestPrice(b);
        } else {
          aValue = a[field as keyof typeof a];
          bValue = b[field as keyof typeof b];
        }

        // Now compare strings
        if (typeof aValue === "string" && typeof bValue === "string") {
          return order === "asc"
            ? aValue.localeCompare(bValue)
            : bValue.localeCompare(aValue);
        }

        // Compare numbers (including price)
        if (typeof aValue === "number" && typeof bValue === "number") {
          return order === "asc" ? aValue - bValue : bValue - aValue;
        }

        return 0;
      });
    }
  }

  if (status) {
    sortedProducts = sortedProducts
      .filter((product) => product.status === status)
  }

  if (category) {
    sortedProducts = sortedProducts
      .filter((product) => product.category?.name === category)
  }

  function renderImage(product: Product) {
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
      <h5 className="text-xl font-semibold flex ">
        {formattedName}
        <Tooltip>
          <TooltipTrigger asChild>
            <span className={cn(
              "flex size-2 rounded-full ml-2 animate-pulse cursor-pointer",
              product.status === "draft" && "bg-gray-500",
              product.status === "active" && "bg-green-500",
              product.status === "inactive" && "bg-red-500"
            )} title="New"/>
          </TooltipTrigger>
          <TooltipContent>
            <p>
              {["draft", "inactive"].includes(product.status)
                ? `This product is marked as ${product.status}.`
                : "This product is live and ready for sale!"}
            </p>
          </TooltipContent>
        </Tooltip>

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

  function renderActionMenu(product: Product) {
    return (
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button className="ml-auto my-auto cursor-pointer" variant="ghost">
            <IconDotsVertical />
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent className="w-48" align="start">
          <DropdownMenuLabel>Manage Product</DropdownMenuLabel>
          <DropdownMenuGroup>
            <DropdownMenuItem
              className="cursor-pointer"
              onClick={async (e) => {
                e.preventDefault();
                await navigate({to: `/stores/products/${product.id}`})
              }}
            >Edit Details</DropdownMenuItem>
            <DropdownMenuSub>
              <DropdownMenuSubTrigger>Change Status</DropdownMenuSubTrigger>
              <DropdownMenuPortal>
                <DropdownMenuSubContent>
                  {product.status === "active" && (
                    <>
                      <DropdownMenuItem
                        className="cursor-pointer"
                        onClick={(e) => {
                          e.preventDefault();
                          setProductStatus("draft", product)
                        }}
                      >Set as Draft</DropdownMenuItem>
                      <DropdownMenuItem
                        className="cursor-pointer"
                        onClick={(e) => {
                          e.preventDefault();
                          setProductStatus("inactive", product)
                        }}
                      >Set as Inactive</DropdownMenuItem>
                    </>
                  )}
                  {["draft", "inactive"].includes(product.status) && (
                    <DropdownMenuItem
                      className="cursor-pointer"
                      onClick={(e) => {
                        e.preventDefault();
                        setProductStatus("publish", product)
                      }}
                    >Set as Active</DropdownMenuItem>
                  )}
                </DropdownMenuSubContent>
              </DropdownMenuPortal>
            </DropdownMenuSub>
          </DropdownMenuGroup>

          <DropdownMenuSeparator/>
          <DropdownMenuLabel>Sales Insights</DropdownMenuLabel>
          <DropdownMenuGroup>
            <DropdownMenuItem
              className="cursor-pointer"
              disabled
            >Key Metrics</DropdownMenuItem>
            <DropdownMenuItem
              className="cursor-pointer"
              disabled
            >Performance Charts</DropdownMenuItem>
          </DropdownMenuGroup>

          <DropdownMenuSeparator/>
          <DropdownMenuItem
            className="cursor-pointer"
            variant="destructive"
            onClick={(e) => {
              e.preventDefault();
              setSelectedProduct(product);
              setBoolState(ProductActionDeleteModalState, true);
            }}
          >Delete</DropdownMenuItem>
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
    if (!product.variants || product.variants.length === 1) return null;

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
              {type.charAt(0).toUpperCase() + type.slice(1)}
            </p>
            <div className="flex flex-row gap-2 flex-wrap">
              {variants.map((v) => (
                <Tooltip key={v.id} >
                  <TooltipTrigger asChild>
                    <Badge variant="outline">
                      {v.name.toUpperCase()}
                    </Badge>
                  </TooltipTrigger>
                  <TooltipContent>
                    <p>{v.description}</p>
                  </TooltipContent>
                </Tooltip>
              ))}
            </div>
          </div>
        ))}
      </div>
    );
  }

  function setProductStatus(status: string, product: Product) {
    if (!product) return;

    if (!confirm(`Are you sure you want to set this product as ${status}?`)) return;

    editProduct({
      id: product.id, body: JSON.stringify({status})
    }, {
      onSuccess: async () => {
        await queryClient.invalidateQueries({ queryKey: ['products'] })
        toast.success("Product status update successfully");
      },
      onError: (error) => {
        if (error && isHTTPResponse<null>(error)) {
          if (typeof error.data === "string") {
            toast.error(error.data);
            return;
          }

          if (typeof error.data === "object" && error.data !== null) {
            // TODO: apply this
            // const data = error.data as ProductErrorResponse;
            //
            // if (data.name && data.name.length > 0) {
            //   form.setError("name", {type: "manual", message: data.name[0]})
            // }
          }
        }

        if (error instanceof Error) {
          const clientError = error as Error
          toast.error(clientError.message);
        }
      },
    })
  }

  if (sortedProducts.length < 1) {
    return (
      <div className="text-center py-12">
        <IconBoxOff className="my-8 w-24 h-24 mx-auto"/>
        <h4 className="text-primary text-xl font-bold tracking-tight">
          No Products Found
        </h4>
        <p className="text-secondary-foreground text-md font-normal">
          {products?.length === 0
            ? "No products exist yet. Add your first product!"
            : status !== undefined
              ? "There are no products available in this filter at the moment."
              : "No products match the current selection."}
        </p>
      </div>
    )
  }

  return (
    <div className="grid gap-4 grid-cols-2 2xl:grid-cols-3 auto-rows-fr mt-4">
      {sortedProducts?.map((product) => (
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
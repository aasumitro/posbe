import {type ReactElement, useEffect, useState} from "react";
import {Tabs, TabsContent, TabsList, TabsTrigger} from "@/components/ui/tabs";
import {ProductContainer} from "@/features/stores/components/catalog/product-container";
import {AddonContainer} from "@/features/stores/components/catalog/addon-container";
import {
  DropdownMenu, DropdownMenuCheckboxItem,
  DropdownMenuContent, DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger
} from "@/components/ui/dropdown-menu";
import {Button} from "@/components/ui/button";
import {File, ListFilter, Plus, SortAscIcon} from "lucide-react";
import {useActionState} from "@/states/action-state";
import {useNavigate, useSearch} from "@tanstack/react-router";
import {AddonActionAdd, AddonActionAddModalState} from "@/features/stores/components/catalog/addon-action-add";
import {ProductActionDeleteModal} from "@/features/stores/components/catalog/product-action-delete";
import {AddonActionEdit} from "@/features/stores/components/catalog/addon-action-edit";
import {AddonActionDelete} from "@/features/stores/components/catalog/addon-action-delete";
import {useProductAddonList, useProductList} from "@/hooks/use-product";
import {useProductState} from "@/states/product-state";
import {Route} from "@/routes/_authenticated/stores/route";
import {useCategoryList} from "@/hooks/use-attribute";
import {IconSort09, IconSort90, IconSortAZ, IconSortZA} from "@tabler/icons-react";

interface CatalogPageProps {
  query: {
    tab: "products" | "addons";
    sort?: Record<string, "asc" | "desc">;
    status?: "draft" | "active" | "inactive";
    category?: string;
  };
}

export function ProductCatalogPage({query}: CatalogPageProps) {
  const { setBoolState } = useActionState();
  const [activeTab, setActiveTab] = useState<string>(query.tab);
  const navigate = useNavigate();
  const { sort: rawSort, status, category } = useSearch({from: '/_authenticated/stores/catalogs'});
  const {data: products} = useProductList();
  const {data: addons} = useProductAddonList();
  const {data: categories} = useCategoryList();
  const {setProducts, setAddons} = useProductState();

  useEffect(() => {
    if (products?.data) setProducts(products.data);
    if (addons?.data) setAddons(addons.data);
  }, [products?.data, addons?.data]);

  async function onAddNewItem() {
    if (activeTab === "products") {
      await navigate({to: "/stores/products"})
      return;
    }

    setBoolState(AddonActionAddModalState, true);
  }

  const handleTabChange = async (value: string) => {
    setActiveTab(value);
    await navigate({from:Route.fullPath, search: {tab: value}});
  }

  const updateSearchQuery = async (
    updates: { status?: string; category?: string; sort?: string }
  ) => {
    let search: { tab: string; sort?: string; status?: string; category?: string } = {
      tab: activeTab,
      sort: rawSort ? decodeURIComponent(rawSort) : undefined,
      status,
      category,
    };

    search = { ...search, ...updates };

    if (updates.category) {
      delete search.status;
    }

    if (updates.status) {
      delete search.category;
    }

    Object.keys(search).forEach((key) => {
      if (search[key as keyof typeof search] === "" ||
        search[key as keyof typeof search] === undefined) {
        delete search[key as keyof typeof search];
      }
    });

    await navigate({from: Route.fullPath, search});
  };

  const FilterDropdown = ({tab}: {tab: string}) => {
    if (tab !== "products") return null;

    return (
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button
            variant="outline"
            size="sm"
            className="h-7 gap-1 text-sm cursor-pointer"
          >
            <ListFilter className="h-3.5 w-3.5"/>
            <span className="sr-only sm:not-sr-only">Filter</span>
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="end">
          <DropdownMenuLabel>Filter by</DropdownMenuLabel>
          <DropdownMenuSeparator/>

          <DropdownMenuLabel className="text-muted-foreground text-xs">Status</DropdownMenuLabel>
          {(["all", "draft", "active", "inactive"] as const).map((s) => (
            <DropdownMenuCheckboxItem
              key={s}
              className="cursor-pointer"
              checked={s === "all" ? !status : status === s}
              onClick={async () => {
                if (s === "all") {
                  await updateSearchQuery({ status: "" });
                } else {
                  await updateSearchQuery({ status: s });
                }
              }}
            >
              {s === "all" ? "All" : s.charAt(0).toUpperCase() + s.slice(1)}
            </DropdownMenuCheckboxItem>
          ))}

          {categories.data && categories.data.length > 0 && (
            <>
              <DropdownMenuLabel className="text-muted-foreground text-xs">
                Category
              </DropdownMenuLabel>
              {categories?.data?.map((c) => (
                <DropdownMenuCheckboxItem
                  key={c.id}
                  className="cursor-pointer"
                  checked={category === c.name}
                  onClick={() => updateSearchQuery({category: c.name})}
                >
                  {c.name.charAt(0).toUpperCase() + c.name.slice(1)}
                </DropdownMenuCheckboxItem>
              ))}
            </>
          )}

          <DropdownMenuSeparator/>
          <DropdownMenuItem
            className="cursor-pointer"
            variant="destructive"
            onClick={async () => {
              await navigate({
                from:Route.fullPath,
                search: {
                  tab: activeTab,
                  sort: rawSort ? decodeURIComponent(rawSort) : undefined,
                }
              });
            }}
          >Clear</DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    )
  }

  const renderSortOptions = (
    options: string[], icons: [ReactElement, ReactElement]
  ) => options.map((filter, index) => (
      <DropdownMenuCheckboxItem
        key={filter}
        className="cursor-pointer"
        checked={rawSort === filter}
        onClick={() => updateSearchQuery({ sort: filter })}
      >
        {index === 0 ? (
          <>{icons[0]} Ascending</>
        ) : (
          <>{icons[1]} Descending</>
        )}
      </DropdownMenuCheckboxItem>
  ));

  const SortDropdown = ({tab}: {tab: string}) => {
    return (
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button
            variant="outline"
            size="sm"
            className="h-7 gap-1 text-sm cursor-pointer"
          >
            <SortAscIcon className="h-3.5 w-3.5"/>
            <span className="sr-only sm:not-sr-only">Sort</span>
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="end">
          <DropdownMenuLabel>Sort by</DropdownMenuLabel>
          <DropdownMenuSeparator/>

          <DropdownMenuLabel className="text-muted-foreground text-xs">Name</DropdownMenuLabel>
          {renderSortOptions(["name:asc", "name:desc"], [<IconSortAZ />, <IconSortZA />])}


          <DropdownMenuLabel className="text-muted-foreground text-xs">Price</DropdownMenuLabel>
          {renderSortOptions(["price:asc", "price:desc"], [<IconSort09 />, <IconSort90 />])}

          {tab === "addons" && (
            <>
              <DropdownMenuLabel className="text-muted-foreground text-xs">Orders</DropdownMenuLabel>
              {renderSortOptions(["usage:asc", "usage:desc"], [<IconSort09 />, <IconSort90 />])}
            </>
          )}

          <DropdownMenuSeparator/>
          <DropdownMenuItem
            className="cursor-pointer"
            variant="destructive"
            onClick={() => updateSearchQuery({sort: ""})}
          >Clear</DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    )
  }

  return (
    <div className="w-full pt-4 xl:pt-6 space-y-6">
      <aside className="flex items-center justify-between gap-4 px-8">
        <div className="block">
          <h3 className="text-lg font-semibold mb-2">Catalogs</h3>
          <p className="text-sm text-muted-foreground">
            Manage core catalog data such as products and addons.
            Organize and control the items available in your system.
          </p>
        </div>
      </aside>

      <aside className="mt-4 gap-6 px-8">
        <Tabs
          value={activeTab}
          className="h-full w-full"
          onValueChange={handleTabChange}
        >
          <div className="flex items-center">
            <TabsList>
              <TabsTrigger
                value="products"
                className="flex items-center space-x-2 px-4 cursor-pointer"
              >
                Products
              </TabsTrigger>
              <TabsTrigger
                value="addons"
                className="flex items-center space-x-2 px-4 cursor-pointer"
              >
                Addons
              </TabsTrigger>
            </TabsList>

            <div className="ml-auto flex items-center gap-2">
              <Button
                size="sm"
                variant="outline"
                className="h-7 gap-1 text-sm cursor-pointer"
                onClick={async (e) => {
                  e.preventDefault();
                  await onAddNewItem();
                }}
              >
                <Plus className="h-3.5 w-3.5"/>
                <span className="sr-only sm:not-sr-only capitalize">New {activeTab}</span>
              </Button>

              <FilterDropdown tab={activeTab} />

              <SortDropdown tab={activeTab} />

              <Button
                size="sm"
                variant="outline"
                className="h-7 gap-1 text-sm"
                disabled
              >
                <File className="h-3.5 w-3.5"/>
                <span className="sr-only sm:not-sr-only">Export</span>
              </Button>
            </div>
          </div>

          <TabsContent value="products">
            <ProductContainer
              sort={query.sort}
              status={query.status}
              category={query.category}
            />
            <ProductActionDeleteModal />
          </TabsContent>
          <TabsContent value="addons">
            <AddonContainer sort={query.sort} />
            <AddonActionAdd />
            <AddonActionEdit />
            <AddonActionDelete />
          </TabsContent>
        </Tabs>
      </aside>
    </div>
  )
}
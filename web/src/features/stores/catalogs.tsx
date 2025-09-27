import {useEffect, useState} from "react";
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

interface CatalogPageProps {
  query: {
    tab: "products" | "addons";
    sort?: Record<string, "asc" | "desc">;
    status?: "draft" | "active" | "inactive";
  };
}

export function ProductCatalogPage({query}: CatalogPageProps) {
  const { setBoolState } = useActionState();
  const [activeTab, setActiveTab] = useState<string>(query.tab);
  const navigate = useNavigate();
  const { sort: rawSort, status } = useSearch({from: '/_authenticated/stores/catalogs'});
  const {data: products} = useProductList();
  const {data: addons} = useProductAddonList();
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
    await navigate({
      from:Route.fullPath,
      search: {
        tab: value,
        sort: rawSort ? decodeURIComponent(rawSort) : undefined,
        status: status
      }
    });
  }

  const statusQuery = async (status: string) => {
    let search: {tab: string; sort?: string; status?: string} = {
      tab: activeTab,
      sort: rawSort ? decodeURIComponent(rawSort) : undefined,
    }

    if (status) {
      search.status = status;
    }

    await navigate({
      from:Route.fullPath,
      search
    });
  }

  const sortQuery = async (sort: string) => {
    let search: {tab: string; sort?: string; status?: string} = {
      tab: activeTab,
      status: status
    }

    if (sort) {
      search.sort = sort;
    }

    await navigate({
      from:Route.fullPath,
      search
    });
  }

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
          <DropdownMenuLabel className="text-muted-foreground text-xs">
            Status
          </DropdownMenuLabel>
          <DropdownMenuCheckboxItem
            className="cursor-pointer"
            checked={!status}
            onClick={() => statusQuery("")}
          >All</DropdownMenuCheckboxItem>
          <DropdownMenuCheckboxItem
            className="cursor-pointer"
            checked={status === "draft"}
            onClick={() => statusQuery("draft")}
          >Draft</DropdownMenuCheckboxItem>
          <DropdownMenuCheckboxItem
            className="cursor-pointer"
            checked={status === "active"}
            onClick={() => statusQuery("active")}
          >Active</DropdownMenuCheckboxItem>
          <DropdownMenuCheckboxItem
            className="cursor-pointer"
            checked={status === "inactive"}
            onClick={() => statusQuery("inactive")}
          >Inactive</DropdownMenuCheckboxItem>
        </DropdownMenuContent>
      </DropdownMenu>
    )
  }

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

          <DropdownMenuLabel className="text-muted-foreground text-xs">
            Name
          </DropdownMenuLabel>
          <DropdownMenuCheckboxItem
            className="cursor-pointer"
            checked={rawSort === "name:asc"}
            onClick={() => sortQuery("name:asc")}
          >asc</DropdownMenuCheckboxItem>
          <DropdownMenuCheckboxItem
            className="cursor-pointer"
            checked={rawSort === "name:desc"}
            onClick={() => sortQuery("name:desc")}
          >desc</DropdownMenuCheckboxItem>

          <DropdownMenuLabel className="text-muted-foreground text-xs">
            Price
          </DropdownMenuLabel>
          <DropdownMenuCheckboxItem
            className="cursor-pointer"
            checked={rawSort === "price:asc"}
            onClick={() => sortQuery("price:asc")}
          >asc</DropdownMenuCheckboxItem>
          <DropdownMenuCheckboxItem
            className="cursor-pointer"
            checked={rawSort === "price:desc"}
            onClick={() => sortQuery("price:desc")}
          >desc</DropdownMenuCheckboxItem>

          {tab === "addons" && (
            <>
              <DropdownMenuLabel className="text-muted-foreground text-xs">
                Total Orders
              </DropdownMenuLabel>
              <DropdownMenuCheckboxItem
                className="cursor-pointer"
                checked={rawSort === "usage:asc"}
                onClick={() => sortQuery("usage:asc")}
              >asc</DropdownMenuCheckboxItem>
              <DropdownMenuCheckboxItem
                className="cursor-pointer"
                checked={rawSort === "usage:desc"}
                onClick={() => sortQuery("usage:desc")}
              >desc</DropdownMenuCheckboxItem>
            </>
          )}

          <DropdownMenuSeparator/>
          <DropdownMenuItem
            className="cursor-pointer"
            variant="destructive"
            onClick={() => sortQuery("")}
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
            <ProductContainer sort={query.sort} status={query.status} />
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
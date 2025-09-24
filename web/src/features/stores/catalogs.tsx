import {useEffect, useState} from "react";
import {Tabs, TabsContent, TabsList, TabsTrigger} from "@/components/ui/tabs";
import {ProductContainer} from "@/features/stores/components/catalog/product-container";
import {AddonContainer} from "@/features/stores/components/catalog/addon-container";
import {
  DropdownMenu, DropdownMenuCheckboxItem,
  DropdownMenuContent,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger
} from "@/components/ui/dropdown-menu";
import {Button} from "@/components/ui/button";
import {File, ListFilter, Plus} from "lucide-react";
import {useActionState} from "@/states/action-state";
import {useNavigate} from "@tanstack/react-router";
import {AddonActionAdd, AddonActionAddModalState} from "@/features/stores/components/catalog/addon-action-add";
import {ProductActionDeleteModal} from "@/features/stores/components/catalog/product-action-delete";
import {AddonActionEdit} from "@/features/stores/components/catalog/addon-action-edit";
import {AddonActionDelete} from "@/features/stores/components/catalog/addon-action-delete";
import {useProductAddonList} from "@/hooks/use-product";
import {useProductState} from "@/states/product-state";

export function ProductCatalogPage() {
  const { setBoolState } = useActionState();
  const [activeTab, setActiveTab] = useState("products");
  const navigate = useNavigate();
  const {data: addons} = useProductAddonList();
  const {setAddons} = useProductState();

  useEffect(() => {
    if (addons?.data) {
      setAddons(addons.data);
    }
  }, [addons?.data]);

  async function onAddNewItem() {
    if (activeTab === "products") {
      await navigate({to: "/stores/products"})
      return;
    }
    setBoolState(AddonActionAddModalState, true);
  }

  return (
    <div className="w-full pt-4 xl:pt-6 space-y-6">
      <aside className="flex items-center justify-between gap-4 px-8">
        <div className="block">
          <h3 className="text-lg font-semibold mb-2">Catalogs</h3>
          <p className="text-sm text-muted-foreground">
            Manage core catalog data such as products and addons. Organize and control the items available in your system.
          </p>
        </div>
      </aside>

      <aside className="mt-4 gap-6 px-8">
        <Tabs value={activeTab} onValueChange={setActiveTab} className="h-full w-full">
          <div className="flex items-center">
            <TabsList>
              <TabsTrigger value="products" className="flex items-center space-x-2 px-4">
                <span>Products</span>
              </TabsTrigger>
              <TabsTrigger value="addons" className="flex items-center space-x-2 px-4">
                <span>Addons</span>
              </TabsTrigger>
            </TabsList>

            <div className="ml-auto flex items-center gap-2">
              <Button
                size="sm"
                variant="outline"
                className="h-7 gap-1 text-sm"
                onClick={async (e) => {
                  e.preventDefault();
                  await onAddNewItem();
                }}
              >
                <Plus className="h-3.5 w-3.5"/>
                <span className="sr-only sm:not-sr-only capitalize">New {activeTab}</span>
              </Button>

              {activeTab === "products" && (
                <>
                  <DropdownMenu>
                    <DropdownMenuTrigger asChild>
                      <Button
                        variant="outline"
                        size="sm"
                        className="h-7 gap-1 text-sm"
                      >
                        <ListFilter className="h-3.5 w-3.5"/>
                        <span className="sr-only sm:not-sr-only">Filter</span>
                      </Button>
                    </DropdownMenuTrigger>
                    <DropdownMenuContent align="end">
                      <DropdownMenuLabel>Filter by</DropdownMenuLabel>
                      <DropdownMenuSeparator/>
                      <DropdownMenuCheckboxItem checked>
                        All
                      </DropdownMenuCheckboxItem>
                      <DropdownMenuCheckboxItem>
                        Active
                      </DropdownMenuCheckboxItem>
                      <DropdownMenuCheckboxItem>
                        Inactive
                      </DropdownMenuCheckboxItem>
                      <DropdownMenuCheckboxItem>
                        Draft
                      </DropdownMenuCheckboxItem>
                      <DropdownMenuCheckboxItem>
                        In Stock
                      </DropdownMenuCheckboxItem>
                      <DropdownMenuCheckboxItem>
                        Out of Stock
                      </DropdownMenuCheckboxItem>
                    </DropdownMenuContent>
                  </DropdownMenu>

                  <Button
                    size="sm"
                    variant="outline"
                    className="h-7 gap-1 text-sm"
                  >
                    <File className="h-3.5 w-3.5"/>
                    <span className="sr-only sm:not-sr-only">Export</span>
                  </Button>
                </>
              )}
            </div>
          </div>
          <TabsContent value="products">
            <ProductContainer />
            <ProductActionDeleteModal />
          </TabsContent>
          <TabsContent value="addons">
            <AddonContainer />
            <AddonActionAdd />
            <AddonActionEdit />
            <AddonActionDelete />
          </TabsContent>
        </Tabs>
      </aside>
    </div>
  )
}
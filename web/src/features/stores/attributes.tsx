import {useEffect, useState} from "react";
import {Tabs, TabsContent, TabsList, TabsTrigger} from "@/components/ui/tabs";
import {UnitContainer} from "@/features/stores/components/attribute/unit-container";
import {CategoryContainer} from "@/features/stores/components/attribute/category-container";
import {useAttributeState} from "@/states/attribute-state";
import {useCategoryList, useUnitList} from "@/hooks/use-attribute";

export function MasterDataPage() {
  const [activeTab, setActiveTab] = useState("units")
  const {data: units} = useUnitList()
  const {data: categories} = useCategoryList()
  const {setUnits, setCategories} =  useAttributeState();

  useEffect(() => {
    if (units?.data) {
      setUnits(units.data)
    }

    if (categories?.data) {
      setCategories(categories.data)
    }
  }, [units?.data, categories?.data]);

  return (
    <div className="w-full pt-4 xl:pt-6 space-y-6 select-none">
      <aside className="flex items-center justify-between gap-4 px-8">
        <div className="block">
          <h3 className="text-lg font-semibold mb-2">Attributes</h3>
          <p className="text-sm text-muted-foreground">
            Manage your store’s core data, including units, categories, subcategories, and more. Keep your system organized and consistent by configuring the essential building blocks used across other modules.
          </p>
        </div>
      </aside>

      <aside className="mt-4 gap-6 px-8">
        <Tabs value={activeTab} onValueChange={setActiveTab} className="h-full w-full">
          <TabsList>
            <TabsTrigger value="units" className="flex items-center space-x-2 px-4">
              Units
            </TabsTrigger>
            <TabsTrigger value="categories" className="flex items-center space-x-2 px-4">
              Categories
            </TabsTrigger>
          </TabsList>
          <TabsContent value="units">
            <UnitContainer />
          </TabsContent>
          <TabsContent value="categories">
            <CategoryContainer />
          </TabsContent>
        </Tabs>
      </aside>
    </div>
  )
}
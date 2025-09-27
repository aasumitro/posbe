import {useEffect, useState} from "react";
import {Tabs, TabsContent, TabsList, TabsTrigger} from "@/components/ui/tabs";
import {UnitContainer} from "@/features/stores/components/attribute/unit-container";
import {CategoryContainer} from "@/features/stores/components/attribute/category-container";
import {useAttributeState} from "@/states/attribute-state";
import {useCategoryList, useUnitList} from "@/hooks/use-attribute";
import {useNavigate, useSearch} from "@tanstack/react-router";
import {Route} from "@/routes/_authenticated/stores/route";

interface AttributePageProps {
  query: {
    tab: "units" | "categories";
    sort?: Record<string, "asc" | "desc">;
  };
}

export function MasterDataPage({ query }: AttributePageProps) {
  const [activeTab, setActiveTab] = useState<string>(query.tab);
  const { sort } = useSearch({from: '/_authenticated/stores/attributes'});
  const navigate = useNavigate()
  const {data: units} = useUnitList()
  const {data: categories} = useCategoryList()
  const {setUnits, setCategories} =  useAttributeState();

  useEffect(() => {
    if (units?.data) setUnits(units.data)
    if (categories?.data)  setCategories(categories.data);
  }, [units?.data, categories?.data]);

  const handleTabChange = async (value: string) => {
    setActiveTab(value);
    await navigate({
      from: Route.fullPath,
      search: {
        tab: value,
        sort: sort ? decodeURIComponent(sort) : undefined
      }
    });
  }

  return (
    <div className="w-full pt-4 xl:pt-6 space-y-6 select-none">
      <aside className="flex items-center justify-between gap-4 px-8">
        <div className="block">
          <h3 className="text-lg font-semibold mb-2">Attributes</h3>
          <p className="text-sm text-muted-foreground">
            Manage your store’s core data, including units, categories, subcategories, and more.
            Keep your system organized and consistent by configuring the essential building blocks used across other modules.
          </p>
        </div>
      </aside>

      <aside className="mt-4 gap-6 px-8">
        <Tabs
          value={activeTab}
          className="h-full w-full"
          onValueChange={handleTabChange}
        >
          <TabsList>
            <TabsTrigger value="units" className="flex items-center space-x-2 px-4 cursor-pointer">
              Units
            </TabsTrigger>
            <TabsTrigger value="categories" className="flex items-center space-x-2 px-4 cursor-pointer">
              Categories
            </TabsTrigger>
          </TabsList>
          <TabsContent value="units">
            <UnitContainer sort={query.sort} />
          </TabsContent>
          <TabsContent value="categories">
            <CategoryContainer sort={query.sort} />
          </TabsContent>
        </Tabs>
      </aside>
    </div>
  )
}
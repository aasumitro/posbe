import {type ReactElement, useEffect, useState} from "react";
import {Tabs, TabsContent, TabsList, TabsTrigger} from "@/components/ui/tabs";
import {UnitContainer} from "@/features/stores/components/attribute/unit-container";
import {CategoryContainer} from "@/features/stores/components/attribute/category-container";
import {useAttributeState} from "@/states/attribute-state";
import {useCategoryList, useUnitList} from "@/hooks/use-attribute";
import {useNavigate, useSearch} from "@tanstack/react-router";
import {Route} from "@/routes/_authenticated/stores/route";
import {
  DropdownMenu,
  DropdownMenuCheckboxItem,
  DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuSeparator,
  DropdownMenuTrigger
} from "@/components/ui/dropdown-menu";
import {Button} from "@/components/ui/button";
import {SortAscIcon} from "lucide-react";
import {IconSortAZ, IconSortZA} from "@tabler/icons-react";

interface AttributePageProps {
  query: {
    tab: "units" | "categories";
    sort?: Record<string, "asc" | "desc">;
  };
}

export function MasterDataPage({ query }: AttributePageProps) {
  const [activeTab, setActiveTab] = useState<string>(query.tab);
  const { sort: rawSort } = useSearch({from: '/_authenticated/stores/attributes'});
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
    await navigate({from: Route.fullPath, search: {tab: value}});
  }

  const updateSearchQuery = async (
    updates: { status?: string; category?: string; sort?: string }
  ) => {
    let search: { tab: string; sort?: string } = {
      tab: activeTab,
      sort: rawSort ? decodeURIComponent(rawSort) : undefined,
    };

    search = { ...search, ...updates };

    await navigate({from: Route.fullPath, search});
  };

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

  const SortDropdown = () => {
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
          <div className="flex items-center">
            <TabsList>
              <TabsTrigger value="units" className="flex items-center space-x-2 px-4 cursor-pointer">
                Units
              </TabsTrigger>
              <TabsTrigger value="categories" className="flex items-center space-x-2 px-4 cursor-pointer">
                Categories
              </TabsTrigger>
            </TabsList>

            <div className="ml-auto flex items-center gap-2">
              <SortDropdown />
            </div>
          </div>

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
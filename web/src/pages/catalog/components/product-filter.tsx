import {Popover, PopoverContent, PopoverTrigger} from "@/components/ui/popover.tsx";
import {Button} from "@/components/ui/button.tsx";
import {EraserIcon, SearchIcon, SlidersHorizontalIcon} from "lucide-react";
import {cn} from "@/lib/utils.ts";
import {Input} from "@/components/ui/input.tsx";
import {Label} from "@/components/ui/label.tsx";
import {useState} from "react";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select.tsx";

export const ProductFilter = () => {
  const [search, setSearch] = useState<string>("")
  const [filter, setFilter] = useState<string[]>([])

  return (
    <section className="flex">
      <Popover>
        <PopoverTrigger asChild>
          <Button variant="ghost">
            <SearchIcon className={cn(
              "w-4 h-4", {
                "mr-2": search.length > 0
              }
            )}/>
            {search}
          </Button>
        </PopoverTrigger>
        <PopoverContent>
          <Input
            type="search"
            placeholder="search . . ."
            onChange={(e) => setSearch(e.target.value)}
            value={search}
          />
        </PopoverContent>
      </Popover>

      <Popover>
        <PopoverTrigger asChild>
          <Button variant="ghost">
            <SlidersHorizontalIcon className={cn(
              "w-4 h-4", {
                "mr-2": filter.length > 0
              }
            )}/>
            {filter.join(", ")}
          </Button>
        </PopoverTrigger>
        <PopoverContent className="w-80">
          <div className="grid gap-4">
            <div className="space-y-2 flex items-center justify-between">
              <section>
                <h4 className="font-medium leading-none">Filter</h4>
                <p className="text-sm text-muted-foreground">
                  Set value to filter the products.
                </p>
              </section>
              <Button variant="ghost" onClick={() => setFilter([])}>
                <EraserIcon className="w-4 h-4" />
              </Button>
            </div>
            <div className="grid gap-2">
              <div className="grid grid-cols-3 items-center gap-4">
                <Label htmlFor="width">Category</Label>
                <Select onValueChange={(value) => setFilter([value])}>
                  <SelectTrigger className="col-span-2 h-8">
                    <SelectValue placeholder="Select category" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="foods">Foods</SelectItem>
                    <SelectItem value="drinks">Drinks</SelectItem>
                  </SelectContent>
                </Select>
              </div>
            </div>
          </div>
        </PopoverContent>
      </Popover>
    </section>
  )
}
import {Tabs, TabsContent, TabsList, TabsTrigger} from "@/components/ui/tabs.tsx";
import {Button} from "@/components/ui/button.tsx";
import {Circle, Edit3, Layers3, PlusIcon, SaveIcon, Square} from "lucide-react";
import {Card, CardContent, CardDescription, CardHeader, CardTitle} from "@/components/ui/card.tsx";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger
} from "@/components/ui/dropdown-menu.tsx";
import {useState} from "react";
import {Popover, PopoverContent, PopoverTrigger} from "@/components/ui/popover.tsx";
import {Input} from "@/components/ui/input.tsx";
import {Label} from "@/components/ui/label.tsx";
import {capitalize} from "@/lib/str.ts";
import {cn} from "@/lib/utils.ts";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select.tsx";

interface MapTabHeaderProp {
  edit: boolean

  editCallback(): void
}

function MapTabHeader(props: MapTabHeaderProp) {
  return (
    <div className="flex items-center">
      <TabsList>
        <TabsTrigger value="1st">1st Floor</TabsTrigger>
        <TabsTrigger value="2nd">2nd Floor</TabsTrigger>
      </TabsList>
      <div className="ml-auto flex items-center gap-2">
        {!props.edit && <Button size="sm" className="h-8 gap-1" onClick={props.editCallback}>
          <Edit3 className="h-3.5 w-3.5"/>
          <span className="sr-only sm:not-sr-only sm:whitespace-nowrap">
            Edit Map
          </span>
        </Button>}
        {props.edit && <>
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button size="sm" className="h-8 gap-1">
                <PlusIcon className="w-4 h-4"/>
                <span className="sr-only sm:not-sr-only sm:whitespace-nowrap">
                  New Items
                </span>
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent className="w-56">
              <DropdownMenuLabel>Map Content</DropdownMenuLabel>
              <DropdownMenuSeparator/>
              <DropdownMenuGroup>
                <DropdownMenuItem>
                  <Layers3 className="mr-2 h-4 w-4"/>
                  <span>Floor</span>
                </DropdownMenuItem>
                <DropdownMenuItem>
                  <Circle className="mr-2 h-4 w-4"/>
                  <span>Circle Room/Table</span>
                </DropdownMenuItem>
                <DropdownMenuItem>
                  <Square className="mr-2 h-4 w-4"/>
                  <span>Square Room/Table</span>
                </DropdownMenuItem>
              </DropdownMenuGroup>
            </DropdownMenuContent>
          </DropdownMenu>
          <Button size="sm" className="h-8 gap-1" onClick={props.editCallback}>
            <SaveIcon className="h-3.5 w-3.5"/>
            <span className="sr-only sm:not-sr-only sm:whitespace-nowrap">
              Save Map
            </span>
          </Button>
        </>}
      </div>
    </div>
  )
}

interface MapContent {
  id: number
  floor_id: number
  variant: string
  type: string
  name: string
  capacity: number
  x_pos: number
  y_pos: number
  w_size: number
  h_size: number
}

function RoomTableContent(prop: MapContent) {
  return (
    <Popover>
      <PopoverTrigger asChild>
        <Button
          className={cn(
            "absolute bg-gray-700",
            (prop.type === "circle" ? "rounded-full" : "rounded-sm"),
          )}
          style={{
            width: `${prop.w_size}px`,
            height: `${prop.h_size}px`,
            left: `${prop.x_pos}px`,
            top: `${prop.y_pos}px`
          }}
        >{prop.name}</Button>
      </PopoverTrigger>
      <PopoverContent side="right" className="w-80">
        <div className="grid gap-4">
          <div className="space-y-2">
            <h4 className="font-medium leading-none">
              {prop.name} ({capitalize(prop.variant)} - {capitalize(prop.type)})
            </h4>
            <p className="text-sm text-muted-foreground">
              Set the configuration for the {prop.variant}.
            </p>
          </div>
          <div className="grid gap-2">
            <div className="grid grid-cols-3 items-center gap-4">
              <Label htmlFor="name">Name</Label>
              <Input
                id="name"
                defaultValue={prop.name}
                className="col-span-2 h-8"
              />
            </div>
            <div className="grid grid-cols-3 items-center gap-4">
              <Label htmlFor="variant">Variant</Label>
              <Select value={prop.variant}>
                <SelectTrigger className="col-span-2 h-8">
                  <SelectValue placeholder="variant"/>
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="table">Table</SelectItem>
                  <SelectItem value="room">Room</SelectItem>
                </SelectContent>
              </Select>
            </div>
            <div className="grid grid-cols-3 items-center gap-4">
              <Label htmlFor="type">Type</Label>
              <Select value={prop.type}>
                <SelectTrigger className="col-span-2 h-8">
                  <SelectValue placeholder="type"/>
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="circle">Circle</SelectItem>
                  <SelectItem value="square">Square</SelectItem>
                </SelectContent>
              </Select>
            </div>
            <div className="grid grid-cols-3 items-center gap-4">
              <Label htmlFor="capacity">Capacity</Label>
              <Select value={String(prop.capacity)}>
                <SelectTrigger className="col-span-2 h-8">
                  <SelectValue placeholder="capacity"/>
                </SelectTrigger>
                <SelectContent>
                  {Array.from({length: 50}, (_, i) => i + 1).map((num) => (
                    <SelectItem key={num} value={String(num)}>
                      {num}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </div>
            <div className="grid grid-cols-3 items-center gap-4">
              <Label htmlFor="w_size">WSize</Label>
              <Input
                id="w_size"
                defaultValue={prop.w_size}
                type="number"
                min="30"
                max="60"
                className="col-span-2 h-8"
              />
            </div>
            <div className="grid grid-cols-3 items-center gap-4">
              <Label htmlFor="h_size">HSize</Label>
              <Input
                id="h_size"
                defaultValue={prop.h_size}
                type="number"
                min="30"
                max="60"
                className="col-span-2 h-8"
              />
            </div>
            <div className="grid grid-cols-3 items-center gap-4">
              <Label htmlFor="x_pos">XPos</Label>
              <Input
                id="x_pos"
                defaultValue={prop.x_pos}
                className="col-span-2 h-8"
              />
            </div>
            <div className="grid grid-cols-3 items-center gap-4">
              <Label htmlFor="y_pos">YPos</Label>
              <Input
                id="y_pos"
                defaultValue={prop.y_pos}
                className="col-span-2 h-8"
              />
            </div>
          </div>
          <div className="flex justify-between">
            <Button variant="destructive">Delete</Button>
            <Button>Save</Button>
          </div>
        </div>
      </PopoverContent>
    </Popover>
  )
}

const contents: MapContent[] = [
  {
    id: 1,
    variant: "table",
    type: "square",
    floor_id: 1,
    name: "T1",
    x_pos: 10,
    y_pos: 10,
    w_size: 60,
    h_size: 30,
    capacity: 4
  },
  {
    id: 2,
    variant: "table",
    type: "circle",
    floor_id: 1,
    name: "T2",
    x_pos: 100,
    y_pos: 100,
    w_size: 40,
    h_size: 40,
    capacity: 4
  },
  {
    id: 3,
    variant: "room",
    type: "square",
    floor_id: 1,
    name: "R2",
    x_pos: 250,
    y_pos: 250,
    w_size: 60,
    h_size: 60,
    capacity: 12
  }
]

function MapTabContent() {
  return (
    <TabsContent value="1st">
      <Card className="min-h-full">
        <CardHeader>
          <CardTitle>Store Maps</CardTitle>
          <CardDescription>
            Manage your store floor, table and room.
          </CardDescription>
        </CardHeader>
        <CardContent className="mx-auto py-6">
          <div className="relative h-96 overflow-hidden rounded-xl border border-dashed border-gray-400 opacity-75">
            <svg className="absolute inset-0 h-full w-full stroke-gray-900/10" fill="none">
              <defs>
                <pattern
                  id="pattern-d09edaee-fc6a-4f25-aca5-bf9f5f77e14a"
                  x="0" y="0" width="10" height="10"
                  patternUnits="userSpaceOnUse"
                >
                  <path d="M-3 13 15-5M-5 5l18-18M-1 21 17 3"></path>
                </pattern>
              </defs>
              <rect
                stroke="none" width="100%" height="100%"
                fill="url(#pattern-d09edaee-fc6a-4f25-aca5-bf9f5f77e14a)"
              ></rect>
            </svg>
            {contents.map((item, index) => (
              <RoomTableContent {...item} key={index}/>
            ))}
          </div>
        </CardContent>
      </Card>
    </TabsContent>
  )
}

export function Map() {
  const [edit, setEdit] = useState(false)

  const saveMap = () => {
    // maybe need to save all
    // then get all the data again
    setEdit(!edit)
  }

  return (<>
    <main className="grid flex-1 items-start gap-4 p-4 sm:px-6 sm:py-0 md:gap-8">
      <Tabs defaultValue="1st">
        <MapTabHeader edit={edit} editCallback={saveMap}/>
        <MapTabContent/>
      </Tabs>
    </main>
  </>)
}

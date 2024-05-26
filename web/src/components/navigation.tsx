import {
  Home,
  Layers3,
  Package,
  Package2,
  PanelLeft,
  Search,
  Settings,
  ShoppingCart,
  Users2
} from "lucide-react";
import {Tooltip, TooltipContent, TooltipProvider, TooltipTrigger} from "@/components/ui/tooltip.tsx";
import {Button} from "@/components/ui/button.tsx";
import {Sheet, SheetContent, SheetTrigger} from "@/components/ui/sheet.tsx";
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList, BreadcrumbPage,
  BreadcrumbSeparator
} from "@/components/ui/breadcrumb.tsx";
import {Input} from "@/components/ui/input.tsx";
import {
  DropdownMenu,
  DropdownMenuContent, DropdownMenuItem,
  DropdownMenuLabel, DropdownMenuSeparator,
  DropdownMenuTrigger
} from "@/components/ui/dropdown-menu.tsx";
import {Avatar, AvatarFallback, AvatarImage} from "@/components/ui/avatar.tsx";
import {useLocation} from 'react-router-dom';
import {Fragment} from "react";
import {capitalize} from "@/lib/str.ts";
import {cn} from "@/lib/utils.ts";

const menus = [
  {
    name: "Home",
    route: "/",
    icon: <Home className="h-5 w-5"/>
  },
  {
    name: "Store Maps",
    route: "/fe/maps",
    icon: <Layers3 className="h-5 w-5"/>
  },
  {
    name: "Transactions",
    route: "/fe/transactions",
    icon: <ShoppingCart className="h-5 w-5"/>
  },
  {
    name: "Products",
    route: "/fe/products",
    icon: <Package className="h-5 w-5"/>
  },
  {
    name: "Users",
    route: "/fe/users",
    icon: <Users2 className="h-5 w-5"/>
  },
  {
    name: "Settings",
    route: "/fe/settings",
    icon: <Settings className="h-5 w-5"/>
  }
]

function topNav() {
  const location = useLocation();
  const currentPath = location.pathname;
  return (
    <nav className="flex flex-col items-center gap-4 px-2 sm:py-5">
      <a href="/"
         className="group flex h-9 w-9 shrink-0 items-center justify-center gap-2 rounded-full bg-primary text-lg font-semibold text-primary-foreground md:h-8 md:w-8 md:text-base">
        <Package2 className="h-4 w-4 transition-all group-hover:scale-110"/>
        <span className="sr-only">POSBE</span>
      </a>
      <TooltipProvider>
        {menus.filter((menu) =>
          menu.name.toLowerCase() !== "settings"
        ).map((menu, index) => (
          <Tooltip key={index}>
            <TooltipTrigger asChild>
              <a
                href={menu.route}
                className={cn(
                  "flex h-9 w-9 items-center justify-center rounded-lg  transition-colors hover:text-foreground md:h-8 md:w-8",
                  currentPath.split("/").includes(menu.name.toLowerCase()) ||
                  (currentPath === "/maps" && menu.name.toLowerCase() === "store maps") ||
                  (currentPath === "/" && menu.name.toLowerCase() === "home")
                    ? " bg-accent text-accent-foreground"
                    : "text-muted-foreground"
                )}
              >
                {menu.icon}
                <span className="sr-only">{menu.name}</span>
              </a>
            </TooltipTrigger>
            <TooltipContent side="right">
              {menu.name}

            </TooltipContent>
          </Tooltip>
        ))}
      </TooltipProvider>
    </nav>
  )
}

function Sidebar() {
  const location = useLocation();
  const currentPath = location.pathname;
  return (
    <aside className="fixed inset-y-0 left-0 z-10 hidden w-14 flex-col border-r bg-background sm:flex">
      {topNav()}
      <nav className="mt-auto flex flex-col items-center gap-4 px-2 sm:py-5">
        <TooltipProvider>
          <Tooltip>
            <TooltipTrigger asChild>
              <a
                href="/fe/settings"
                className={cn(
                  "flex h-9 w-9 items-center justify-center rounded-lg  transition-colors hover:text-foreground md:h-8 md:w-8",
                  (currentPath === "/settings")
                    ? " bg-accent text-accent-foreground"
                    : "text-muted-foreground"
                )}
              >
                <Settings className="h-5 w-5"/>
                <span className="sr-only">Settings</span>
              </a>
            </TooltipTrigger>
            <TooltipContent side="right">Settings</TooltipContent>
          </Tooltip>
        </TooltipProvider>
      </nav>
    </aside>
  )
}

function menu() {
  const location = useLocation();
  const currentPath = location.pathname;
  return (
    <Sheet>
      <SheetTrigger asChild>
        <Button size="icon" variant="outline" className="sm:hidden">
          <PanelLeft className="h-5 w-5"/>
          <span className="sr-only">Toggle Menu</span>
        </Button>
      </SheetTrigger>
      <SheetContent side="left" className="sm:max-w-xs">
        <nav className="grid gap-6 text-lg font-medium">
          <a
            href="/"
            className="group flex h-10 w-10 shrink-0 items-center justify-center gap-2 rounded-full bg-primary text-lg font-semibold text-primary-foreground md:text-base"
          >
            <Package2 className="h-5 w-5 transition-all group-hover:scale-110"/>
            <span className="sr-only">POSBE</span>
          </a>
          {menus.map((menu, index) => (
            //
            <a
              key={index}
              href={menu.route}
              className={cn(
                currentPath.split("/").includes(menu.name.toLowerCase()) ||
                (currentPath === "/maps" && menu.name.toLowerCase() === "store maps") ||
                (currentPath === "/" && menu.name.toLowerCase() === "home")
                  ? " flex items-center gap-4 px-2.5 text-foreground"
                  : "flex items-center gap-4 px-2.5 text-muted-foreground hover:text-foreground",
              )}
            >
              {menu.icon}
              {menu.name}
            </a>
          ))}
        </nav>
      </SheetContent>
    </Sheet>
  )
}

function breadcrumb() {
  const location = useLocation();
  const currentPath = location.pathname;
  const isHome = currentPath === "/";
  return (
    <Breadcrumb className="hidden md:flex">
      <BreadcrumbList>
        {isHome && <BreadcrumbItem>
          <BreadcrumbPage>Home</BreadcrumbPage>
        </BreadcrumbItem>}
        {!isHome && currentPath.split("/").map((path, index, arr) => {
          const fullPath = arr.slice(0, index + 1).join("/") || "/";
          const isLast = index === arr.length - 1;
          return (
            <Fragment key={index}>
              <BreadcrumbItem>
                {index === 0 && (
                  <BreadcrumbLink asChild>
                    <a href="/">Home</a>
                  </BreadcrumbLink>
                )}
                {(index > 1 && !isLast) && (
                  <BreadcrumbLink asChild>
                    <a href={fullPath}>{capitalize(path) || "Home"}</a>
                  </BreadcrumbLink>
                )}
                {isLast && <BreadcrumbPage>
                  {path === "" ? "Home" : capitalize(path)}
                </BreadcrumbPage>}
              </BreadcrumbItem>
              {!isLast && <BreadcrumbSeparator/>}
            </Fragment>
          );
        })}
      </BreadcrumbList>
    </Breadcrumb>
  )
}

function search() {
  return (
    <div className="relative ml-auto flex-1 md:grow-0">
      <Search className="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground"/>
      <Input
        type="search"
        placeholder="Search..."
        className="w-full rounded-lg bg-background pl-8 md:w-[200px] lg:w-[336px]"
      />
    </div>
  )
}

function user() {
  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button
          variant="outline"
          size="icon"
          className="overflow-hidden rounded-full"
        >
          <Avatar className="overflow-hidden rounded-full">
            <AvatarImage src="https://github.com/shadcn.png"/>
            <AvatarFallback>CN</AvatarFallback>
          </Avatar>
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end">
        <DropdownMenuLabel>My Account</DropdownMenuLabel>
        <DropdownMenuSeparator/>
        <DropdownMenuItem>Profile</DropdownMenuItem>
        <DropdownMenuSeparator/>
        <DropdownMenuItem>Logout</DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  )
}

function Header() {
  return (
    <header
      className="sticky top-0 z-30 flex h-14 items-center gap-4 border-b bg-background px-4 sm:static sm:h-auto sm:border-0 sm:bg-transparent sm:px-6">
      {menu()}

      {breadcrumb()}

      {search()}

      {user()}
    </header>
  )
}

export {
  Sidebar,
  Header
}

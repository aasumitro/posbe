import {Home,  Layers3, Package, ShoppingCart, StoreIcon} from "lucide-react";
import {Button} from "@/components/ui/button.tsx";
import {useNavigate} from "react-router-dom";
import {isCurrentRoute} from "@/lib/path-validation.ts";
import {useRef} from "react";
import {KEYS} from "@/lib/keys.ts";

const menu = [
  {
    name: "Home",
    route: "/home",
    access: ["*"],
    icon: <Home className="h-4 w-4"/>
  },
  {
    name: "Layouts",
    route: "/layouts",
    access: ["*"],
    icon: <Layers3 className="h-4 w-4"/>
  },
  {
    name: "Transactions",
    route: "/transactions",
    access: ["admin", "cashier"],
    icon: <ShoppingCart className="h-4 w-4"/>
  },
  {
    name: "Catalogs",
    route: "/catalogs",
    access: ["admin"],
    icon: <Package className="h-4 w-4"/>
  },
  {
    name: "Store",
    route: "/store",
    access: ["admin"],
    icon:  <StoreIcon className="h-4 w-4"/>
  }
]

export const MainNavigationMenu = () => {
  const navigate = useNavigate();
  const role = useRef("")

  const getUserRole = () => {
    const key = KEYS.UI_USER_PROFILE
    const profile = localStorage.getItem(key);
    if (profile) {
      const user = JSON.parse(profile);
      role.current = user.role.name;
    }
  }

  const filteredMenu = menu.filter(item => {
    if (item.access.includes("*")) return true;
    getUserRole();
    return item.access.includes(role.current);
  });

  return (
    <div id="mega-menu-full" className="items-center justify-between font-medium hidden w-full lg:flex lg:w-auto">
      <ul
        className="flex flex-col p-4 md:p-0 mt-4 border border-gray-100 rounded-lg bg-gray-50 md:space-x-8 rtl:space-x-reverse md:flex-row md:mt-0 md:border-0 md:bg-white dark:bg-gray-800 md:dark:bg-gray-900 dark:border-gray-700"
      >
        {filteredMenu.map((item, index) => (
          <li key={index}>
            <Button
              variant={isCurrentRoute(item.route) ? "default" : "ghost"}
              onClick={() => navigate(item.route)}
              className="gap-2"
            >
              {item.icon}
              {item.name}
            </Button>
          </li>
        ))}
      </ul>
    </div>
  )
}
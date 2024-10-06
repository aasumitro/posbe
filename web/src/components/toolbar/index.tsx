import {UserNavigationMenu} from "@/components/toolbar/user.tsx";
import {MainNavigationMenu} from "@/components/toolbar/menu.tsx";
import {MenuToggle} from "@/components/toolbar/menu-toggle.tsx";
import {ShiftSwitcherNavigation} from "@/components/toolbar/shift.tsx";

export const Toolbar = () => {
  return (
    <nav className="bg-white border-gray-200 dark:border-gray-600 dark:bg-gray-900 shadow-sm border-y px-8 py-6">
      <div className="flex flex-wrap justify-between items-center mx-auto">
        <div className="w-[200px]">
          <ShiftSwitcherNavigation/>
        </div>
        <MenuToggle/>
        <MainNavigationMenu/>
        <div className="hidden lg:flex items-center w-[200px]">
          <UserNavigationMenu/>
        </div>
      </div>
    </nav>
  )
}
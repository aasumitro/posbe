import {FolderIcon} from "lucide-react";

export const ShiftSwitcherNavigation = () => {
  return (
    <div className="flex items-center space-x-3 rtl:space-x-reverse">
      <FolderIcon className="w-6 h-6"/>
      <span className="self-center text-xl font-semibold whitespace-nowrap dark:text-white">
        Lorem Store!
      </span>
    </div>
  )
}
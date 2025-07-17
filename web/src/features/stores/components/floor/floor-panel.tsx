import {cn} from "@/lib/utils";
import {ChevronDown} from "lucide-react";

export function FloorPanel() {
  return (
    <div
      className={cn(
        "absolute top-16 right-5.5 select-none px-4 py-2 flex items-center justify-between w-28",
        "rounded-lg text-muted-foreground bg-gray-100 font-mono text-[10px] cursor-not-allowed gap-2",
      )}
      onClick={(e) => {
        e.preventDefault();
        alert("Floor feature currently not available");
      }}
    >
      1st Floor
      <ChevronDown className="w-3.5 h-3.5" />
    </div>
  )
}
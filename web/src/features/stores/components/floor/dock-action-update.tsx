import {Dock, DockIcon} from "@/components/ui/dock";
import {Tooltip, TooltipContent, TooltipProvider, TooltipTrigger} from "@/components/ui/tooltip";
import {Separator} from "@/components/ui/separator";
import {Save, SaveOff} from "lucide-react";

export function UpdateActionDock({revert, save}: ({revert: () => void, save: () => void})) {
  return (
    <TooltipProvider>
      <Dock direction="middle" className="absolute bottom-2 left-0 right-0">
        <Tooltip>
          <TooltipTrigger asChild>
            <DockIcon onTap={revert}>
              <SaveOff className="size-4" />
            </DockIcon>
          </TooltipTrigger>
          <TooltipContent>
            <p>Revert changes</p>
          </TooltipContent>
        </Tooltip>

        <Separator orientation="vertical" className="h-full py-2" />

        <p className="px-6">Save changes?</p>

        <Separator orientation="vertical" className="h-full py-2" />

        <Tooltip>
          <TooltipTrigger asChild>
            <DockIcon onTap={save}>
              <Save className="size-4" />
            </DockIcon>
          </TooltipTrigger>
          <TooltipContent>
            <p>Save and  Apply</p>
          </TooltipContent>
        </Tooltip>
      </Dock>
    </TooltipProvider>
  )
}
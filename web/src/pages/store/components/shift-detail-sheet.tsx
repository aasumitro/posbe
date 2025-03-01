import {useEffect, useState} from "react";
import {useGlobalStateStore} from "@/stores/global-state.ts";
import {Sheet, SheetContent, SheetDescription, SheetHeader, SheetTitle, SheetTrigger} from "@/components/ui/sheet.tsx";
import {KEYS} from "@/lib/keys.ts";
import {Toggle} from "@/components/ui/toggle.tsx";
import {EyeIcon, UserPen} from "lucide-react";

export function ShiftDetailSheet() {
  const [open, setOpen] = useState(false)
  const {boolStates, setBoolState} = useGlobalStateStore();
  const [editMode, setEditMode] = useState(false)

  useEffect(() => {
    if (boolStates[KEYS.DISPLAY_SHIFT_DETAIL_SHEET]) {
      setOpen(true);
    }
  }, [boolStates])

  const onClose = () => {
    setOpen(!open);
    setBoolState(KEYS.DISPLAY_SHIFT_DETAIL_SHEET, false);
  }

  const onModeChange = () => {
    setEditMode(!editMode)
  }

  return (
    <Sheet open={open} onOpenChange={onClose}>
      <SheetTrigger asChild></SheetTrigger>
      <SheetContent side="right">
        <SheetHeader>
          <SheetTitle>
            {editMode ? "Edit Shift" : "Shift Detail"}
            <Toggle
              className="mt-4 ml-4"
              size="sm"
              onClick={onModeChange}
            >
              {!editMode ? <>
                <EyeIcon className="mr-2 h-4 w-4"/>
                View Mode
              </> : <>
                <UserPen className="mr-2 h-4 w-4"/>
                Edit Mode
              </>}
            </Toggle>
          </SheetTitle>
          <SheetDescription>
            {editMode
              ? "Make changes to your shift here. Click save when you're done."
              : "Current selected shift detail. "
            }
          </SheetDescription>
        </SheetHeader>
      </SheetContent>
    </Sheet>
  )
}
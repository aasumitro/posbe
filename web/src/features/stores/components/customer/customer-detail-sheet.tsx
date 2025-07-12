import {useEffect, useState} from "react";
import {useActionState} from "@/states/action-state";
import {Sheet, SheetContent, SheetDescription, SheetHeader, SheetTitle} from "@/components/ui/sheet";
import {Label} from "@/components/ui/label";
import {Input} from "@/components/ui/input";
import {Button} from "@/components/ui/button";

export const CustomerDetailActionSheetState = "customer_detail_action_sheet_state"

export function CustomerDetailSheet() {
  const [open, setOpen] = useState(false);
  const { bool, setBoolState } = useActionState();

  useEffect(() => {
    if (bool[CustomerDetailActionSheetState]){
      setOpen(bool[CustomerDetailActionSheetState])
    }
  }, [bool]);

  function onOpenChange(newOpen: boolean) {
    setOpen(newOpen);
    if (!newOpen) {
      setBoolState(CustomerDetailActionSheetState, false);
    }
  }

  return (
    <Sheet open={open} onOpenChange={onOpenChange}>
      <SheetContent>
        <SheetHeader>
          <SheetTitle>Customer Info</SheetTitle>
          <SheetDescription>
            Preview customer details or remove them from the list.
          </SheetDescription>
        </SheetHeader>

        <div className="px-4 w-full space-y-2">
          <Label>Name</Label>
          <Input id="name" value="Zeros Mardigu" disabled />
        </div>

        <div className="px-4 w-full space-y-2">
          <Label>Phone number</Label>
          <Input id="phone" value="+6282275558899" disabled />
        </div>

        <div className="flex justify-between items-center px-4 mt-2">
          <Button
            variant="link"
            className="text-red-500"
          >
            DELETE
          </Button>

          <div className="space-x-2">
            <Button variant="outline">Cancel</Button>
            <Button>Save</Button>
          </div>
        </div>
      </SheetContent>
    </Sheet>
  )
}
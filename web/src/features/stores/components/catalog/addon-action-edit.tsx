import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "@/components/ui/alert-dialog"
import {useEffect, useState} from "react";
import {useActionState} from "@/states/action-state";

export const AddonActionEditModalState = "addon_action_edit_modal_state"

export function AddonActionEdit() {
  const [open, setOpen] = useState(false);
  const { bool, setBoolState } = useActionState();

  useEffect(() => {
    if (bool[AddonActionEditModalState]){
      setOpen(bool[AddonActionEditModalState])
    }
  }, [bool]);

  function onOpenChange(newOpen: boolean) {
    setOpen(newOpen);
    if (!newOpen) {
      setBoolState(AddonActionEditModalState, false);
    }
  }

  return (
    <AlertDialog open={open} onOpenChange={onOpenChange}>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Edit Addon</AlertDialogTitle>
          <AlertDialogDescription>
            Tests
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel>Cancel</AlertDialogCancel>
          <AlertDialogAction>Continue</AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  )
}
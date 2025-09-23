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

export const AddonActionAddModalState = "addon_action_add_modal_state"

export function AddonActionAdd() {
  const [open, setOpen] = useState(false);
  const { bool, setBoolState } = useActionState();

  useEffect(() => {
    if (bool[AddonActionAddModalState]){
      setOpen(bool[AddonActionAddModalState])
    }
  }, [bool]);

  function onOpenChange(newOpen: boolean) {
    setOpen(newOpen);
    if (!newOpen) {
      setBoolState(AddonActionAddModalState, false);
    }
  }

  return (
    <AlertDialog open={open} onOpenChange={onOpenChange}>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Add new Addon</AlertDialogTitle>
          <AlertDialogDescription>
            Test
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
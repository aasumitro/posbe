import {useEffect, useState} from "react";
import {useGlobalStateStore} from "@/stores/global-state.ts";
import {KEYS} from "@/lib/keys.ts";
import {
  AlertDialog, AlertDialogContent,
  AlertDialogDescription,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger
} from "@/components/ui/alert-dialog.tsx";

export function NewShiftModal() {
  const [open, setOpen] = useState(false)
  const {boolStates, setBoolState} = useGlobalStateStore();

  useEffect(() => {
    if (boolStates[KEYS.DISPLAY_NEW_SHIFT_MODAL]) {
      setOpen(true);
    }
  }, [boolStates])

  const onClose = () => {
    setOpen(!open);
    setBoolState(KEYS.DISPLAY_NEW_SHIFT_MODAL, false);
  }

  return (
    <AlertDialog open={open} onOpenChange={onClose}>
      <AlertDialogTrigger asChild></AlertDialogTrigger>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Create new Shift</AlertDialogTitle>
          <AlertDialogDescription>
            You are about to create a new store shift. Please make sure that all the provided details are accurate.
          </AlertDialogDescription>
        </AlertDialogHeader>
      </AlertDialogContent>
    </AlertDialog>
  )
}
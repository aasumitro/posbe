"use client"

import {useActionState} from "@/states/action-state";
import {useEffect, useState} from "react";
import {
  AlertDialog,
  // AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  // AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle
} from "@/components/ui/alert-dialog";
// import {Button} from "@/components/ui/button";
// import {Loader} from "lucide-react";

export const CloseShiftModalState = "close_shift_modal_state"

export function ShiftActionCloseAlertDialog() {
  const [isDialogOpen, setDialogOpen] = useState(false);
  const { bool, setBoolState } = useActionState();

  useEffect(() => {
    if (bool[CloseShiftModalState]) setDialogOpen(bool[CloseShiftModalState])
  }, [bool]);

  // const onSubmit = (e: Event) => {
  //   e.preventDefault();
  // }

  function onOpenChange(newOpen: boolean) {
    setDialogOpen(newOpen);
    if (!newOpen) {
      setBoolState(CloseShiftModalState, false);
      // form.reset();
    }
  }

  return (
    <AlertDialog open={isDialogOpen} onOpenChange={onOpenChange}>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>
             Open Shift
          </AlertDialogTitle>
          <AlertDialogDescription>
            You need to start a new shift before creating or processing orders.
          </AlertDialogDescription>
        </AlertDialogHeader>
        {/*<form onSubmit={onSubmit}>*/}
        {/*  <AlertDialogFooter>*/}
        {/*    <AlertDialogCancel*/}
        {/*      onClick={onClose}*/}
        {/*      disabled={isPending}*/}
        {/*    >Cancel</AlertDialogCancel>*/}
        {/*    <Button*/}
        {/*      className="bg-red-500 hover:bg-red-600 text-white"*/}
        {/*      type="submit"*/}
        {/*      disabled={isSubmitted || isPending}*/}
        {/*    >*/}
        {/*      <Loader className={*/}
        {/*        isSubmitted*/}
        {/*          ? "block animate-spin"*/}
        {/*          : "hidden"*/}
        {/*      }/>*/}
        {/*      Confirm*/}
        {/*    </Button>*/}
        {/*  </AlertDialogFooter>*/}
        {/*</form>*/}
      </AlertDialogContent>
    </AlertDialog>
  )
}
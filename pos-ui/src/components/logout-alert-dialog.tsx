"use client"

import {
  AlertDialog,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "@/components/ui/alert-dialog";
import { useState, useEffect } from "react";
import { Loader} from "lucide-react";
import { Button } from "@/components/ui/button";
import { useActionState } from "@/states/action-state";

export const LogoutModalState = "logout_modal_state"

export function LogoutAlertDialog() {
  const [isDialogOpen, setDialogOpen] = useState(false);
  const [isSubmitted, setSubmitted] = useState(false);
  const { bool, setBoolState } = useActionState();

  useEffect(() => {
    if (bool[LogoutModalState]){
      setDialogOpen(bool[LogoutModalState])
    }
  }, [bool]);

  const onSubmit = () => {
    setSubmitted(true);
  };

  const onClose = () => {
    setBoolState(LogoutModalState, false)
    setDialogOpen(false)
    setSubmitted(false);
  }

  return (
    <AlertDialog open={isDialogOpen} onOpenChange={onClose}>
      <AlertDialogContent className="w-96">
        <AlertDialogHeader>
          <AlertDialogTitle>
            Confirm Logout
          </AlertDialogTitle>
          <AlertDialogDescription>
            Are you sure you want to log out?
          </AlertDialogDescription>
        </AlertDialogHeader>
        <form onSubmit={onSubmit}>
          <AlertDialogFooter>
            <AlertDialogCancel
              onClick={onClose}
            >Cancel</AlertDialogCancel>
            <Button
              className="bg-red-500 hover:bg-red-600 text-white"
              type="submit"
            >
              <Loader className={
                isSubmitted
                  ? "block animate-spin"
                  : "hidden"
              }/>
              Confirm
            </Button>
          </AlertDialogFooter>
        </form>
      </AlertDialogContent>
    </AlertDialog>
  );
}
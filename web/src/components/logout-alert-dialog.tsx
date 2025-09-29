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
import { useState, useEffect, type FormEvent } from "react";
import { Loader} from "lucide-react";
import { Button } from "@/components/ui/button";
import { useActionState } from "@/states/action-state";
import {useAuthStore} from "@/states/auth-state";
import {useLogout} from "@/hooks/use-auth";
import {useNavigate} from "@tanstack/react-router";
import {toast} from "sonner";
import {HTTP_STATUS_CODE, isHTTPResponse} from "@/lib/api";

export const LogoutModalState = "logout_modal_state"

export function LogoutAlertDialog() {
  const [isDialogOpen, setDialogOpen] = useState(false);
  const [isSubmitted, setSubmitted] = useState(false);
  const { bool, setBoolState } = useActionState();
  const { mutate: logout, isPending} = useLogout();
  const navigate = useNavigate();

  useEffect(() => {
    if (bool[LogoutModalState]){
      setDialogOpen(bool[LogoutModalState])
    }
  }, [bool]);

  const onSubmit = (event: FormEvent) => {
    event.preventDefault();
    setSubmitted(true);

    logout(undefined, {
      onError: async (error) => {
        setSubmitted(false);
        if (error instanceof Error) {
          const clientError = error as Error
          if (clientError.message.includes("401")) {
            toast.success("Logout successfully!");
          }
          await toSignIn();
        }
        if (error && isHTTPResponse<null>(error) && error.code === HTTP_STATUS_CODE.UNAUTHORIZED) {
          toast.success("Logout successfully!");
          await toSignIn();
        }
      }
    })
  };

  const onClose = () => {
    setBoolState(LogoutModalState, false)
    setDialogOpen(false)
    setSubmitted(false);
  }

  const toSignIn = async () => {
    onClose();
    useAuthStore.getState().auth.reset();
    await navigate({to: "/login", replace: true})
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
              disabled={isPending}
            >Cancel</AlertDialogCancel>
            <Button
              className="bg-red-500 hover:bg-red-600 text-white"
              type="submit"
              disabled={isSubmitted || isPending}
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
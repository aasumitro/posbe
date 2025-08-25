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
import {toast} from "sonner";
import {useUserState} from "@/states/user-state";
import {useDeleteUser} from "@/hooks/use-user";
import {isHTTPResponse} from "@/lib/api";
import {useQueryClient} from "@tanstack/react-query";

export const ConfirmDeleteModalState = "team_confirm_delete_modal_state"

export function TeamConfirmDeleteAlertDialog() {
  const [isDialogOpen, setDialogOpen] = useState(false);
  const [isSubmitted, setSubmitted] = useState(false);
  const { bool, setBoolState } = useActionState();
  const { selectedUser } =  useUserState();
  const { mutate: deleteUser, isPending} = useDeleteUser();
  const queryClient = useQueryClient();

  useEffect(() => {
    if (bool[ConfirmDeleteModalState]){
      setDialogOpen(bool[ConfirmDeleteModalState])
    }
  }, [bool]);

  const onSubmit = (event: FormEvent) => {
    event.preventDefault();
    setSubmitted(true);

    if (!selectedUser) return;

    deleteUser(selectedUser?.id, {
      onSuccess: async () => {
        await queryClient.invalidateQueries({ queryKey: ['users'] })
        toast.success("User delete successfully");
        onClose();
      },
      onError: async (error) => {
        if (error && isHTTPResponse<null>(error)) {
          toast.error(error.data);
          return;
        }
        if (error instanceof Error) {
          const clientError = error as Error
          toast.error(clientError.message);
        }
      }
    })
  };

  const onClose = () => {
    setBoolState(ConfirmDeleteModalState, false)
    setDialogOpen(false)
    setSubmitted(false);
  }

  return (
    <AlertDialog open={isDialogOpen} onOpenChange={onClose}>
      <AlertDialogContent className="w-96">
        <AlertDialogHeader>
          <AlertDialogTitle>
            Confirm Delete
          </AlertDialogTitle>
          <AlertDialogDescription>
            Are you sure you want to delete this user ({selectedUser?.name})?
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
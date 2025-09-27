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
import {type FormEvent, useEffect, useState} from "react";
import {useActionState} from "@/states/action-state";
import {useQueryClient} from "@tanstack/react-query";
import {useProductState} from "@/states/product-state";
import {useDeleteProductAddon} from "@/hooks/use-product";
import {toast} from "sonner";
import {isHTTPResponse} from "@/lib/api";

export const AddonActionDeleteModalState = "addon_action_delete_modal_state"

export function AddonActionDelete() {
  const [open, setOpen] = useState(false);
  const { bool, setBoolState } = useActionState();
  const { selectedAddon, setSelectedAddon } =  useProductState();
  const { mutate: deleteAddon, isPending} = useDeleteProductAddon();
  const queryClient = useQueryClient();

  useEffect(() => {
    if (!selectedAddon) return;

    if (bool[AddonActionDeleteModalState]){
      setOpen(bool[AddonActionDeleteModalState])
    }
  }, [selectedAddon, bool]);

  function onOpenChange(newOpen: boolean) {
    setOpen(newOpen);
    if (!newOpen) {
      setBoolState(AddonActionDeleteModalState, false);
    }
  }

  const onSubmit = (e: FormEvent) => {
    e.preventDefault();
    if (!selectedAddon) {
      toast.warning("Please retry open the confirm delete button");
      onOpenChange(false);
      return;
    }

    deleteAddon(selectedAddon.id, {
      onSuccess: async () => {
        toast.success("Addon delete successfully");
        await queryClient.invalidateQueries({ queryKey: ['addons'] })
        setSelectedAddon(null);
        onOpenChange(false);
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
      },
    })
  }

  return (
    <AlertDialog open={open} onOpenChange={onOpenChange}>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Are you absolutely sure?</AlertDialogTitle>
          <AlertDialogDescription>
            This action cannot be undone.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <form onSubmit={onSubmit}>
          <AlertDialogFooter>
            <AlertDialogCancel
              type="button"
              disabled={isPending}
            >Cancel</AlertDialogCancel>
            <AlertDialogAction
              className="bg-red-500 hover:bg-red-600 text-white"
              type="submit"
              disabled={isPending}
            >Delete</AlertDialogAction>
          </AlertDialogFooter>
        </form>
      </AlertDialogContent>
    </AlertDialog>
  )
}
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
import {toast} from "sonner";
import {useQueryClient} from "@tanstack/react-query";
import {isHTTPResponse} from "@/lib/api";
import {useAttributeState} from "@/states/attribute-state";
import {useDeleteUnit} from "@/hooks/use-attribute";

export const UnitActionDeleteModalState = "unit_action_delete_modal_state"

export function UnitActionDelete() {
  const [open, setOpen] = useState(false);
  const { bool, setBoolState } = useActionState();
  const { selectedUnit, setSelectedUnit } =  useAttributeState();
  const { mutate: deleteUnit, isPending} = useDeleteUnit();
  const queryClient = useQueryClient();

  useEffect(() => {
    if (!selectedUnit) return;

    if (bool[UnitActionDeleteModalState]){
      setOpen(bool[UnitActionDeleteModalState])
    }
  }, [selectedUnit, bool]);

  function onOpenChange(newOpen: boolean) {
    setOpen(newOpen);
    if (!newOpen) {
      setBoolState(UnitActionDeleteModalState, false);
    }
  }

  const onSubmit = (e: FormEvent) => {
    e.preventDefault();
    if (!selectedUnit) {
      toast.warning("Please retry open the confirm delete button");
      onOpenChange(false);
      return;
    }

    deleteUnit(selectedUnit.id, {
      onSuccess: async () => {
        toast.success("Unit delete successfully");
        await queryClient.invalidateQueries({ queryKey: ['units'] })
        setSelectedUnit(null);
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
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
import {useSeatingState} from "@/states/seating-state";
import {toast} from "sonner";
import {useDeleteTable} from "@/hooks/use-seating";
import {useQueryClient} from "@tanstack/react-query";
import {isHTTPResponse} from "@/lib/api";

export const TableActionDeleteModalState = "table_action_delete_modal_state"

export function TableActionDelete() {
  const [open, setOpen] = useState(false);
  const { bool, setBoolState } = useActionState();
  const { selectedFloor, selectedTableId, setSelectedTableId } =  useSeatingState();
  const { mutate: deleteTable, isPending} = useDeleteTable();
  const queryClient = useQueryClient();

  useEffect(() => {
    if (!selectedTableId) return;

    if (bool[TableActionDeleteModalState]){
      setOpen(bool[TableActionDeleteModalState])
    }
  }, [selectedTableId, bool]);

  function onOpenChange(newOpen: boolean) {
    setOpen(newOpen);
    if (!newOpen) {
      setBoolState(TableActionDeleteModalState, false);
    }
  }

  const onSubmit = (e: FormEvent) => {
    e.preventDefault();
    if (!selectedTableId) {
      toast.warning("Please retry open the confirm delete button");
      onOpenChange(false);
      return;
    }

    deleteTable(selectedTableId, {
      onSuccess: async () => {
        toast.success("Table delete successfully");
        await queryClient.invalidateQueries({ queryKey: ['floors'] })
        await queryClient.invalidateQueries({ queryKey: ["floor.tables", selectedFloor?.id] })
        setSelectedTableId(null);
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
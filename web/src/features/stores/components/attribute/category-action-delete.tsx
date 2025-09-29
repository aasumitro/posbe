import {type FormEvent, useEffect, useState} from "react";
import {useActionState} from "@/states/action-state";
import {useAttributeState} from "@/states/attribute-state";
import {useDeleteCategory} from "@/hooks/use-attribute";
import {useQueryClient} from "@tanstack/react-query";
import {toast} from "sonner";
import {isHTTPResponse} from "@/lib/api";
import {
  AlertDialog, AlertDialogAction, AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription, AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle
} from "@/components/ui/alert-dialog";

export const CategoryActionDeleteModalState = "category_action_delete_modal_state"

export function CategoryActionDelete() {
  const [open, setOpen] = useState(false);
  const { bool, setBoolState } = useActionState();
  const { selectedCategory, setSelectedCategory } =  useAttributeState();
  const { mutate: deleteCategory, isPending} = useDeleteCategory();
  const queryClient = useQueryClient();

  useEffect(() => {
    if (!selectedCategory) return;

    if (bool[CategoryActionDeleteModalState]){
      setOpen(bool[CategoryActionDeleteModalState])
    }
  }, [selectedCategory, bool]);

  function onOpenChange(newOpen: boolean) {
    setOpen(newOpen);
    if (!newOpen) {
      setBoolState(CategoryActionDeleteModalState, false);
    }
  }

  const onSubmit = (e: FormEvent) => {
    e.preventDefault();
    if (!selectedCategory) {
      toast.warning("Please retry open the confirm delete button");
      onOpenChange(false);
      return;
    }

    deleteCategory(selectedCategory.id, {
      onSuccess: async () => {
        toast.success("Category delete successfully");
        await queryClient.invalidateQueries({ queryKey: ['categories'] })
        setSelectedCategory(null);
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
              className="cursor-pointer"
              type="button"
              disabled={isPending}
            >Cancel</AlertDialogCancel>
            <AlertDialogAction
              className="bg-red-500 hover:bg-red-600 text-white cursor-pointer"
              type="submit"
              disabled={isPending}
            >Delete</AlertDialogAction>
          </AlertDialogFooter>
        </form>
      </AlertDialogContent>
    </AlertDialog>
  )
}
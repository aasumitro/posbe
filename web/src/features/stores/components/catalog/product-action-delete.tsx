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
import {useProductState} from "@/states/product-state";
import {useDeleteProduct} from "@/hooks/use-product";
import {useQueryClient} from "@tanstack/react-query";
import {toast} from "sonner";
import {isHTTPResponse} from "@/lib/api";

export const ProductActionDeleteModalState = "product_action_delete_modal_state"

export function ProductActionDeleteModal() {
  const [open, setOpen] = useState(false);
  const { bool, setBoolState } = useActionState();
  const {selectedProduct, setSelectedProduct} = useProductState();
  const { mutate: deleteProduct, isPending} = useDeleteProduct();
  const queryClient = useQueryClient();

  useEffect(() => {
    if (!selectedProduct) return;

    if (bool[ProductActionDeleteModalState]){
      setOpen(bool[ProductActionDeleteModalState])
    }
  }, [selectedProduct, bool]);

  const onSubmit = (e: FormEvent) => {
    e.preventDefault();
    if (!selectedProduct) {
      toast.warning("Please retry open the confirm delete button");
      onOpenChange(false);
      return;
    }

    deleteProduct(selectedProduct.id, {
      onSuccess: async () => {
        toast.success("Product delete successfully");
        await queryClient.invalidateQueries({ queryKey: ['products'] })
        setSelectedProduct(null);
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

  function onOpenChange(newOpen: boolean) {
    setOpen(newOpen);
    if (!newOpen) {
      setBoolState(ProductActionDeleteModalState, false);
    }
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
              className="cursor-pointer"
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
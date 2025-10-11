import {useEffect, useState} from "react";
import {useActionState} from "@/states/action-state";
import {
  AlertDialog,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogHeader,
  AlertDialogTitle
} from "@/components/ui/alert-dialog";
import {useOrderState} from "@/states/order-state";

export const AddToCartModalState = "add_to_cart_state"

export function AddToCartModal() {
  const [isOpen, setIsOpen] = useState(false);
  const { bool, setBoolState } = useActionState();
  const { selectedProduct, setSelectedProduct } =  useOrderState();

  useEffect(() => {
    if (!bool[AddToCartModalState]) return;
    if (!selectedProduct) return;
    setIsOpen(bool[AddToCartModalState]);
  }, [bool, selectedProduct]);

  // const onSubmit = () => {}

  const onClose = () => {
    setBoolState(AddToCartModalState, false)
    setIsOpen(false);
    setSelectedProduct(null);
  }

  return (
    <AlertDialog open={isOpen} onOpenChange={onClose}>
      <AlertDialogContent className="min-w-[425px]">
        <AlertDialogHeader>
          <AlertDialogTitle>
            Add to Cart
          </AlertDialogTitle>
          <AlertDialogDescription>
            add {selectedProduct?.name} to cart
          </AlertDialogDescription>
        </AlertDialogHeader>
      </AlertDialogContent>
    </AlertDialog>
  )
}
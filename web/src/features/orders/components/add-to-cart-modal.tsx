import {useEffect, useState} from "react";
import {useActionState} from "@/states/action-state";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle
} from "@/components/ui/dialog";
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
    <Dialog open={isOpen} onOpenChange={onClose}>
      <DialogContent className="min-w-[425px]">
        <DialogHeader>
          <DialogTitle>
            Add to Cart
          </DialogTitle>
          <DialogDescription>
            add {selectedProduct?.name} to cart
          </DialogDescription>
        </DialogHeader>
      </DialogContent>
    </Dialog>
  )
}
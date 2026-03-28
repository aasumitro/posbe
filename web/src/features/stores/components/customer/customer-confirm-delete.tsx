"use client"

import {type FormEvent, useRef} from "react";
import {Loader2Icon} from "lucide-react";
import { Button } from "@/components/ui/button";
import {toast} from "sonner";
import {isHTTPResponse} from "@/lib/api";
import {useQueryClient} from "@tanstack/react-query";
import {Popover, PopoverContent, PopoverTrigger} from "@/components/ui/popover";
import {PopoverClose} from "@radix-ui/react-popover";
import {useDeleteCustomer} from "@/hooks/use-customer";
import {useCustomerState} from "@/states/customer-state";

export function CustomerConfirmDeleteAlertDialog({action}: {action: (state: boolean) => void}) {
  const { selectedCustomer, setSelectedCustomer } =  useCustomerState();
  const { mutate: deleteCustomer, isPending} = useDeleteCustomer();
  const queryClient = useQueryClient();
  const cancelRef = useRef<HTMLButtonElement>(null);

  const onSubmit = (event: FormEvent) => {
    event.preventDefault();

    if (!selectedCustomer) return;

    cancelRef.current?.click();

    deleteCustomer(selectedCustomer?.id, {
      onSuccess: async () => {
        await queryClient.invalidateQueries({ queryKey: ['customers'] })
        toast.success("Customer delete successfully");
        setSelectedCustomer(null);
        action(false);
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

  return (
    <Popover>
      <PopoverTrigger asChild>
        <Button
          variant="link"
          className="text-red-500 cursor-pointer"
        >
          DELETE
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-80">
        <div className="grid gap-4 mb-4">
          <div className="space-y-2">
            <h4 className="leading-none font-medium">Confirm Delete</h4>
            <p className="text-muted-foreground text-sm">
              Are you sure you want to delete this customer ({selectedCustomer?.name})?
            </p>
          </div>
        </div>
        <form onSubmit={onSubmit} className="text-right space-x-2">
          <PopoverClose asChild ref={cancelRef}>
            <Button
              type="button"
              variant="outline"
              className="cursor-pointer"
              disabled={isPending}
            >Cancel</Button>
          </PopoverClose>
          <Button
            type="submit"
            className="bg-red-500 hover:bg-red-600 text-white cursor-pointer"
            disabled={isPending || isPending}
          >
            <Loader2Icon className={
              isPending
                ? "block animate-spin"
                : "hidden"
            }/>
            Confirm
          </Button>
        </form>
      </PopoverContent>
    </Popover>
  );
}
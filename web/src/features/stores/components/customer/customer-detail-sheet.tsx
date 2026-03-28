import {useEffect, useState} from "react";
import {useActionState} from "@/states/action-state";
import {Sheet, SheetContent, SheetDescription, SheetHeader, SheetTitle, SheetFooter} from "@/components/ui/sheet";
import {Input} from "@/components/ui/input";
import {Button} from "@/components/ui/button";
import {useCustomerState} from "@/states/customer-state";
import {useUpdateCustomer} from "@/hooks/use-customer";
import {useQueryClient} from "@tanstack/react-query";
import {useForm} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import {FormSchema, type FormValues, type CustomerErrorResponse} from "./form-schema";
import type {Customer} from "@/types/customer";
import {z} from "zod";
import {toast} from "sonner";
import {isHTTPResponse} from "@/lib/api";
import {Form, FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/components/ui/form";
import {Textarea} from "@/components/ui/textarea";
import {Loader2Icon} from "lucide-react";
import {CustomerConfirmDeleteAlertDialog} from "@/features/stores/components/customer/customer-confirm-delete";

export const CustomerDetailActionSheetState = "customer_detail_action_sheet_state"

export function CustomerDetailSheet() {
  const [open, setOpen] = useState(false);
  const { bool, setBoolState } = useActionState();
  const { setSelectedCustomer, selectedCustomer } = useCustomerState();
  const { mutate: update, isPending } = useUpdateCustomer();
  const queryClient = useQueryClient();

  const form = useForm<FormValues>({
    resolver: zodResolver(FormSchema),
    mode: "onChange",
    defaultValues: {
      name: "",
      phone: "",
      email: "",
      description: ""
    }
  })

  useEffect(() => {
    if (bool[CustomerDetailActionSheetState]){
      setOpen(bool[CustomerDetailActionSheetState])
    }
    if (selectedCustomer) reset(selectedCustomer)
  }, [bool, selectedCustomer]);

  const reset = (selectedCustomer: Customer) => {
    const id = setTimeout(() => {
      form.reset({
        name: selectedCustomer.name,
        email: selectedCustomer.email,
        phone: selectedCustomer.phone,
        description: selectedCustomer.description
      })
    }, 100)
    return () => clearTimeout(id);
  }

  const watchAll = form.watch();

  const isChanged = (
    watchAll.name !== selectedCustomer?.name ||
    watchAll.phone !== selectedCustomer?.phone ||
    watchAll.email !== selectedCustomer?.email ||
    watchAll.description !== selectedCustomer?.description
  );

  function onSubmit(data: z.infer<typeof FormSchema>) {
    if (!selectedCustomer) return;

    let filteredValues = Object.fromEntries(
      Object.entries(data).filter(([_, value]) =>
        value !== undefined && value !== null && value !== ""
      )
    ) as Record<string, unknown>;

    update({
      id: selectedCustomer.id,
      body: JSON.stringify(filteredValues)
    }, {
      onSuccess: async () => {
        await queryClient.invalidateQueries({ queryKey: ['customers'] })
        toast.success("Update customer successfully");
        onOpenChange(false);
      },
      onError: async (error) => {
        if (error && isHTTPResponse<null>(error)) {
          if (typeof error.data === "string") {
            toast.error(error.data);
            return;
          }

          if (typeof error.data === "object" && error.data !== null) {
            const data = error.data as CustomerErrorResponse;

            if (data.name && data.name.length > 0) {
              form.setError("name", {type: "manual", message: data.name[0]})
            }

            if (data.phone && data.phone.length > 0) {
              form.setError("phone", {type: "manual", message: data.phone[0]})
            }

            if (data.email && data.email.length > 0) {
              form.setError("email", {type: "manual", message: data.email[0]})
            }

            if (data.description && data.description.length > 0) {
              form.setError("description", {type: "manual", message: data.description[0]})
            }
          }
        }
        if (error instanceof Error) {
          const clientError = error as Error
          toast.error(clientError.message);
        }
      }
    })
  }

  function onOpenChange(newOpen: boolean) {
    setOpen(newOpen);
    if (!newOpen) {
      setBoolState(CustomerDetailActionSheetState, false);
      setSelectedCustomer(null);
    }
  }

  return (
    <Sheet open={open} onOpenChange={onOpenChange}>
      <Form {...form}>
        <SheetContent>
          <SheetHeader>
            <SheetTitle>Customer Info</SheetTitle>
            <SheetDescription>
              Preview customer details or remove them from the list.
            </SheetDescription>
          </SheetHeader>

          <form id="form-update" onSubmit={form.handleSubmit(onSubmit)}>
            <section className="px-4 grid gap-4 pt-2 pb-8">
              <FormField
                control={form.control}
                name="name"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Name</FormLabel>
                    <FormControl>
                      <Input
                        className="w-full h-10"
                        type="text"
                        placeholder="Zeros Mardigu"
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="phone"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Phone number</FormLabel>
                    <FormControl>
                      <Input
                        className="w-full h-10"
                        type="text"
                        placeholder="+62822XXXXXX"
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="email"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Email</FormLabel>
                    <FormControl>
                      <Input
                        className="w-full h-10"
                        type="email"
                        placeholder="lorem@posbe.id"
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="description"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Description</FormLabel>
                    <FormControl>
                      <Textarea
                        placeholder="Describe the customer (e.g., our regular cs)"
                        className="resize-none"
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </section>
          </form>

          <SheetFooter className="flex flex-row items-center justify-between">
            <CustomerConfirmDeleteAlertDialog action={onOpenChange} />
            <div className="flex justify-end items-center space-x-2">
              <Button
                variant="outline"
                type="button"
                className="cursor-pointer"
                onClick={() => onOpenChange(false)}
              >Cancel</Button>
              <Button
                type="submit"
                className="cursor-pointer"
                form="form-update"
                disabled={!isChanged || isPending}
              >
                {isPending && <Loader2Icon className="w-4 animate-spin" />}
                Save
              </Button>
            </div>
          </SheetFooter>
        </SheetContent>
      </Form>
    </Sheet>
  )
}
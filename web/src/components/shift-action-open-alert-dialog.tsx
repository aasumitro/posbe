"use client"

import {useActionState} from "@/states/action-state";
import {useEffect, useState} from "react";
import {
  AlertDialog,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle
} from "@/components/ui/alert-dialog";
import {Button} from "@/components/ui/button";
import {ContactRound, Loader} from "lucide-react";
import {useForm} from "react-hook-form";
import {z} from "zod";
import {zodResolver} from "@hookform/resolvers/zod";
import {useAuthStore} from "@/states/auth-state";
import {useNavigate} from "@tanstack/react-router";
import {useStoreState} from "@/states/store-state";
import {Form, FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/components/ui/form";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select";
import {Input} from "@/components/ui/input";
import {useQueryClient} from "@tanstack/react-query";
import {useShiftAction} from "@/hooks/use-order";
import {IconExternalLink} from "@tabler/icons-react";

export const OpenShiftModalState = "open_shift_modal_state"

const OpenShiftSchema = z.object({
  shift_id: z.coerce.number().min(1, "Please select shift"),
  cash: z.coerce.number().min(0, "Please set the current cash"),
})

export function ShiftActionOpenAlertDialog() {
  const [isDialogOpen, setDialogOpen] = useState(false);
  const { bool, setBoolState } = useActionState();
  const {auth} = useAuthStore();
  const navigate = useNavigate();
  const {shifts} =  useStoreState();
  const queryClient = useQueryClient();
  const { mutate: action, isPending} = useShiftAction();

  useEffect(() => {
    if (bool[OpenShiftModalState]) setDialogOpen(bool[OpenShiftModalState]);
  }, [bool]);

  const form = useForm<z.infer<typeof OpenShiftSchema>>({
    resolver: zodResolver(OpenShiftSchema),
    defaultValues: {
      shift_id: 0,
      cash: 0,
    },
  });

  const onSubmit = (values: z.infer<typeof OpenShiftSchema>) => {
    const body: Record<string, unknown> = {};
    body.action = "open";
    body.shift_id = values.shift_id;
    body.cash = values.cash;

    action(JSON.stringify(body), {
      onSuccess: (resp) => {
        console.log(resp);
      },
      onError: (error) => {
        console.log(error);
      }
    });
  }

  async function onOpenChange(newOpen: boolean) {
    setDialogOpen(newOpen);
    if (!newOpen) {
      form.reset();
      setBoolState(OpenShiftModalState, false);
      await queryClient.refetchQueries({ queryKey: ['active.shift'] })
      if (auth.user?.role?.name === "waiter") {
        await navigate({to: "/", replace: true})
      }
    }
  }

  return (
    <AlertDialog open={isDialogOpen} onOpenChange={onOpenChange} >
      <AlertDialogContent className="w-96" autoFocus>
        <AlertDialogHeader>
          <AlertDialogTitle>Open Shift</AlertDialogTitle>
          <AlertDialogDescription>
            You need to start a new shift before creating or processing orders.
          </AlertDialogDescription>
        </AlertDialogHeader>

        {auth.user?.role && !["admin", "cashier"].includes(auth.user?.role?.name) && (
          <section className="text-center py-6 border rounded-lg">
            <ContactRound className="my-4 w-12 h-12 mx-auto"/>
            <h4 className="text-primary text-lg font-bold tracking-tight">
              Shift Action Restricted
            </h4>
            <p className="text-secondary-foreground text-sm font-normal px-4">
              You’ll need your admin or cashier to open a new shift before you can take orders.
              <span className="font-light mt-4 block">Press <code>ESC</code> to close this window.</span>
            </p>
          </section>
        )}

        {auth.user?.role && ["admin", "cashier"].includes(auth.user?.role?.name) && (
          <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)}>
              <section className="space-y-4">
                <FormField
                  control={form.control}
                  name="shift_id"
                  render={({ field }) => (
                    <FormItem className="w-full">
                      <FormLabel>Shift</FormLabel>
                      <FormControl>
                        <Select
                          value={field.value ? String(field.value) : ""}
                          onValueChange={(val) => field.onChange(Number(val))}
                        >
                          <SelectTrigger className="w-full min-h-12">
                            <SelectValue placeholder="Select Shift" />
                          </SelectTrigger>
                          <SelectContent position="popper">
                            {shifts?.map((shift) => (
                              <SelectItem key={shift.id} value={String(shift.id)}>
                                {shift.name.charAt(0).toUpperCase() + shift.name.slice(1)}
                              </SelectItem>
                            ))}
                          </SelectContent>
                        </Select>
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />

                <FormField
                  control={form.control}
                  name="cash"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Cash</FormLabel>
                      <FormControl>
                        <Input
                          type="number"
                          className="w-full h-12"
                          placeholder="Enter the cyrrent cash"
                          {...field}
                        />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </section>

              <AlertDialogFooter className="mt-6">
                <Button
                  onClick={async () => await navigate({to: "/", replace: true})}
                  type="button"
                  variant="link"
                  className="mr-auto cursor-pointer"
                >
                  <IconExternalLink className="rotate-y-180" />
                  Home
                </Button>

                <>
                  <AlertDialogCancel
                    disabled={isPending}
                    className="cursor-pointer"
                  >
                    Cancel
                  </AlertDialogCancel>
                  <Button
                    type="submit"
                    className="cursor-pointer"
                    disabled={isPending}
                  >
                    <Loader className={
                      isPending
                        ? "block animate-spin"
                        : "hidden"
                    }/>
                    Confirm
                  </Button>
                </>
              </AlertDialogFooter>
            </form>
          </Form>
        )}
      </AlertDialogContent>
    </AlertDialog>
  )
}
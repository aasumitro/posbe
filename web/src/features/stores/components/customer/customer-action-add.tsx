import {useState} from "react";
import {z} from "zod";
import {useForm} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import {
  AlertDialog, AlertDialogCancel,
  AlertDialogContent, AlertDialogDescription, AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger
} from "@/components/ui/alert-dialog";
import {cn} from "@/lib/utils";
import {IconPlus} from "@tabler/icons-react";
import {Form, FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/components/ui/form";
import {Input} from "@/components/ui/input";
import {Button} from "@/components/ui/button";
import {useQueryClient} from "@tanstack/react-query";
import {useNewCustomer} from "@/hooks/use-customer";
import {Textarea} from "@/components/ui/textarea";
import {isHTTPResponse} from "@/lib/api";
import {toast} from "sonner";
import {Loader2Icon} from "lucide-react";
import {type FormValues, FormSchema, type CustomerErrorResponse} from "./form-schema";

export function CustomerActionAdd() {
  const [open, setOpen] = useState(false)
  const { mutate: addCustomer, isPending } = useNewCustomer();
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

  function onSubmit(data: z.infer<typeof FormSchema>) {
    const payload = {
      name: data.name,
      ...(data.phone.trim() && { phone: data.phone }),
      ...(data.email.trim() && { email: data.email }),
      ...(data.description && { description: data.description }),
    };

    addCustomer(JSON.stringify(payload), {
      onSuccess: async () => {
        await queryClient.invalidateQueries({ queryKey: ['customers'] })
        toast.success("New customer added successfully");
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

  function onOpenChange(state: boolean) {
    setOpen(state);
    if (!state) {
      form.reset();
    }
  }

  return (
    <AlertDialog open={open} onOpenChange={onOpenChange}>
      <AlertDialogTrigger asChild>
        <div className={cn(
          "aspect-video rounded-xl border-2 border-dashed hover:bg-muted/50",
          "flex flex-col items-center justify-center group cursor-pointer min-h-56 max-h-56 w-96"
        )}>
          <IconPlus className="w-12 h-12 text-gray-400 group-hover:text-gray-500" />
        </div>
      </AlertDialogTrigger>
      <AlertDialogContent className="select-none">
        <AlertDialogHeader>
          <AlertDialogTitle>Add Customer</AlertDialogTitle>
          <AlertDialogDescription>
            Fill in the details below to add a new customer. Adding an email or phone number is required to identify them.
          </AlertDialogDescription>
        </AlertDialogHeader>

        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)}>
            <section className="grid gap-4 pt-2 pb-8">
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

            <AlertDialogFooter>
              <AlertDialogCancel>Cancel</AlertDialogCancel>
              <Button
                type="submit"
                className="cursor-pointer"
                disabled={!form.formState.isDirty || !form.formState.isValid || isPending}
              >
                {isPending && <Loader2Icon className="w-4 animate-spin" />}
                Save
              </Button>
            </AlertDialogFooter>
          </form>
        </Form>
      </AlertDialogContent>
    </AlertDialog>
  )
}
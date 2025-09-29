import {
  AlertDialog,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "@/components/ui/alert-dialog"
import {useEffect, useState} from "react";
import {useActionState} from "@/states/action-state";
import {useUpdateProductAddon} from "@/hooks/use-product";
import {z} from "zod";
import {useQueryClient} from "@tanstack/react-query";
import {useProductState} from "@/states/product-state";
import {useForm} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import {toast} from "sonner";
import {isHTTPResponse} from "@/lib/api";
import {Form, FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/components/ui/form";
import {Input} from "@/components/ui/input";
import {Textarea} from "@/components/ui/textarea";
import {Button} from "@/components/ui/button";
import {Loader2Icon} from "lucide-react";

export const AddonActionEditModalState = "addon_action_edit_modal_state"

const FormSchema = z.object({
  price: z.coerce.number({
    required_error: "Please set a price.",
    invalid_type_error: "Price must be a number.",
  }).nonnegative("Price must be greater than or equal to 0"),
  name: z.string({
    required_error: "Please set a name for the addon.",
  }).min(1, "Name must be at least 1 character.")
    .max(50, "Name must be at most 50 characters."),
  description: z.string({
    required_error: "Please provide a description.",
  }).min(1, "Description must be at least 1 character.")
    .max(200, "Description must be at most 200 characters."),
})

type AddonErrorResponse = {
  price?: string[]
  name?: string[]
  description?: string[]
}

export function AddonActionEdit() {
  const [open, setOpen] = useState(false);
  const { selectedAddon, setSelectedAddon } =  useProductState();
  const { bool, setBoolState } = useActionState();
  const {mutate: editAddon, isPending} = useUpdateProductAddon();
  const queryClient = useQueryClient();

  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
    defaultValues: {
      price: 0,
      name: "",
      description: ""
    }
  })

  useEffect(() => {
    if (!selectedAddon) return;

    if (bool[AddonActionEditModalState]){
      setOpen(bool[AddonActionEditModalState])
    }

    if (selectedAddon) {
      form.reset({
        price: selectedAddon.price,
        name: selectedAddon.name,
        description: selectedAddon.description,
      })
    }
  }, [selectedAddon, bool]);

  function onSubmit(data: z.infer<typeof FormSchema>) {
    if (!selectedAddon) return;

    const body: Record<string, unknown> = {};

    if (data.name != selectedAddon.name) body.name = data.name;
    if (data.price != selectedAddon.price) body.price = data.price;
    if (data.description != selectedAddon.description) body.description = data.description;

    editAddon({
      id: selectedAddon.id,
      body: JSON.stringify(body),
    }, {
      onSuccess: async  () => {
        await queryClient.invalidateQueries({ queryKey: ['addons'] })
        toast.success("Update addon successfully");
        onOpenChange(false);
        setSelectedAddon(null);
      },
      onError: (error) => {
        if (error && isHTTPResponse<null>(error)) {
          if (typeof error.data === "string") {
            toast.error(error.data);
            return;
          }

          if (typeof error.data === "object" && error.data !== null) {
            const data = error.data as AddonErrorResponse;

            if (data.name && data.name.length > 0) {
              form.setError("name", {type: "manual", message: data.name[0]})
            }

            if (data.price && data.price.length > 0) {
              form.setError("price", {type: "manual", message: data.price[0]})
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
      setBoolState(AddonActionEditModalState, false);
    }
  }

  return (
    <AlertDialog open={open} onOpenChange={onOpenChange}>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Edit Addon</AlertDialogTitle>
          <AlertDialogDescription>
            Update the details of this addon. You can change its name, description,
            or price as needed.
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
                        placeholder="oat milk"
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="price"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Price</FormLabel>
                    <FormControl>
                      <Input
                        className="w-full h-10"
                        placeholder="oat milk"
                        type="number"
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
                        placeholder="Describe the addon (e.g., oat milk instead of regular milk)"
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
              <AlertDialogCancel
                disabled={isPending}
                className="cursor-pointer"
              >Cancel</AlertDialogCancel>
              <Button
                type="submit"
                disabled={isPending}
                className="cursor-pointer"
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
import {useState} from "react";
import {
  AlertDialog,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger
} from "@/components/ui/alert-dialog";
import {IconPlus} from "@tabler/icons-react";
import {Button} from "@/components/ui/button";
import {Form, FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/components/ui/form";
import {Input} from "@/components/ui/input";
import {z} from "zod";
import {useForm} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import {Loader2Icon} from "lucide-react";
import {useNewFloor} from "@/hooks/use-seating";
import {useQueryClient} from "@tanstack/react-query";
import {toast} from "sonner";
import {isHTTPResponse} from "@/lib/api";

const FormSchema = z.object({
  name: z.string().min(3),
})

type FloorErrorResponse = {
  name?: string[]
}

export function FloorActionAdd() {
  const [open, setOpen] = useState(false)
  const {mutate: addFloor, isPending} = useNewFloor();
  const queryClient = useQueryClient();

  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
    defaultValues: {
      name: "",
    }
  })

  function onOpenChange(state: boolean) {
    setOpen(state);
    if (!open) {
      form.reset();
    }
  }

  function onSubmit(data: z.infer<typeof FormSchema>) {
    addFloor(JSON.stringify({
      name: data.name,
    }), {
      onSuccess: async () => {
        await queryClient.invalidateQueries({ queryKey: ['floors'] })
        toast.success("New floor added successfully");
        onOpenChange(false);
      },
      onError: async (error) => {
        if (error && isHTTPResponse<null>(error)) {
          if (typeof error.data === "string") {
            toast.error(error.data);
            return;
          }

          if (typeof error.data === "object" && error.data !== null) {
            const data = error.data as FloorErrorResponse;

            if (data.name && data.name.length > 0) {
              form.setError("name", {type: "manual", message: data.name[0]})
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

  return (
    <AlertDialog open={open} onOpenChange={onOpenChange}>
      <AlertDialogTrigger asChild>
        <Button className="w-full cursor-pointer">
          <IconPlus />
          Add new Floor
        </Button>
      </AlertDialogTrigger>
      <AlertDialogContent className="select-none">
        <AlertDialogHeader>
          <AlertDialogTitle>Add Floor</AlertDialogTitle>
          <AlertDialogDescription>
            Fill in the details below to add a new floor for your store.
          </AlertDialogDescription>
        </AlertDialogHeader>

        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)}>
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
                      autoComplete="off"
                      placeholder="1st floor"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <AlertDialogFooter className="mt-6">
              <AlertDialogCancel className="cursor-pointer">Cancel</AlertDialogCancel>
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
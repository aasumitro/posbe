import {cn} from "@/lib/utils";
import {IconPlus} from "@tabler/icons-react";
import {
  AlertDialog,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog"
import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import { z } from "zod"
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form"
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select"
import {Input} from "@/components/ui/input";
import {Button} from "@/components/ui/button";
import {useState} from "react";
import {useQueryClient} from "@tanstack/react-query";
import {useNewUnit} from "@/hooks/use-attribute";
import {toast} from "sonner";
import {isHTTPResponse} from "@/lib/api";
import {Loader2Icon} from "lucide-react";

const FormSchema = z.object({
  magnitude: z.string({
    required_error: "Please select a magnitude to display.",
  }),
  symbol: z.string({
    required_error: "Please set a symbol to display.",
  }).min(1, "Symbol must be at least 1 character.").max(3, "Symbol must be at most 3 characters."),
  name: z.string({
    required_error: "Please set a name to display.",
  }).min(1, "Name must be at least 1 character.").max(20, "Name must be at most 20 characters."),
})

type UnitErrorResponse = {
  magnitude?: string[]
  name?: string[]
  symbol?: string[]
}

export function UnitActionAdd() {
  const [open, setOpen] = useState(false)
  const {mutate: addUnit, isPending} = useNewUnit();
  const queryClient = useQueryClient();

  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
    defaultValues: {
      magnitude: "",
      name: "",
      symbol: ""
    }
  })

  function onSubmit(data: z.infer<typeof FormSchema>) {
    addUnit(JSON.stringify({
      name: data.name,
      symbol: data.symbol,
      magnitude: data.magnitude,
    }), {
      onSuccess: async () => {
        await queryClient.invalidateQueries({ queryKey: ['units'] })
        toast.success("New unit added successfully");
        onOpenChange(false);
      },
      onError: async (error) => {
        if (error && isHTTPResponse<null>(error)) {
          if (typeof error.data === "string") {
            toast.error(error.data);
            return;
          }

          if (typeof error.data === "object" && error.data !== null) {
            const data = error.data as UnitErrorResponse;

            if (data.name && data.name.length > 0) {
              form.setError("name", {type: "manual", message: data.name[0]})
            }

            if (data.magnitude && data.magnitude.length > 0) {
              form.setError("magnitude", {type: "manual", message: data.magnitude[0]})
            }

            if (data.symbol && data.symbol.length > 0) {
              form.setError("symbol", {type: "manual", message: data.symbol[0]})
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
    if (!open) {
      form.reset();
    }
  }

  return (
    <AlertDialog open={open} onOpenChange={onOpenChange}>
      <AlertDialogTrigger asChild>
        <div className={cn(
          "aspect-video rounded-xl border-2 border-dashed hover:bg-muted/50",
          "flex flex-col items-center justify-center group cursor-pointer"
        )}>
          <IconPlus className="w-12 h-12 text-gray-400 group-hover:text-gray-500" />
        </div>
      </AlertDialogTrigger>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Add Unit</AlertDialogTitle>
          <AlertDialogDescription>
            Create a new unit to represent a quantity, such as mass, volume, or count, for your products.
          </AlertDialogDescription>
        </AlertDialogHeader>

        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)}>
            <section className="grid gap-4 pt-2 pb-8">
              <FormField
                control={form.control}
                name="magnitude"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Magnitude</FormLabel>
                    <Select
                      onValueChange={field.onChange}
                      defaultValue={field.value}
                    >
                      <FormControl>
                        <SelectTrigger className="w-full min-h-10">
                          <SelectValue placeholder="Select a unit magnitude" />
                        </SelectTrigger>
                      </FormControl>
                      <SelectContent>
                        <SelectItem value="mass">Mass – for solid or powdered products</SelectItem>
                        <SelectItem value="volume">Volume – for liquids or fluids</SelectItem>
                        <SelectItem value="count">Count – for discrete items you can count</SelectItem>
                      </SelectContent>
                    </Select>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="symbol"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Symbol</FormLabel>
                    <FormControl>
                      <Input
                        className="w-full h-10"
                        maxLength={3}
                        placeholder="g"
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="name"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Name</FormLabel>
                    <FormControl>
                      <Input
                        className="w-full h-10"
                        placeholder="gram"
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </section>

            <AlertDialogFooter>
              <AlertDialogCancel disabled={isPending}>Cancel</AlertDialogCancel>
              <Button type="submit" disabled={isPending}>
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
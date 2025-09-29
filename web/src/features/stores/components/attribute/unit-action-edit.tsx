import {useEffect, useState} from "react";
import {useActionState} from "@/states/action-state";
import {useAttributeState} from "@/states/attribute-state";
import {useQueryClient} from "@tanstack/react-query";
import {useUpdateUnit} from "@/hooks/use-attribute";
import {
  AlertDialog, AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription, AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle
} from "@/components/ui/alert-dialog";
import {z} from "zod";
import {useForm} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import {Form, FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/components/ui/form";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select";
import {Input} from "@/components/ui/input";
import {Button} from "@/components/ui/button";
import {Loader2Icon} from "lucide-react";
import {isHTTPResponse} from "@/lib/api";
import {toast} from "sonner";

export const UnitActionEditModalState = "unit_action_edit_modal_state"

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

export function UnitActionEdit() {
  const [open, setOpen] = useState(false);
  const { bool, setBoolState } = useActionState();
  const { selectedUnit, setSelectedUnit } =  useAttributeState();
  const { mutate: updateUnit, isPending} = useUpdateUnit();
  const queryClient = useQueryClient();

  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
    defaultValues: {
      magnitude: "",
      name: "",
      symbol: ""
    }
  })

  useEffect(() => {
    if (!selectedUnit) return;

    if (bool[UnitActionEditModalState]){
      setOpen(bool[UnitActionEditModalState])
    }

    if (selectedUnit) {
      form.reset({
        name: selectedUnit.name,
        symbol: selectedUnit.symbol,
        magnitude: selectedUnit.magnitude,
      })
    }
  }, [selectedUnit, bool]);

  function onSubmit(data: z.infer<typeof FormSchema>) {
    if (!selectedUnit) return;

    const body: Record<string, unknown> = {};

    if (data.name != selectedUnit.name) body.name = data.name;
    if (data.magnitude != selectedUnit.magnitude) body.magnitude = data.magnitude;
    if (data.symbol != selectedUnit.symbol) body.symbol = data.symbol;

    updateUnit({
      id: selectedUnit.id,
      body: JSON.stringify(body),
    }, {
      onSuccess: async  () => {
        await queryClient.invalidateQueries({ queryKey: ['units'] })
        toast.success("Update unit successfully");
        onOpenChange(false);
        setSelectedUnit(null);
      },
      onError: (error) => {
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

  function onOpenChange(newOpen: boolean) {
    setOpen(newOpen);
    if (!newOpen) {
      setBoolState(UnitActionEditModalState, false);
    }
  }

  return (
    <AlertDialog open={open} onOpenChange={onOpenChange}>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Update Unit</AlertDialogTitle>
          <AlertDialogDescription>
            Edit the unit details, such as magnitude, symbol, and name, to represent product quantities like mass, volume, or count.
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
              <AlertDialogCancel className="cursor-pointer" disabled={isPending}>Cancel</AlertDialogCancel>
              <Button type="submit" className="cursor-pointer" disabled={isPending}>
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
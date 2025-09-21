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
import {useSeatingState} from "@/states/seating-state";
import type {Table} from "@/types/seating";
import {z} from "zod";
import {useForm} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import {Form, FormControl, FormField, FormItem, FormMessage} from "@/components/ui/form";
import {useQueryClient} from "@tanstack/react-query";
import {RadioGroup, RadioGroupItem} from "@/components/ui/radio-group";
import {Label} from "@/components/ui/label";
import circleTable from "@/assets/table/circle-min.png";
import rectangleTable from "@/assets/table/rectangle-min.png";
import {Input} from "@/components/ui/input";
import {Slider} from "@/components/ui/slider";
import {cn} from "@/lib/utils";
import type {CircleTableConfig, RectangleTableConfig} from "@/components/table";
import {Table as CTable} from "@/components/table";
import {Button} from "@/components/ui/button";
import {useUpdateTable} from "@/hooks/use-seating";
import {Loader2Icon} from "lucide-react";
import {isHTTPResponse} from "@/lib/api";
import {toast} from "sonner";

export const TableActionUpdateModalState = "table_action_update_modal_state"

const formSchema = z.object({
  type: z.string(),
  name: z
    .string()
    .min(1, { message: "Label must be at least 1 character" })
    .max(3, { message: "Label must be at most 3 characters" })
    .regex(/^[a-zA-Z0-9]+$/, { message: "Label must contain only letters or numbers" }),
  chair: z.number().min(1).max(12)
})

type tableErrorResponse = {
  name?: string[]
  capacity?: string[]
  type?: string[]
}

export function TableActionUpdate() {
  const [open, setOpen] = useState(false);
  const {tables, selectedTableId, setSelectedTableId} =  useSeatingState();
  const {bool, setBoolState} = useActionState();
  const [table, setTable] = useState<Table| null>(null);
  const { mutate: updateTable, isPending} = useUpdateTable();
  const queryClient = useQueryClient();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: "",
      type: "",
      chair: 1,
    },
    // mode: "onChange"
  })

  useEffect(() => {
    if (!selectedTableId) return;

    if (bool[TableActionUpdateModalState]){
      setOpen(bool[TableActionUpdateModalState])
    }

    if (selectedTableId) {
      const t = tables?.find((t) =>
        t.id === selectedTableId) ?? null
      setTable(t);

      form.reset({
        type: t?.type,
        name: t?.name,
        chair: t?.capacity,
      });
    }
  }, [selectedTableId, bool]);

  function onOpenChange(newOpen: boolean) {
    setOpen(newOpen);
    if (!newOpen) {
      setBoolState(TableActionUpdateModalState, false);
      setTable(null);
      form.reset();
    }
  }

  function TablePreview() {
    const type = form.watch("type")
    const name = form.watch("name")
    const chair = form.watch("chair")

    if (type === "rectangle") {
      const defaultHeight = 50
      const defaultWidth =
        chair <= 4 ? 50 :
          chair <= 6 ? 100 :
            chair <= 8 ? 150 :
              chair <= 10 ? 200 : 250

      return (
        <CTable
          id={0}
          name={name}
          status="available"
          config={{
            shape: "rectangle",
            width: defaultWidth,
            height: defaultHeight,
            chairs: chair,
          } as RectangleTableConfig}
        />
      )
    }

    return (
      <CTable
        id={0}
        name={name}
        status="available"
        config={{
          shape: "circle",
          diameter: 50,
          chairs: chair,
        } as CircleTableConfig}
      />
    )
  }

  function onSubmit(values: z.infer<typeof formSchema>) {
    if (!table) return;

    const body: Record<string, unknown> = {};

    if (table.name !== values.name) {
      body.name = values.name;
    }

    if (table.type !== values.type) {
      body.type = values.type;

      if (values.type === "rectangle") {
        body.h_size = 50;
        body.w_size =
          values.chair <= 4 ? 50 :
            values.chair <= 6 ? 100 :
              values.chair <= 8 ? 150 :
                values.chair <= 10 ? 200 : 250;
        body.d_size = 0
      } else {
        body.d_size = 50;
        body.h_size = 0;
        body.w_size = 0;
      }
    }

    if (table.capacity !== values.chair) {
      body.capacity = values.chair;

      if (table.type === values.type) {
        if (values.type === "rectangle") {
          body.h_size = 50;
          body.w_size =
            values.chair <= 4 ? 50 :
              values.chair <= 6 ? 100 :
                values.chair <= 8 ? 150 :
                  values.chair <= 10 ? 200 : 250;
          body.d_size = 0
        } else if (values.type === "circle") {
          body.d_size = 50;
          body.h_size = 0;
          body.w_size = 0;
        }
      }
    }

    if (Object.keys(body).length === 0) {
      toast.info("No changes detected");
      return;
    }

    updateTable({
      id: table.id, body: JSON.stringify(body),
    }, {
      onSuccess: async () => {
        setSelectedTableId(null);
        await queryClient.invalidateQueries({ queryKey: ["floors"] });
        await queryClient.invalidateQueries({ queryKey: ["floor.tables", table?.floor_id] });
        onOpenChange(false);
      },
      onError: (error) => {
        if (error && isHTTPResponse<null>(error)) {
          if (typeof error.data === "string") {
            toast.error(error.data);
            return;
          }

          if (typeof error.data === "object" && error.data !== null) {
            const data = error.data as tableErrorResponse;

            if (data.type && data.type.length > 0) {
              form.setError("type", {type: "manual", message: data.type[0]})
            }

            if (data.name && data.name.length > 0) {
              form.setError("name", {type: "manual", message: data.name[0]})
            }

            if (data.capacity && data.capacity.length > 0) {
              form.setError("chair", {type: "manual", message: data.capacity[0]})
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
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Update table {table?.name}</AlertDialogTitle>
          <AlertDialogDescription>
            Modify the details of this table. Changes will take effect immediately.
          </AlertDialogDescription>
        </AlertDialogHeader>

        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)}>
            <section className="space-y-4">
              <FormField
                control={form.control}
                name="type"
                render={({ field }) => (
                  <FormItem className="space-y-3">
                    <FormControl>
                      <RadioGroup
                        className="grid gap-3 md:grid-cols-2"
                        onValueChange={(value) => {
                          field.onChange(value)
                          form.setValue("chair", 1)
                        }}
                        defaultValue={field.value}
                      >
                        {["circle", "rectangle"].map((table) => (
                          <Label
                            className="has-[[data-state=checked]]:border-ring has-[[data-state=checked]]:bg-input/20 flex items-start gap-3 rounded-lg border p-8"
                            key={table}
                          >
                            <RadioGroupItem
                              value={table}
                              id={table}
                              className="data-[state=checked]:border-primary hidden"
                            />
                            <div className="grid gap-1 font-normal">
                              <img
                                src={table === "circle" ? circleTable : rectangleTable}
                                className="w-full h-full select-none"
                                alt={table}
                                draggable={false}
                              />
                            </div>
                          </Label>
                        ))}
                      </RadioGroup>

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
                    <div className="grid grid-cols-3 items-center gap-4">
                      <Label htmlFor="label">Label/Name</Label>
                      <Input
                        id="label"
                        type="text"
                        maxLength={3}
                        className="col-span-2 h-8"
                        placeholder="e.g: TA1"
                        {...field}
                      />
                    </div>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="chair"
                render={({ field }) => (
                  <FormItem>
                    <div className="grid grid-cols-3 items-center gap-4">
                      <Label htmlFor="chair">Total chair ({form.watch("chair")})</Label>
                      <Slider
                        value={[field.value]}
                        onValueChange={(val) => field.onChange(val[0])}
                        max={form.watch("type") === "circle" ? 6 : 12}
                        step={1}
                        min={1}
                        className="col-span-2 h-8"
                      />
                    </div>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <div className="w-full h-56 border-2 border-dashed rounded-md py-2 px-4 relative inset-0 overflow-hidden">
                <p className="text-muted-foreground text-xs">
                  Table {form.watch("name").toUpperCase()} Preview
                </p>

                <div className="flex justify-center items-center w-full h-full">
                  <div
                    className={cn(
                      "flex items-center justify-center",
                      form.watch("type") === "rectangle" && form.watch("chair") <= 2 && "rotate-90"
                    )}
                    style={{
                      transform:
                        form.watch("type") === "circle"
                          ? "scale(0.6)"
                          : form.watch("chair") >= 8 ? "scale(0.8)" : "scale(0.5)",
                      transformOrigin: "center",
                    }}
                  >
                    <TablePreview />
                  </div>
                </div>
              </div>
            </section>

            <AlertDialogFooter className="mt-6">
              <AlertDialogCancel
                type="button"
                className="cursor-pointer"
                disabled={isPending}
              >Cancel</AlertDialogCancel>
              <Button
                type="submit"
                className="cursor-pointer"
                disabled={isPending}
              >
                {isPending && <Loader2Icon className="w-4 animate-spin" />}
                Save Changes
              </Button>
            </AlertDialogFooter>
          </form>
        </Form>
      </AlertDialogContent>
    </AlertDialog>
  )
}

import {useState} from "react";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover"
import circleTable from "@/assets/table/circle-min.png"
import rectangleTable from "@/assets/table/rectangle-min.png"
import {Table} from "@/components/table";
import type {CircleTableConfig, RectangleTableConfig} from "@/components/table";
import { z } from "zod"
import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormMessage,
} from "@/components/ui/form"
import {RadioGroup, RadioGroupItem} from "@/components/ui/radio-group";
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import {Slider} from "@/components/ui/slider";
import {cn} from "@/lib/utils";
import {IconPlus} from "@tabler/icons-react";
import {useSeatingState} from "@/states/seating-state";
import {useNewTable} from "@/hooks/use-seating";
import {useQueryClient} from "@tanstack/react-query";
import {toast} from "sonner";
import {isHTTPResponse} from "@/lib/api";
import {Loader2Icon} from "lucide-react";

const formSchema = z.object({
  type: z.string(),
  label: z
    .string()
    .min(1, { message: "Label must be at least 1 character" })
    .max(3, { message: "Label must be at most 3 characters" })
    .regex(/^[a-zA-Z0-9]+$/, { message: "Label must contain only letters or numbers" }),
  chair: z.number().min(1).max(12)
})

type tableErrorResponse = {
  floor_id?: string[]
  name?: string[]
  capacity?: string[]
  type?: string[]
}

export function TableActionNew() {
  const [open, setOpen] = useState<boolean>(false)
  const {selectedFloor} =  useSeatingState();
  const {mutate: addTable, isPending} = useNewTable();
  const queryClient = useQueryClient();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      type: "",
      label: "",
      chair: 1,
    },
    mode: "onChange"
  })

  function onSubmit(values: z.infer<typeof formSchema>) {
    if (!selectedFloor) return;

    let w_size = 0;
    let h_size = 0;
    let d_size = 0;

    if (values.type === "rectangle") {
      h_size = 50;
      w_size =
        values.chair <= 4 ? 50 :
          values.chair <= 6 ? 100 :
            values.chair <= 8 ? 150 :
              values.chair <= 10 ? 200 : 250;
    } else if (values.type === "circle") {
      d_size = 50;
    }

    addTable(JSON.stringify({
      floor_id: selectedFloor.id,
      name: values.label,
      capacity: values.chair,
      type: values.type,
      x_pos: 400,
      y_pos: 400,
      w_size,
      h_size,
      d_size,
    }), {
      onSuccess: async () => {
        await queryClient.invalidateQueries({ queryKey: ['floors'] })
        await queryClient.invalidateQueries({ queryKey: ["floor.tables", selectedFloor.id] })
        toast.success("New table added successfully");
        onOpenChange(false);
      },
      onError: async (error) => {
        if (error && isHTTPResponse<null>(error)) {
          if (typeof error.data === "string") {
            toast.error(error.data);
            return;
          }

          if (typeof error.data === "object" && error.data !== null) {
            const data = error.data as tableErrorResponse;

            if (data.name && data.name.length > 0) {
              form.setError("label", {type: "manual", message: data.name[0]})
            }

            if (data.capacity && data.capacity.length > 0) {
              form.setError("chair", {type: "manual", message: data.capacity[0]})
            }

            if (data.type && data.type.length > 0) {
              form.setError("type", {type: "manual", message: data.type[0]})
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

  const onOpenChange = (newOpen: boolean) => {
    setOpen(newOpen);
    if (!newOpen) {
      form.reset();
    }
  }

  function TablePreview() {
    const type = form.watch("type")
    const label = form.watch("label")
    const chair = form.watch("chair")

    if (type === "rectangle") {
      const defaultHeight = 50
      const defaultWidth =
        chair <= 4 ? 50 :
          chair <= 6 ? 100 :
            chair <= 8 ? 150 :
              chair <= 10 ? 200 : 250

      return (
        <Table
          id={0}
          name={label}
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
      <Table
        id={0}
        name={label}
        status="available"
        config={{
          shape: "circle",
          diameter: 50,
          chairs: chair,
        } as CircleTableConfig}
      />
    )
  }

  return (
    <Popover open={open} onOpenChange={onOpenChange}>
      <PopoverTrigger asChild>
        <Button className="absolute left-8 top-8 cursor-pointer">
          <IconPlus />
          Add new Table
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-[400px]">
        <div className="grid gap-4">
          <div className="space-y-2">
            <h4 className="leading-none font-medium">Add New Table</h4>
            <p className="text-muted-foreground text-sm">
              Configure and place a new table on the floor layout.
            </p>
          </div>

          <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
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

              {form.watch("type") && (
                <>
                  <FormField
                    control={form.control}
                    name="label"
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
                      Table {form.watch("label").toUpperCase()} Preview
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

                  <Button
                    type="submit"
                    className="w-full"
                    disabled={form.watch("label") === "" || form.watch("chair") <= 0 || isPending}
                  >
                    {isPending && <Loader2Icon className="w-4 animate-spin" />}
                    Submit
                  </Button>
                </>
              )}
            </form>
          </Form>
        </div>
      </PopoverContent>
    </Popover>
  )
}

import {useState} from "react";
import {
  AlertDialog, AlertDialogCancel,
  AlertDialogContent, AlertDialogDescription, AlertDialogFooter,
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
import {useNewStoreShift} from "@/hooks/use-store";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select";
import {strTimeToInt, timeReference} from "@/lib/time";
import {toast} from "sonner";
import {isHTTPResponse} from "@/lib/api";
import {useQueryClient} from "@tanstack/react-query";

const FormSchema = z.object({
  name: z.string(),
  start_time: z.string(),
  end_time: z.string()
})

type ShiftErrorResponse = {
  name?: string[]
  start_time?: string[]
  end_time?: string[]
}

export function ShiftActionAdd() {
  const [open, setOpen] = useState(false)
  const {mutate: addShift, isPending} = useNewStoreShift();
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
    if (strTimeToInt(data.end_time) <  strTimeToInt(data.start_time)) {
      form.setError("end_time", {type: "manual", message: "End time should not less than start time."})
      return;
    }
    addShift(JSON.stringify({
      name: data.name,
      start_time: strTimeToInt(data.start_time),
      end_time: strTimeToInt(data.end_time),
    }), {
      onSuccess: async () => {
        await queryClient.invalidateQueries({ queryKey: ['store.shifts'] })
        toast.success("New shift added successfully");
        onOpenChange(false);
      },
      onError: async (error) => {
        if (error && isHTTPResponse<null>(error)) {
          if (typeof error.data === "string") {
            toast.error(error.data);
            return;
          }

          if (typeof error.data === "object" && error.data !== null) {
            const data = error.data as ShiftErrorResponse;

            if (data.name && data.name.length > 0) {
              form.setError("name", {type: "manual", message: data.name[0]})
            }

            if (data.start_time && data.start_time.length > 0) {
              form.setError("start_time", {type: "manual", message: data.start_time[0]})
            }

            if (data.end_time && data.end_time.length > 0) {
              form.setError("end_time", {type: "manual", message: data.end_time[0]})
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
        <div className="flex items-center justify-center gap-1">
          <IconPlus className="w-4 h-4" />
          Add new Shift
        </div>
      </AlertDialogTrigger>
      <AlertDialogContent className="select-none">
        <AlertDialogHeader>
          <AlertDialogTitle>Add Store Shift</AlertDialogTitle>
          <AlertDialogDescription>
            Fill in the details below to add a new shift for your store.
          </AlertDialogDescription>
        </AlertDialogHeader>

        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)}>
            <section className="grid gap-4 pt-2">
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
                        placeholder="shift 3"
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </section>

            <section className="flex flex-col md:flex-row gap-2 mt-4 mb-8">
              <FormField
                control={form.control}
                name="start_time"
                render={({ field }) => (
                  <FormItem className="w-full">
                    <FormLabel>Start Time</FormLabel>
                    <FormControl>
                      <Select
                        value={field.value}
                        onValueChange={field.onChange}
                      >
                        <SelectTrigger className="w-full">
                          <SelectValue placeholder="Select Start Time" />
                        </SelectTrigger>
                        <SelectContent position="popper">
                          {timeReference.map((time, index) => (
                            <SelectItem key={index} value={time}>{time}</SelectItem>
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
                name="end_time"
                render={({ field }) => (
                  <FormItem className="w-full">
                    <FormLabel>End Time</FormLabel>

                    <FormControl>
                      <Select
                        value={field.value}
                        onValueChange={field.onChange}
                      >
                        <SelectTrigger className="w-full">
                          <SelectValue placeholder="Select End Time" />
                        </SelectTrigger>
                        <SelectContent position="popper">
                          {timeReference.map((time, index) => (
                            <SelectItem key={index} value={time}>{time}</SelectItem>
                          ))}
                        </SelectContent>
                      </Select>
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
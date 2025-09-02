import {type FormEvent, useEffect, useState} from "react";
import {useActionState} from "@/states/action-state";
import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetFooter,
  SheetHeader,
  SheetTitle
} from "@/components/ui/sheet";
import {useStoreState} from "@/states/store-state";
import {cn} from "@/lib/utils";
import {intToTimeValue, intToTime} from "@/lib/time";
import {Avatar, AvatarFallback, AvatarImage} from "@/components/ui/avatar";
import {Button} from "@/components/ui/button";
import {Popover, PopoverContent, PopoverTrigger} from "@/components/ui/popover";
import {PopoverClose} from "@radix-ui/react-popover";
import {useDeleteShift, useEditStoreShift, useStoreShiftDetail} from "@/hooks/use-store";
import {toast} from "sonner";
import {isHTTPResponse} from "@/lib/api";
import {useQueryClient} from "@tanstack/react-query";
import {Loader2Icon} from "lucide-react";
import {z} from "zod";
import {useForm} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import {Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage} from "@/components/ui/form";
import {Input} from "@/components/ui/input";
import {TimePicker} from "@/components/ui/time-picker";
import type {StoreShift} from "@/types/shift";
import {Skeleton} from "@/components/ui/skeleton";

export const ShiftDetailActionSheetState = "shift_detail_action_sheet_state"

const FormSchema = z.object({
  name: z.string().min(3),
  start_time: z.any(),
  end_time: z.any()
})

type EditShiftErrorResponse = {
  name?: string[]
  start_time?: string[]
  end_time?: string[]
}

export function ShiftDetailActionSheet() {
  const [open, setOpen] = useState(false);
  const { bool, setBoolState } = useActionState();
  const { settings, selectedShift, setSelectedShifts } =  useStoreState();
  const [isSubmitted, setSubmitted] = useState(false);
  const { mutate: deleteShift, isPending: isPendingDelete} = useDeleteShift();
  const { mutate: editShift, isPending: isPendingUpdate} = useEditStoreShift();
  const {data: shift, isPending: isPendingShift} = useStoreShiftDetail(selectedShift?.id ?? undefined)
  const queryClient = useQueryClient();

  useEffect(() => {
    if (bool[ShiftDetailActionSheetState]){
      setOpen(bool[ShiftDetailActionSheetState])
    }
  }, [bool]);

  function onOpenChange(newOpen: boolean) {
    setOpen(newOpen);
    if (!newOpen) {
      setSubmitted(false);
      setSelectedShifts(null);
      setBoolState(ShiftDetailActionSheetState, false);
    }
  }

  const onSubmitDelete = (event: FormEvent) => {
    event.preventDefault();

    if (!selectedShift) return;

    setSubmitted(true);

    deleteShift(selectedShift?.id, {
      onSuccess: async () => {
        await queryClient.invalidateQueries({queryKey: ['store.shifts']})
        toast.success("Shift delete successfully");
        onOpenChange(false);
      },
      onError: async (error) => {
        setSubmitted(false);
        if (error && isHTTPResponse<null>(error)) {
          toast.error(error.data);
          return;
        }
        if (error instanceof Error) {
          const clientError = error as Error
          toast.error(clientError.message);
        }
      }
    });
  }

  const getResetValues = (shift: typeof selectedShift | null) => ({
    name: shift?.name,
    start_time: intToTimeValue(shift?.start_time),
    end_time: intToTimeValue(shift?.end_time),
  });

  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
    defaultValues: getResetValues(selectedShift)
  })

  // Reset form when organization changes
  // noinspection DuplicatedCode
  useEffect(() => {
    if (!selectedShift) return;
    const id = setTimeout(() => {
      form.reset(getResetValues(selectedShift));
    }, 100);
    return () => clearTimeout(id);
    // eslint-disable-next-line
  }, [selectedShift]);

  function formatTime(val: any): number | null {
    if (!val) return null;
    const hh = String(val.hour).padStart(2, "0");
    const mm = String(val.minute).padStart(2, "0");
    return parseInt(hh + mm, 10);
  }

  function onSubmitUpdate(data: z.infer<typeof FormSchema>) {
    if (!selectedShift) return;

    const start = formatTime(data.start_time);
    const end = formatTime(data.end_time);
    if (end && start && (end < start)) {
      form.setError("end_time", {type: "manual",
        message: "End time should not less than start time."})
      return;
    }

    setSubmitted(true);

    editShift({
      id: selectedShift.id,
      body: JSON.stringify({
        name: data.name,
        start_time: start,
        end_time: end,
      })
    }, {
      onSuccess: async () => {
        toast.success("Shift update successfully");
        await queryClient.invalidateQueries({ queryKey: ['store.shifts'] })
        setSubmitted(false);
      },
      onError: async (error) => {
        setSubmitted(false);
        if (error && isHTTPResponse<null>(error)) {
          if (typeof error.data === "string") {
            toast.error(error.data);
            return;
          }

          if (typeof error.data === "object" && error.data !== null) {
            const editErr = error.data as EditShiftErrorResponse;

            if (editErr.name && editErr.name.length > 0) {
              form.setError("name", {type: "manual", message: editErr.name[0]})
            }

            if (editErr.start_time && editErr.start_time.length > 0) {
              form.setError("start_time", {type: "manual", message: editErr.start_time[0]})
            }

            if (editErr.end_time && editErr.end_time.length > 0) {
              form.setError("end_time", {type: "manual", message: editErr.end_time[0]})
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

  function disabledDeleteButton(shift: StoreShift | null): boolean {
    if (!shift) return true;
    return shift.id === 1 || shift.id === 2;
  }

  function Placeholder({avatars}: { avatars: boolean }) {
    if (avatars) return (
      <div className="flex -space-x-2 px-4">
        <Skeleton className="h-8 w-8 rounded-full" />
        <Skeleton className="h-8 w-8 rounded-full" />
        <Skeleton className="h-8 w-8 rounded-full" />
      </div>
    );

    return <Skeleton className="bg-black/25 h-6 w-[100px] mt-2" />
  }

  return (
    <Sheet open={open} onOpenChange={onOpenChange}>
      <SheetContent className="flex flex-col">
        <SheetHeader>
          <SheetTitle>{selectedShift?.name?.toUpperCase()} Detail</SheetTitle>
          <SheetDescription>
            Check the details of this shift and track its activity.
          </SheetDescription>
        </SheetHeader>

        <div className="overflow-auto space-y-4">
          <h1 className="font-bold text-md px-4">Performance Summary</h1>

          <section className="flex flex-col gap-2 px-4">
            <section className="flex flex-row gap-2 justify-between">
              <div className={cn(
                "aspect-video rounded-xl bg-muted/50",
                "flex flex-col items-center justify-center w-full p-4"
              )}>
                <p className="uppercase font-mono text-xs tracking-widest">open</p>
                {isPendingShift ? <Placeholder avatars={false} /> : (
                  <h5 className="font-bold mt-2 tracking-wider">
                    {shift?.data ? intToTime(shift.data.start_time) : "-"}
                  </h5>
                )}
              </div>
              <div className={cn(
                "aspect-video rounded-xl bg-muted/50",
                "flex flex-col items-center justify-center w-full p-4"
              )}>
                <p className="uppercase font-mono text-xs tracking-widest">close</p>
                {isPendingShift ? <Placeholder avatars={false}/> : (
                  <h5 className="font-bold mt-2 tracking-wider">
                    {shift?.data ? intToTime(shift.data.end_time) : "-"}
                  </h5>
                )}
              </div>
            </section>

            <section className="flex flex-row gap-2 justify-between">
              <div className={cn(
                "aspect-video rounded-xl bg-muted/50",
                "flex flex-col items-center justify-center w-full p-4"
              )}>
                <p className="uppercase font-mono text-xs tracking-widest">USE</p>
                {isPendingShift ? <Placeholder avatars={false}/> : (
                  <h5 className="font-bold mt-2 text-xl tracking-wider">
                    {shift?.data && shift.data.total_usage > 0 ? `${shift.data.total_usage}x` : "-"}
                  </h5>
                )}
              </div>
              <div className={cn(
                "aspect-video rounded-xl bg-muted/50",
                "flex flex-col items-center justify-center w-full p-4"
              )}>
                <p className="uppercase font-mono text-xs tracking-widest">TRX</p>
                {isPendingShift ? <Placeholder avatars={false}/> : (
                  <h5 className="font-bold mt-2 text-xl tracking-wider">
                    {shift?.data && shift.data.total_transaction > 0 ? `${shift.data.total_transaction}x` : "-"}
                  </h5>
                )}
              </div>
            </section>

            <div className={cn(
              "aspect-video rounded-xl bg-muted/50",
              "flex flex-col items-center justify-center w-full h-28 p-4"
            )}>
              <p className="uppercase font-mono text-xs tracking-widest">
                Last Activity
              </p>
              <h5 className="font-bold mt-2 text-xl tracking-wider ">
                Today (active)
              </h5>
            </div>
          </section>

          <section className="flex flex-col gap-2 px-4">
            <section className="flex flex-row gap-2 justify-between">
              <div className={cn(
                "aspect-video rounded-xl bg-muted/50",
                "flex flex-col items-center justify-center w-full p-4"
              )}>
                <p className="uppercase font-mono text-xs tracking-widest">Surplus</p>
                {isPendingShift ? <Placeholder avatars={false}/> : (
                  <h5 className="font-bold mt-2 text-xl tracking-wider">
                    {shift?.data && shift.data.total_surplus > 0 ? `${shift.data.total_surplus}x` : "-"}
                  </h5>
                )}
              </div>
              <div className={cn(
                "aspect-video rounded-xl bg-muted/50",
                "flex flex-col items-center justify-center w-full p-4"
              )}>
                <p className="uppercase font-mono text-xs tracking-widest">Deficit</p>
                {isPendingShift ? <Placeholder avatars={false}/> : (
                  <h5 className="font-bold mt-2 text-xl tracking-wider">
                    {shift?.data && shift.data.total_deficit > 0 ? `${shift.data.total_deficit}x` : "-"}
                  </h5>
                )}
              </div>
            </section>

            <section className="flex flex-row gap-2 justify-between">
              <div className={cn(
                "aspect-video rounded-xl bg-muted/50",
                "flex flex-col items-center justify-center w-full p-4"
              )}>
                <p className="uppercase font-mono text-xs tracking-widest">Profit</p>
                {isPendingShift ? <Placeholder avatars={false}/> : (
                  <h5 className={cn(
                    "font-bold mt-2 text-xl tracking-wider",
                    shift?.data && shift.data.profit && "text-green-500"
                  )}>
                    {shift?.data && shift.data.profit > 0 ? `${settings?.currency} 5M` : "-"}
                  </h5>
                )}
              </div>
              <div className={cn(
                "aspect-video rounded-xl bg-muted/50",
                "flex flex-col items-center justify-center w-full p-4"
              )}>
                <p className="uppercase font-mono text-xs tracking-widest">Loss</p>
                {isPendingShift ? <Placeholder avatars={false}/> : (
                  <h5 className={cn(
                    "font-bold mt-2 text-xl tracking-wider",
                    shift?.data && shift.data.loss && "text-red-500"
                  )}>
                    {shift?.data && shift.data.loss > 0 ? `${settings?.currency} 1M` : "-"}
                  </h5>
                )}
              </div>
            </section>

            <div className={cn(
              "aspect-video rounded-xl bg-muted/50",
              "flex flex-col items-center justify-center w-full h-28 p-4"
            )}>
              <p className="uppercase font-mono text-xs tracking-widest">
                Net Profit
              </p>
              {isPendingShift ? <Placeholder avatars={false}/> : (
                <h5 className={cn(
                  "font-bold mt-2 text-xl tracking-wider",
                  shift?.data && shift.data.net ? (shift.data.net > 0 ? "text-green-500" : "text-red-500") : ""
                )}>
                  {shift?.data && shift.data.net > 0 ? `${settings?.currency} 1M` : "-"}
                </h5>
              )}
            </div>
          </section>

          <h1 className="font-bold text-md px-4">Users</h1>
          {isPendingShift ? <Placeholder avatars={true}/> : (
            <div className="*:data-[slot=avatar]:ring-background flex -space-x-2 *:data-[slot=avatar]:ring-2 *:data-[slot=avatar]:grayscale px-4">
              <Avatar>
                <AvatarImage src="https://github.com/shadcn.png" alt="@shadcn" />
                <AvatarFallback>CN</AvatarFallback>
              </Avatar>
              <Avatar>
                <AvatarImage src="https://github.com/leerob.png" alt="@leerob" />
                <AvatarFallback>LR</AvatarFallback>
              </Avatar>
              <Avatar>
                <AvatarImage
                  src="https://github.com/evilrabbit.png"
                  alt="@evilrabbit"
                />
                <AvatarFallback>ER</AvatarFallback>
              </Avatar>
            </div>
          )}
        </div>

        <SheetFooter className="flex flex-row justify-between">
          <Popover>
            <PopoverTrigger asChild>
              <Button
                variant="link"
                className="text-red-500 cursor-pointer"
                disabled={disabledDeleteButton(selectedShift)}
              >
                DELETE
              </Button>
            </PopoverTrigger>
            <PopoverContent className="w-80">
              <div className="grid gap-4 mb-4">
                <div className="space-y-2">
                  <h4 className="leading-none font-medium">Confirm Delete</h4>
                  <p className="text-muted-foreground text-sm">
                    Are you sure you want to delete this shift?
                  </p>
                </div>
              </div>
              <form onSubmit={onSubmitDelete} className="text-right space-x-2">
                <PopoverClose asChild>
                  <Button
                    type="button"
                    variant="outline"
                    className="cursor-pointer"
                    disabled={isPendingDelete}
                  >Cancel</Button>
                </PopoverClose>
                <Button
                  className="bg-red-500 hover:bg-red-600 text-white"
                  type="submit"
                  disabled={isSubmitted || isPendingDelete}
                >
                  <Loader2Icon className={
                    isSubmitted
                      ? "block animate-spin"
                      : "hidden"
                  }/>
                  Confirm
                </Button>
              </form>
            </PopoverContent>
          </Popover>

          <Popover>
            <PopoverTrigger asChild>
              <Button className="cursor-pointer">
                Update
              </Button>
            </PopoverTrigger>

            <PopoverContent className="w-[400px]">
              <div className="grid gap-4 mb-4">
                <div className="space-y-2">
                  <h4 className="leading-none font-medium">Update Shift</h4>
                  <p className="text-muted-foreground text-sm">
                    Make changes to this shift and save your updates.
                  </p>
                </div>
              </div>

              <Form {...form}>
                <form onSubmit={form.handleSubmit(onSubmitUpdate)} className="text-right space-x-2">
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

                  <div className="mt-4 space-y-6">
                    <div className="flex items-center space-x-2 text-sm">
                      <div className="flex-grow h-px bg-gray-200" />
                      <span className="whitespace-nowrap">Work Time</span>
                      <div className="flex-grow h-px bg-gray-200" />
                    </div>
                  </div>

                  <section className="flex flex-col md:flex-row gap-2 mt-4 mb-8">
                    <FormField
                      control={form.control}
                      name="start_time"
                      render={({ field }) => (
                        <FormItem className="w-full">
                          <FormControl>
                            <TimePicker
                              aria-label="Start time"
                              value={field.value}
                              onChange={field.onChange}
                            />
                          </FormControl>
                          {!form.formState.errors.start_time && (
                            <FormDescription>This is the shift start time</FormDescription>
                          )}
                          <FormMessage />
                        </FormItem>
                      )}
                    />

                    <FormField
                      control={form.control}
                      name="end_time"
                      render={({ field }) => (
                        <FormItem className="w-full">
                          <FormControl>
                            <TimePicker
                              aria-label="end time"
                              value={field.value}
                              onChange={field.onChange}
                            />
                          </FormControl>
                          {!form.formState.errors.end_time && (
                            <FormDescription>This is the shift end time</FormDescription>
                          )}
                          <FormMessage />
                        </FormItem>
                      )}
                    />
                  </section>

                  <PopoverClose asChild>
                    <Button
                      type="button"
                      variant="outline"
                      className="cursor-pointer"
                      disabled={isPendingUpdate}
                    >Cancel</Button>
                  </PopoverClose>
                  <Button
                    type="submit"
                    disabled={isSubmitted || isPendingUpdate}
                  >
                    <Loader2Icon className={
                      isSubmitted
                        ? "block animate-spin"
                        : "hidden"
                    }/>
                    Save
                  </Button>
                </form>
              </Form>
            </PopoverContent>
          </Popover>
        </SheetFooter>
      </SheetContent>
    </Sheet>
  )
}

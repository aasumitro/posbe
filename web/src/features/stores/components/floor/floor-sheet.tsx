import {
  Sheet,
  SheetContent,
  SheetDescription, SheetFooter,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "@/components/ui/sheet"
import {IconEdit, IconFocusAuto, IconTrash} from "@tabler/icons-react";
import {Button} from "@/components/ui/button";
import {FloorActionAdd} from "@/features/stores/components/floor/new-floor-action";
import {useSeatingState} from "@/states/seating-state";
import {
  Choicebox,
  ChoiceboxItem,
  ChoiceboxItemHeader,
  ChoiceboxItemSubtitle,
  ChoiceboxItemTitle
} from "@/components/ui/choicebox";
import {Popover, PopoverContent, PopoverTrigger} from "@/components/ui/popover";
import {PopoverClose} from "@radix-ui/react-popover";
import {Form, FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/components/ui/form";
import {Input} from "@/components/ui/input";
import {z} from "zod";
import {useForm} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import type {Floor} from "@/types/seating";
import {type FormEvent, useRef} from "react";
import {toast} from "sonner";
import {isHTTPResponse} from "@/lib/api";
import {useDeleteFloor, useUpdateFloor} from "@/hooks/use-seating";
import {Loader2Icon} from "lucide-react";
import {useQueryClient} from "@tanstack/react-query";

export function FloorSheet() {
  const {floors} =  useSeatingState();

  return (
    <Sheet>
      <SheetTrigger asChild>
        <Button variant="outline" className="w-full cursor-pointer">
          <IconFocusAuto />
          Manage Floors
        </Button>
      </SheetTrigger>
      <SheetContent>
        <SheetHeader>
          <SheetTitle>Manage Floors</SheetTitle>
          <SheetDescription>
            Here you can add new floors, modify details, or delete existing floors.
          </SheetDescription>
        </SheetHeader>

        <section className="px-4">
          <Choicebox defaultValue="1">
            {floors?.map((floor) => (
              <ChoiceboxItem
                key={floor.id}
                value={(floor?.id+1).toString() as string}
              >
                <ChoiceboxItemHeader>
                  <ChoiceboxItemTitle>
                    {floor.name}
                    <ChoiceboxItemSubtitle>
                      ({floor?.total_tables
                      ? `tables: ${floor.total_tables}`
                      : 'no table'})
                    </ChoiceboxItemSubtitle>
                    <div className="ml-auto space-x-2">
                      <UpdateFloor floor={floor} />
                      <DeleteFloor floor={floor} />
                    </div>
                  </ChoiceboxItemTitle>
                </ChoiceboxItemHeader>
              </ChoiceboxItem>
            ))}
          </Choicebox>
        </section>

        <SheetFooter className="border-t">
          <FloorActionAdd />
        </SheetFooter>
      </SheetContent>
    </Sheet>
  )
}

const FormSchema = z.object({
  name: z.string().min(3),
})

type FloorErrorResponse = {
  name?: string[]
}

function UpdateFloor({floor}: {floor: Floor}) {
  const { mutate: updateFloor, isPending: isPendingUpdate} = useUpdateFloor(floor.id);
  const queryClient = useQueryClient();
  const cancelRef = useRef<HTMLButtonElement>(null);

  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
    defaultValues: {
      name: floor.name,
    }
  })

  function onSubmit(data: z.infer<typeof FormSchema>) {
    if (!floor) return;

    cancelRef.current?.click();

    updateFloor(JSON.stringify({
      name: data.name,
    }), {
      onSuccess: async () => {
        await queryClient.invalidateQueries({ queryKey: ['floors'] })
        toast.success("Floor update successfully");
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
    <Popover>
      <PopoverTrigger asChild>
        <Button
          variant="outline"
          className="cursor-pointer"
        >
          <IconEdit />
        </Button>
      </PopoverTrigger>

      <PopoverContent className="w-[400px]">
        <div className="grid gap-4 mb-4">
          <div className="space-y-2">
            <h4 className="leading-none font-medium">Update Floor</h4>
            <p className="text-muted-foreground text-sm">
              Make changes to this floor and save your updates.
            </p>
          </div>
        </div>

        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="text-right space-x-2 space-y-6">
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

            <PopoverClose asChild>
              <Button
                type="button"
                variant="outline"
                className="cursor-pointer"
                disabled={isPendingUpdate}
                ref={cancelRef}
              >Cancel</Button>
            </PopoverClose>
            <Button
              type="submit"
              disabled={isPendingUpdate}
            >
              <Loader2Icon className={
                isPendingUpdate
                  ? "block animate-spin"
                  : "hidden"
              }/>
              Save
            </Button>
          </form>
        </Form>
      </PopoverContent>
    </Popover>
  )
}

function DeleteFloor({floor}: {floor: Floor}) {
  const { mutate: deleteFloor, isPending: isPendingDelete} = useDeleteFloor();
  const queryClient = useQueryClient();
  const cancelRef = useRef<HTMLButtonElement>(null);

  const onSubmitDelete = (event: FormEvent) => {
    event.preventDefault();

    if (!floor) return;

    cancelRef.current?.click();

    deleteFloor(floor.id, {
      onSuccess: async () => {
        await queryClient.invalidateQueries({ queryKey: ['floors'] })
        toast.success("Floor delete successfully");
      },
      onError: async (error) => {
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

  return (
    <Popover>
      <PopoverTrigger asChild>
        <Button
          variant="outline"
          className="text-red-500 cursor-pointer"
        >
          <IconTrash />
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-80">
        <div className="grid gap-4 mb-4">
          <div className="space-y-2">
            <h4 className="leading-none font-medium">Confirm Delete</h4>
            <p className="text-muted-foreground text-sm">
              Are you sure you want to delete the floor "{floor.name}"{floor?.total_tables &&
              `? This action will also remove ${floor.total_tables} associated tables.`}
            </p>
          </div>
        </div>
        <form className="text-right space-x-2" onSubmit={onSubmitDelete}>
          <PopoverClose asChild>
            <Button
              type="button"
              variant="outline"
              className="cursor-pointer"
              disabled={isPendingDelete}
              ref={cancelRef}
            >Cancel</Button>
          </PopoverClose>
          <Button
            className="bg-red-500 hover:bg-red-600 text-white"
            type="submit"
            disabled={isPendingDelete}
          >
            <Loader2Icon className={
              isPendingDelete
                ? "block animate-spin"
                : "hidden"
            }/>
            Confirm
          </Button>
        </form>
      </PopoverContent>
    </Popover>
  )
}
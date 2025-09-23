import {
  AlertDialog,
  AlertDialogCancel,
  AlertDialogContent, AlertDialogDescription, AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "@/components/ui/alert-dialog";
import {cn} from "@/lib/utils";
import {IconPlus, IconTrash} from "@tabler/icons-react";
import {Form, FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/components/ui/form";
import {Input} from "@/components/ui/input";
import {Button} from "@/components/ui/button";
import {Loader2Icon} from "lucide-react";
import {useEffect, useState} from "react";
import {useActionState} from "@/states/action-state";
import {useAttributeState} from "@/states/attribute-state";
import {useQueryClient} from "@tanstack/react-query";
import {z} from "zod";
import {useForm} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import type {Subcategory} from "@/types/category";
import {useDeleteSubcategory, useUpdateCategory} from "@/hooks/use-attribute";
import {Popover, PopoverContent, PopoverTrigger} from "@/components/ui/popover";
import {PopoverClose} from "@radix-ui/react-popover";
import {toast} from "sonner";
import {isHTTPResponse} from "@/lib/api";

export const CategoryActionEditModalState = "category_action_edit_modal_state"

const FormSchema = z.object({
  name: z.string({
    required_error: "Please set a name to display.",
  }).min(3, "Name must be at least 1 character.").max(20, "Name must be at most 20 characters."),
})

type CategoryErrorResponse = {
  name?: string[]
}

export function CategoryActionEdit() {
  const [open, setOpen] = useState(false);
  const { bool, setBoolState } = useActionState();
  const { selectedCategory } =  useAttributeState();
  const [subcategories, setSubcategories] = useState<Subcategory[]>([]);
  const {mutate: editCategory, isPending} = useUpdateCategory();
  const {mutate: deleteSubcategory, isPending: isPendingDelete} = useDeleteSubcategory();
  const queryClient = useQueryClient();

  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
    defaultValues: {
      name: ""
    }
  })

  useEffect(() => {
    if (!selectedCategory) return;

    if (bool[CategoryActionEditModalState]){
      setOpen(bool[CategoryActionEditModalState])
    }

    if (selectedCategory) {
      form.reset({
        name: selectedCategory.name,
      })

      if (!selectedCategory.subcategories) return;

      setSubcategories([...selectedCategory.subcategories])
    }
  }, [selectedCategory, bool]);

  const handleSubcategoryChange = (index: number, subcategory: Subcategory) => {
    const updated = [...subcategories];
    updated[index] = subcategory;
    setSubcategories(updated);
  };

  const handleAddSubcategory = () => {
    if (!selectedCategory) return;
    setSubcategories([...subcategories,
      {id: -1, category_id: selectedCategory?.id, name: ""}]);

    // generate a new negative id (temporary client id)
    const minId = subcategories.length > 0
      ? Math.min(...subcategories.map(s => s.id)) : 0

    const newId = minId <= 0 ? minId - 1 : -1;

    setSubcategories([
      ...subcategories,
      { id: newId, category_id: selectedCategory.id, name: "" }
    ]);
  };

  const onRemoveSubcategory = (index: number) => {
    const subcategory = subcategories.find((_, i) => i === index)

    if (!subcategory) return;

    if (subcategory?.id < 0) {
      const newSubcategories = subcategories.filter(
        (_, i) => i !== index)
      setSubcategories(newSubcategories);
      return;
    }

    deleteSubcategory({
      cid: subcategory?.category_id,
      sid: subcategory?.id,
    }, {
      onSuccess: async () => {
        await queryClient.invalidateQueries({ queryKey: ['categories'] })
        toast.success("Subcategory delete successfully");
      },
      onError: (error) => {
        if (error && isHTTPResponse<null>(error)) {
          if (typeof error.data === "string") {
            toast.error(error.data);
            return;
          }
        }

        if (error instanceof Error) {
          const clientError = error as Error
          toast.error(clientError.message);
        }
      },
    })
  };

  function onSubmit(data: z.infer<typeof FormSchema>) {
    if (!selectedCategory) return;

    const body: Record<string, unknown> = {};
    if (data.name != selectedCategory.name) body.name = data.name;
    const editSubcategory = subcategories
      .filter((s) => s.id > 0 && s.name.trim() !== "")
      .filter((s) => {
        const original = selectedCategory?.subcategories?.find((o) => o.id === s.id);
        return original && original.name !== s.name; // keep only if name changed
      });
    if (editSubcategory.length > 0) {
      body.subcategories = editSubcategory.map((s) => ({
        id: s.id,
        name: s.name,
      }));
    }
    const newSubcategory = subcategories
      .filter((s) => s.id < 0 && s.name.trim() !== "")
    if (newSubcategory.length > 0) {
      body.new_subcategories = newSubcategory.map((s) => s.name);
    }

    editCategory({
      id: selectedCategory?.id,
      body: JSON.stringify(body),
    }, {
      onSuccess: async () => {
        await queryClient.invalidateQueries({ queryKey: ['categories'] })
        toast.success("Category update successfully");
      },
      onError: (error) => {
        if (error && isHTTPResponse<null>(error)) {
          if (typeof error.data === "string") {
            toast.error(error.data);
            return;
          }

          if (typeof error.data === "object" && error.data !== null) {
            const data = error.data as CategoryErrorResponse;

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

  function onOpenChange(newOpen: boolean) {
    setOpen(newOpen);
    if (!newOpen) {
      setBoolState(CategoryActionEditModalState, false);
    }
  }

  return (
    <AlertDialog open={open} onOpenChange={onOpenChange}>
      <AlertDialogContent className="select-none">
        <AlertDialogHeader>
          <AlertDialogTitle>Add Category</AlertDialogTitle>
          <AlertDialogDescription>
            Define a new product category and optionally include related subcategories to organize your items.
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
                      placeholder="foods"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <section className="p-4 border-2 border-dashed mt-4 mb-8 rounded-lg space-y-4">
              <p className="text-xs text-muted-foreground">Subcategories</p>
              {subcategories?.map((value, index) => (
                <div key={index} className="flex flex-row gap-2 items-center justify-between">
                  <Input
                    placeholder="meat"
                    value={value.name}
                    onChange={(e) => {
                      e.preventDefault();
                      handleSubcategoryChange(index, {
                        ...value,
                        name: e.target.value,
                      });
                    }}
                  />
                  <div className={cn(
                    "flex gap-1",
                    subcategories.length > 1 && "min-w-20"
                  )}>
                    {subcategories.length > 1 && (
                       <>
                         {value.id > 0 && (
                           <Popover>
                             <PopoverTrigger asChild>
                               <Button
                                 type="button"
                                 variant="outline"
                                 disabled={isPendingDelete}
                               >
                                 <IconTrash size={16} />
                               </Button>
                             </PopoverTrigger>
                             <PopoverContent className="w-80">
                               <div className="grid gap-4 mb-4">
                                 <div className="space-y-2">
                                   <h4 className="leading-none font-medium">Confirm Delete</h4>
                                   <p className="text-muted-foreground text-sm">
                                     Are you sure you want to delete this subcategory?
                                   </p>
                                 </div>
                               </div>
                               <div className="text-right space-x-2">
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
                                   disabled={isPendingDelete}
                                   onClick={(e) => {
                                     e.preventDefault();
                                     onRemoveSubcategory(index)
                                   }}
                                 >
                                   <Loader2Icon className={
                                     isPendingDelete
                                       ? "block animate-spin"
                                       : "hidden"
                                   }/>
                                   Confirm
                                 </Button>
                               </div>
                             </PopoverContent>
                           </Popover>
                         )}

                         {value.id < 0 && (
                           <Button
                             type="button"
                             variant="outline"
                             onClick={(e) => {
                               e.preventDefault();
                               onRemoveSubcategory(index)
                             }}
                           >
                             <IconTrash size={16} />
                           </Button>
                         )}
                       </>
                    )}
                    {(index === subcategories.length - 1 || subcategories.length === 1) && (
                      <Button
                        type="button"
                        variant="outline"
                        onClick={handleAddSubcategory}
                      >
                        <IconPlus size={16} />
                      </Button>
                    )}
                  </div>
                </div>
              ))}
            </section>

            <AlertDialogFooter>
              <AlertDialogCancel disabled={isPending}>
                Cancel
              </AlertDialogCancel>
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

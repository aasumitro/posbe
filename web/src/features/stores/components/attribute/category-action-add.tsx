import {z} from "zod";
import {useState} from "react";
import {useForm} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import {
  AlertDialog, AlertDialogCancel,
  AlertDialogContent, AlertDialogDescription, AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger
} from "@/components/ui/alert-dialog";
import {cn} from "@/lib/utils";
import {IconPlus, IconTrash} from "@tabler/icons-react";
import {Form, FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/components/ui/form";
import {Input} from "@/components/ui/input";
import {Button} from "@/components/ui/button";
import {useNewCategory} from "@/hooks/use-attribute";
import {useQueryClient} from "@tanstack/react-query";
import {Loader2Icon} from "lucide-react";
import {isHTTPResponse} from "@/lib/api";
import {toast} from "sonner";

const FormSchema = z.object({
  name: z.string({
    required_error: "Please set a name to display.",
  }).min(1, "Name must be at least 1 character.").max(20, "Name must be at most 20 characters."),
})

type CategoryErrorResponse = {
  name?: string[]
}

export function CategoryActionAdd() {
  const [open, setOpen] = useState(false)
  const [subcategories, setSubcategories] = useState<string[]>([""]);
  const {mutate: addCategory, isPending} = useNewCategory();
  const queryClient = useQueryClient();

  const handleSubcategoryChange = (index: number, value: string) => {
    const updated = [...subcategories];
    updated[index] = value;
    setSubcategories(updated);
  };

  const handleAddSubcategory = () => {
    setSubcategories([...subcategories, ""]);
  };

  const handleRemoveSubcategory = (index: number) => {
    setSubcategories(subcategories.filter((_, i) => i !== index));
  };

  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
    defaultValues: {
      name: ""
    }
  })

  function onSubmit(data: z.infer<typeof FormSchema>) {
    addCategory(JSON.stringify({
      name:data.name,
      subcategories: subcategories
        .filter((s) => s.trim() !== "")
    }), {
      onSuccess: async () => {
        await queryClient.invalidateQueries({ queryKey: ['categories'] })
        toast.success("New category added successfully");
        onOpenChange(false);
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

  function onOpenChange(state: boolean) {
    setOpen(state);
    if (!open) {
      form.reset();
      setSubcategories([""]);
    }
  }

  return (
    <AlertDialog open={open} onOpenChange={onOpenChange}>
      <AlertDialogTrigger asChild>
        <div className={cn(
          "h-full rounded-xl border-2 border-dashed hover:bg-muted/50",
          "flex flex-col items-center justify-center group cursor-pointer"
        )}>
          <IconPlus className="w-12 h-12 text-gray-400 group-hover:text-gray-500" />
        </div>
      </AlertDialogTrigger>
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
              {subcategories.map((value, index) => (
                <div key={index} className="flex flex-row gap-2 items-center justify-between">
                  <Input
                    placeholder="meat"
                    value={value}
                    onChange={(e) => {
                      e.preventDefault();
                      handleSubcategoryChange(index, e.target.value);
                    }}
                  />
                  <div className={cn(
                    "flex gap-1",
                    subcategories.length > 1 && "min-w-20"
                  )}>
                    {subcategories.length > 1 && (
                      <Button
                        type="button"
                        variant="outline"
                        onClick={(e) => {
                          e.preventDefault();
                          handleRemoveSubcategory(index)
                        }}
                      >
                        <IconTrash size={16} />
                      </Button>
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

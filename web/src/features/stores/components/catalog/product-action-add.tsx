import {Accordion, AccordionContent, AccordionItem, AccordionTrigger} from "@/components/ui/accordion";
import {IconChevronLeft, IconInfoCircle, IconPlus, IconTrash} from "@tabler/icons-react";
import {z} from "zod";
import { useFieldArray } from "react-hook-form";
import type {Unit} from "@/types/unit";
import {Input} from "@/components/ui/input";
import {Separator} from "@/components/ui/separator";
import ReactQuill from 'react-quill-new';
import 'react-quill-new/dist/quill.snow.css';
import {type Control, useForm, useFormState, useWatch} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import {Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage} from "@/components/ui/form";
import {Button} from "@/components/ui/button";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select";
import {useState} from "react";
import {useNavigate} from "@tanstack/react-router";
import type {Category} from "@/types/category";
import {useNewProduct} from "@/hooks/use-product";
import {isHTTPResponse} from "@/lib/api";
import {toast} from "sonner";
import {useQueryClient} from "@tanstack/react-query";
import {Route} from "@/routes/_authenticated/stores/route";

interface ProductActionAddProps {
  categories?: Category[] | null;
  units?: Unit[] | null;
}

interface VariantContainerProps {
  control: Control<z.infer<typeof FormSchema>>,
  units?: Unit[] | null
}

interface VariantGroupProps {
  control: Control<z.infer<typeof FormSchema>>;
  groupIndex: number;
  onRemoveGroup: () => void;
  units?: Unit[] | null;
}

const FormSchema = z.object({
  image: z
    .any()
    .refine(
      (file) => {
        if (!file) return true;
        return file instanceof File;
      },
      { message: "Invalid file." }
    )
    .refine(
      (file) => {
        if (!file) return true;
        return ["image/jpeg", "image/png", "image/jpg"].includes(file.type);
      },
      { message: "Only .jpeg, .jpg, .png files are allowed." }
    )
    .refine(
      async (file) => {
        if (!file) return true;

        const img = await new Promise<HTMLImageElement>(
          (resolve, reject) => {
            const reader = new FileReader();
            reader.onload = (e) => {
              const image = new Image();
              image.onload = () => resolve(image);
              image.onerror = reject;
              image.src = e.target?.result as string;
            };
            reader.readAsDataURL(file);
          });

        return img.width >= 250 && img.height >= 250;
      }, { message: "Image must be at least 250x250px." }).optional(),
  code: z.string({
    required_error: "Please set a code for the product.",
  }).min(1, "Code must be at least 1 character.")
    .max(16, "Name must be at most 16 characters."),
  name: z.string({
    required_error: "Please set a name for the product.",
  }).min(1, "Name must be at least 1 character.")
    .max(50, "Name must be at most 50 characters."),
  category_id: z.coerce.number().min(1, "Please select product category"),
  subcategory_id: z.coerce.number().min(1, "Please select product subcategory"),
  description: z.string(),
  variant_groups: z.array(
    z.object({
      type: z.string().min(1, "Type is required"),
      variants: z.array(
        z.object({
          name: z.string().min(1, "Name is required"),
          description: z.string().min(3, "Description is required").max(200),
          unit_id: z.coerce.number().optional(),
          unit_size: z.coerce.number().optional(),
          price: z.coerce.number().min(0),
        }).refine((variant) => {
          // If unit_id > 0, unit_size must be > 0
          if (variant.unit_id && variant.unit_id > 0) {
            return (variant.unit_size ?? 0) > 0;
          }
          return true;
        }, {
          message: "Unit size is required when a unit is selected",
          path: ["unit_size"], // points the error to the unit_size field
        })
      )
    })
  ).optional(),
})

export function ProductActionAdd({categories, units}: ProductActionAddProps) {
  const navigate = useNavigate();
  const [status, setStatus] = useState<"draft" | "publish">("draft");
  const {mutate: addProduct, isPending} = useNewProduct();
  const queryClient = useQueryClient();

  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
    defaultValues: {
      image: "",
      code: "",
      name: "",
      category_id: 0,
      subcategory_id: 0,
      description: "",
      variant_groups: [
        {
          type: "",
          variants: [
            {
              name: "",
              description: "",
              unit_id: 0,
              unit_size: 0,
              price: 0,
            }
          ]
        }
      ],
    }
  })

  async function onSubmit(data: z.infer<typeof FormSchema>) {
    const body: Record<string, unknown> = {};
    if (data.image) {
      body.image = await new Promise<string>((resolve, reject) => {
        const reader = new FileReader();
        reader.onload = () => resolve(reader.result as string);
        reader.onerror = reject;
        reader.readAsDataURL(data.image);
      });
    }
    body.status = status
    body.sku = data.code;
    body.name = data.name;
    if (data.category_id > 0) {
      body.category_id = data.category_id;
    }
    if (data.subcategory_id > 0) {
      body.subcategory_id = data.subcategory_id;
    }
    if (data.description !== "") {
      body.description = data.description;
    }
    if (Array.isArray(data.variant_groups) && data.variant_groups.length > 0) {
      let variants: {
        type: string, name: string, description: string,
        price: number, unit_id?: number, unit_size?: number
      }[] = [];
      data.variant_groups.forEach((group) => {
        group.variants.forEach((variant) => {
          variants.push({
            type: group.type,
            name: variant.name,
            description: variant.description,
            price: variant.price,
            unit_id: variant.unit_id,
            unit_size: variant.unit_size
          });
        })
      })
      body.variants = variants;
    }

    addProduct(JSON.stringify(body), {
      onSuccess: async () => {
        await queryClient.invalidateQueries({ queryKey: ['products'] })
        toast.success("Product add successfully");
        await navigate({
          from:Route.fullPath,
          to: "/stores/catalogs",
          search: {tab: "products"}
        })
      },
      onError: (error) => {
        if (error && isHTTPResponse<null>(error)) {
          if (typeof error.data === "string") {
            toast.error(error.data);
            return;
          }

          if (typeof error.data === "object" && error.data !== null) {
            // TODO: apply this
            // const data = error.data as ProductErrorResponse;
            //
            // if (data.name && data.name.length > 0) {
            //   form.setError("name", {type: "manual", message: data.name[0]})
            // }
          }
        }

        if (error instanceof Error) {
          const clientError = error as Error
          toast.error(clientError.message);
        }
      },
    });
  }

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)}>
        <aside className="flex items-center justify-between gap-4 px-8">
          <div className="flex gap-3 items-center">
            <Button
              variant="outline"
              className="w-8 h-8 cursor-pointer"
              onClick={async (e) => {
                e.preventDefault();
                await navigate({
                  from:Route.fullPath,
                  to: "/stores/catalogs",
                  search: {tab: "products"}
                })
              }}
            >
              <IconChevronLeft className="w-4 h-4" />
            </Button>
            <h5 className="text-lg font-semibold">
              Add New Product
            </h5>
          </div>
          <div className="flex gap-2 items-center">
            <Button
              variant="ghost"
              className="cursor-pointer"
              onClick={() => setStatus("draft")}
              disabled={isPending}
            >
              Save as Draft
            </Button>
            <Button
              className="cursor-pointer"
              onClick={() => setStatus("publish")}
              disabled={isPending}
            >
              Publish
            </Button>
          </div>
        </aside>

        <aside className="mx-4 grid grid-cols-2 gap-12 p-8">
          <div className="w-full space-y-4 mb-10">
            <h4 className="text-md font-semibold">Product Image</h4>
            <FormField
              control={form.control}
              name="image"
              render={({ field }) => (
                <FormItem>
                  <FormControl>
                    <label
                      htmlFor="image-upload"
                      className="w-52 h-52 border-2 border-dashed rounded-lg cursor-pointer flex flex-col items-center justify-center text-gray-500 hover:border-black overflow-hidden"
                    >
                      {field.value ? (
                        <img
                          src={typeof field.value === "string" ? field.value : URL.createObjectURL(field.value)}
                          alt="Preview"
                          className="object-cover w-full h-full"
                        />
                      ) : (
                        <>
                          <svg className="w-6 h-6 mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 4v16m8-8H4" />
                          </svg>
                          <p className="text-sm text-center">
                            Drop your image here<br />
                            or <span className="text-blue-600 underline">Click to browse</span>
                          </p>
                        </>
                      )}
                      <input
                        type="file"
                        id="image-upload"
                        accept="image/jpeg,image/png,image/jpg"
                        className="hidden"
                        onChange={(e) => {
                          const file = e.target.files?.[0];
                          if (file) {
                            field.onChange(file); // save file into RHF state
                          }
                        }}
                      />
                    </label>
                  </FormControl>
                  <FormMessage />
                  <p className="mt-2 text-xs text-gray-500">
                    ℹ️ Only .jpeg, .jpg, .png. Minimum size 250x250px for optimal use.
                  </p>
                </FormItem>
              )}
            />

            <Separator className="my-6" />

            <h4 className="text-md font-semibold">Product Information</h4>
            <FormField
              control={form.control}
              name="code"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Code</FormLabel>
                  <FormControl>
                    <Input
                      className="w-full h-12"
                      placeholder="Enter the (UPC, SKU, etc) Code"
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
                      className="w-full h-12"
                      placeholder="Enter the Name (Wagyu A5)"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <div className="flex flex-col md:flex-row gap-2">
              <FormField
                control={form.control}
                name="category_id"
                render={({ field }) => (
                  <FormItem className="w-full">
                    <FormLabel>Category</FormLabel>
                    <FormControl>
                      <Select
                        value={field.value ? String(field.value) : undefined}
                        onValueChange={(val) => field.onChange(Number(val))}
                      >
                        <SelectTrigger className="w-full min-h-12">
                          <SelectValue placeholder="Select category" />
                        </SelectTrigger>
                        <SelectContent position="popper">
                          {categories?.map((category) => (
                            <SelectItem key={category.id} value={String(category.id)}>
                              {category.name.charAt(0).toUpperCase() + category.name.slice(1)}
                            </SelectItem>
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
                name="subcategory_id"
                render={({ field }) => (
                  <FormItem className="w-full">
                    <FormLabel>Subcategory</FormLabel>
                    <FormControl>
                      <Select
                        value={field.value ? String(field.value) : undefined}
                        onValueChange={(val) => field.onChange(Number(val))}
                        disabled={!form.watch("category_id")}
                      >
                        <SelectTrigger className="w-full min-h-12">
                          <SelectValue placeholder="Select subcategory" />
                        </SelectTrigger>
                        <SelectContent position="popper">
                          {categories?.find((c) => c.id === form.watch("category_id"))
                            ?.subcategories?.map((subcategory) => (
                              <SelectItem key={subcategory.id} value={String(subcategory.id)}>
                                {subcategory.name.charAt(0).toUpperCase() + subcategory.name.slice(1)}
                              </SelectItem>
                            ))}
                        </SelectContent>
                      </Select>
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </div>

            <FormField
              control={form.control}
              name="description"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Description</FormLabel>
                  <FormControl>
                    <ReactQuill
                      theme="snow"
                      className="h-64"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>

          <div className="w-full space-y-4">
            <VariantContainer control={form.control} units={units} />
          </div>
        </aside>
      </form>
    </Form>
  )
}

function VariantContainer({ control, units }: VariantContainerProps) {
  const { fields: groupFields, append: addGroup, remove: removeGroup } =
    useFieldArray({
      control,
      name: "variant_groups",
    });

  const [openGroup, setOpenGroup] = useState(`group-0`);

  return (
    <div className="w-full space-y-4">
      <div className="flex justify-between items-center">
        <h4 className="text-md font-semibold flex gap-2 items-center">
          Product Variants
          <IconInfoCircle className="text-muted-foreground cursor-pointer" />
        </h4>
        <Button
          variant="link"
          type="button"
          className="cursor-pointer"
          onClick={() => {
            addGroup({ type: "", variants: [] });
            setOpenGroup(`group-${groupFields.length}`);
          }}
        >
          <IconPlus className="w-4 h-4" /> Add variant group
        </Button>
      </div>

      <Accordion
        collapsible
        type="single"
        className="my-4 w-full space-y-2"
        onValueChange={(val) => setOpenGroup(val)}
        value={openGroup}
      >
        {groupFields.map((group, groupIndex) => (
          <VariantGroup
            key={group.id}
            control={control}
            groupIndex={groupIndex}
            onRemoveGroup={() => removeGroup(groupIndex)}
            units={units}
          />
        ))}
      </Accordion>
    </div>
  );
}

function VariantGroup({control, groupIndex, onRemoveGroup, units}: VariantGroupProps) {
  const { errors } = useFormState({ control });

  const { fields: variantFields, append: addVariant, remove: removeVariant } =
    useFieldArray({control, name: `variant_groups.${groupIndex}.variants`});

  const typeValue = useWatch({
    control, name: `variant_groups.${groupIndex}.type`});

  const displayType = typeValue && typeValue.trim() !== ""
    ? typeValue : `New variant group ${groupIndex}`;

  const variantArray = useWatch({
    control, name: `variant_groups.${groupIndex}.variants` });

  return (
    <AccordionItem
      className="border last:border-b rounded-md px-4 w-full"
      value={`group-${groupIndex}`}
    >
      <AccordionTrigger className="cursor-pointer">
        {displayType}
      </AccordionTrigger>

      <AccordionContent>
        <div className="mt-4 border rounded-xl p-4 bg-gray-50">
          <FormField
            control={control}
            name={`variant_groups.${groupIndex}.type`}
            render={({ field }) => (
              <FormItem>
                <FormLabel>Type (Variant Group Name)</FormLabel>
                <FormControl>
                  <Input
                    className="w-full h-12 bg-white"
                    placeholder="Enter group type (e.g. Size, Doneness)"
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
        </div>

        {variantFields.map((variant, variantIndex) => {
          const unitIdValue = variantArray[variantIndex]?.unit_id;
          const nameValue = variantArray[variantIndex]?.name || `Variant ${variantIndex + 1}`;

          return (
            <div key={variant.id} className="mt-4 border rounded-xl">
              <div className="rounded-t-xl flex items-center justify-between bg-gray-50 p-4">
                <h4 className="text-md">{nameValue}</h4>
                <Button
                  variant="ghost"
                  type="button"
                  className="text-red-500 cursor-pointer"
                  onClick={() => removeVariant(variantIndex)}
                >
                  <IconTrash className="w-4 h-4" />
                </Button>
              </div>

              <div className="px-4 py-6 space-y-4">
                <FormField
                  control={control}
                  name={`variant_groups.${groupIndex}.variants.${variantIndex}.name`}
                  render={({ field }) => (
                    <FormItem className="w-full">
                      <FormLabel>Name</FormLabel>
                      <FormControl>
                        <Input
                          className="w-full h-12"
                          placeholder="Enter the variant name (e.g. L, XL)"
                          {...field}
                        />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />

                <FormField
                  control={control}
                  name={`variant_groups.${groupIndex}.variants.${variantIndex}.description`}
                  render={({ field }) => (
                    <FormItem className="w-full">
                      <FormLabel>Description</FormLabel>
                      <FormControl>
                        <Input
                          className="w-full h-12"
                          placeholder="Enter the variant description (e.g: large, extra large)"
                          {...field}
                        />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />

                <FormField
                  control={control}
                  name={`variant_groups.${groupIndex}.variants.${variantIndex}.unit_id`}
                  render={({ field }) => (
                    <FormItem className="w-full">
                      <FormLabel>Unit</FormLabel>
                      <FormControl>
                        <Select
                          value={field.value ? String(field.value) : undefined}
                          onValueChange={(val) => field.onChange(Number(val))}
                        >
                          <SelectTrigger className="w-full min-h-12">
                            <SelectValue placeholder="Select unit" />
                          </SelectTrigger>
                          <SelectContent position="popper">
                            {units?.map((unit) => (
                              <SelectItem key={unit.id} value={String(unit.id)}>
                                {unit.name.charAt(0).toUpperCase() + unit.name.slice(1)} ({unit.symbol})
                              </SelectItem>
                            ))}
                          </SelectContent>
                        </Select>
                      </FormControl>
                      {!errors?.variant_groups?.[groupIndex]?.variants?.[variantIndex]?.unit_id && (
                        <FormDescription className="mt-1 text-sm">
                          Set the variant unit. This is an optional field, but if you select a unit, you must also specify the unit size.
                        </FormDescription>
                      )}
                      <FormMessage />
                    </FormItem>
                  )}
                />

                <FormField
                  control={control}
                  name={`variant_groups.${groupIndex}.variants.${variantIndex}.unit_size`}
                  render={({ field }) => (
                    <FormItem className="w-full">
                      <FormLabel>
                        Unit Size
                      </FormLabel>
                      <FormControl>
                        <Input
                          type="number"
                          className="w-full h-12"
                          disabled={!unitIdValue}
                          {...field}
                        />
                      </FormControl>
                      {!errors?.variant_groups?.[groupIndex]?.variants?.[variantIndex]?.unit_size && (
                        <FormDescription>
                          Set the unit size, e.g., 200g (the "g" corresponds to the selected unit).
                        </FormDescription>
                      )}
                      <FormMessage />
                    </FormItem>
                  )}
                />

                <FormField
                  control={control}
                  name={`variant_groups.${groupIndex}.variants.${variantIndex}.price`}
                  render={({ field }) => (
                    <FormItem className="w-full">
                      <FormLabel>
                        Price
                      </FormLabel>
                      <FormControl>
                        <Input
                          className="w-full h-12"
                          type="number"
                          placeholder="Variant price *optional"
                          {...field}
                        />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </div>
            </div>
          )
        })}

        <div className="flex justify-between my-4">
          <Button
            type="button"
            variant="ghost"
            className="cursor-pointer"
            onClick={() => addVariant({
              name: "",
              unit_size: 0,
              price: 0,
              description: "",
              unit_id: 0
            })}
          >
            <IconPlus className="w-4 h-4" /> Add Variant
          </Button>

          <Button
            type="button"
            variant="ghost"
            className="text-red-500 cursor-pointer"
            onClick={onRemoveGroup}
          >
            <IconTrash className="w-4 h-4" /> Delete Group
          </Button>
        </div>
      </AccordionContent>
    </AccordionItem>
  );
}

import {z} from "zod";
import {useForm} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import {Card, CardContent, CardFooter} from "@/components/ui/card";
import {Form, FormControl, FormField, FormItem, FormMessage} from "@/components/ui/form";
import {Button} from "@/components/ui/button";
import {Input} from "@/components/ui/input";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select";
import {useStoreState} from "@/states/store-state";
import {useEffect, useMemo} from "react";
import {useUpdateSetting} from "@/hooks/use-store";
import {useQueryClient} from "@tanstack/react-query";
import {toast} from "sonner";
import {isHTTPResponse} from "@/lib/api";
import {Loader2Icon} from "lucide-react";

const StoreInfoFormSchema = z.object({
  name: z.string().min(3).max(100),
  phone: z.string()
    .min(6, "Phone number must be at least 6 digits")
    .regex(/^\+[1-9]\d{6,14}$/,
      "Invalid phone number format. Please use international format, e.g: +628123456789"
    ),
  email: z.string().email(),
  type: z.string(),
  address_line1: z.string(),
  address_line2: z.string(),
  country: z.string(),
  state_province: z.string(),
  city_regency: z.string(),
  postal_code: z.string(),
})

type InfoErrorResponse = {
  name?: string[]
  phone?: string[]
  email?: string[]
  type?: string[]
}

export function InfoSection() {
  const {settings} =  useStoreState();
  const {mutate: update, isPending} = useUpdateSetting()
  const queryClient = useQueryClient();

  const getResetValues = (sett: typeof settings | null) => {
    const location = sett?.address?.split("<>").map(s => s.trim()) ?? []
    const [line1 = "", line2 = "", city_regency = "", state_province = "", country = "", postal_code = ""] = location

    return {
      name: sett?.name ?? "",
      phone: sett?.phone ??  "",
      email: sett?.email ?? "",
      type: sett?.type ?? "",
      address_line1: line1,
      address_line2: line2,
      country: country.toLocaleLowerCase(),
      state_province: state_province.toLocaleLowerCase(),
      city_regency: city_regency.toLocaleLowerCase(),
      postal_code: postal_code,
    }
  };

  // Initialize form with live default values
  const form = useForm<z.infer<typeof StoreInfoFormSchema>>({
    resolver: zodResolver(StoreInfoFormSchema),
    defaultValues: getResetValues(settings ?? null),
  });

  // Reset form when organization changes
  // noinspection DuplicatedCode
  useEffect(() => {
    if (!settings) return;
    const id = setTimeout(() => {
      form.reset(getResetValues(settings));
    }, 100);
    return () => clearTimeout(id);
    // eslint-disable-next-line
  }, [settings]);

  // Watch all form values
  const watchValues = form.watch();

  // Compare form with initial data
  const isDirtyComparedToSettingData = useMemo(() => {
    const current = watchValues;
    const initial = getResetValues(settings);
    return Object.keys(initial).some((key) => {
      return current[key as keyof typeof current] !== initial[key as keyof typeof initial];
    });
    // eslint-disable-next-line
  }, [watchValues, settings]);

  function restoreChanges() {
    if (settings) form.reset(getResetValues(settings));
  }

  function buildAddress(data: z.infer<typeof StoreInfoFormSchema>) {
    return [
      data.address_line1,
      data.address_line2,
      data.city_regency,
      data.state_province,
      data.country,
      data.postal_code
    ]
      .filter(v => v !== "")
      .join(" <> ");
  }

  function buildBody(data: z.infer<typeof StoreInfoFormSchema>, sett: typeof settings | null) {
    const body: Record<string, unknown> = {};

    // generic field comparison
    const fields: (keyof typeof data)[] = ["name", "phone", "email", "type"];
    for (const field of fields) {
      // @ts-ignore
      if (data[field] !== "" && data[field] !== sett?.[field]) {
        body[field] = data[field];
      }
    }

    // address building
    const address = buildAddress(data);
    if (address && address !== sett?.address) {
      body.address = address;
    }

    return body;
  }

  function onSubmit(data: z.infer<typeof StoreInfoFormSchema>) {
    const body = buildBody(data, settings);
    if (Object.keys(body).length === 0) {
      return;
    }

    update(JSON.stringify(body), {
      onSuccess: async () => {
        await queryClient.invalidateQueries({ queryKey: ['store.settings'] })
        toast.success("Store info data updated successfully.");
      },
      onError: async (error) => {
        if (error && isHTTPResponse<null>(error)) {
          if (typeof error.data === "string") {
            toast.error(error.data);
            return;
          }

          if (typeof error.data === "object" && error.data !== null) {
            const data = error.data as InfoErrorResponse;

            if (data.name && data.name.length > 0) {
              form.setError("name", {type: "manual", message: data.name[0]})
            }

            if (data.phone && data.phone.length > 0) {
              form.setError("phone", {type: "manual", message: data.phone[0]})
            }

            if (data.email && data.email.length > 0) {
              form.setError("email", {type: "manual", message: data.email[0]})
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

  return (
    <div className="space-y-6">
      <Card>
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)}>
            <CardContent>
              <div className="grid w-full items-center gap-4">
                <FormField
                  control={form.control}
                  name="name"
                  render={({ field }) => (
                    <FormItem className="w-full">
                      <FormControl>
                        <Input
                          type="text"
                          placeholder="Name"
                          {...field}
                        />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />

                <FormField
                  control={form.control}
                  name="phone"
                  render={({ field }) => (
                    <FormItem className="w-full">
                      <FormControl>
                        <Input
                          type="text"
                          placeholder="Phone number"
                          {...field}
                        />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />

                <FormField
                  control={form.control}
                  name="email"
                  render={({ field }) => (
                    <FormItem className="w-full">
                      <FormControl>
                        <Input
                          type="email"
                          placeholder="Email address"
                          {...field}
                        />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />

                <FormField
                  control={form.control}
                  name="type"
                  render={({ field }) => (
                    <FormItem className="w-full">
                      <FormControl>
                        <Select
                          value={field.value}
                          onValueChange={field.onChange}
                        >
                          <SelectTrigger id="business-category" className="w-full">
                            <SelectValue placeholder="Store Type / Category" />
                          </SelectTrigger>
                          <SelectContent position="popper">
                            <SelectItem value="bar">Bar</SelectItem>
                            <SelectItem value="coffee">Coffee</SelectItem>
                            <SelectItem value="restaurant">Restaurant</SelectItem>
                          </SelectContent>
                        </Select>
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />

                <div className="mt-2 space-y-6">
                  <div className="flex items-center space-x-2 text-sm">
                    <div className="flex-grow h-px bg-gray-200" />
                    <span className="whitespace-nowrap">Address</span>
                    <div className="flex-grow h-px bg-gray-200" />
                  </div>
                </div>

                <FormField
                  control={form.control}
                  name="address_line1"
                  render={({ field }) => (
                    <FormItem className="w-full">
                      <FormControl>
                        <Input
                          type="text"
                          placeholder="Address line 1. e.g: Jl. Merdeka No. 123, RT 04 RW 05"
                          {...field}
                        />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />

                <FormField
                  control={form.control}
                  name="address_line2"
                  render={({ field }) => (
                    <FormItem className="w-full">
                      <FormControl>
                        <Input
                          type="text"
                          placeholder="Address line 2 (optional). e.g: Lantai 2, Ruko Blok C, Kompleks Harmoni"
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
                    name="country"
                    render={({ field }) => (
                      <FormItem className="w-full">
                        <FormControl>
                          <Select
                            value={field.value}
                            onValueChange={field.onChange}
                          >
                            <SelectTrigger className="w-full">
                              <SelectValue placeholder="Select country" />
                            </SelectTrigger>
                            <SelectContent position="popper">
                              <SelectItem value="indonesia">Indonesia</SelectItem>
                            </SelectContent>
                          </Select>
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />

                  <FormField
                    control={form.control}
                    name="state_province"
                    render={({ field }) => (
                      <FormItem className="w-full">
                        <FormControl>
                          <Select
                            value={field.value}
                            onValueChange={field.onChange}
                            disabled={!form.watch("country")}
                          >
                            <SelectTrigger className="w-full">
                              <SelectValue placeholder="Select State / Provice" />
                            </SelectTrigger>
                            <SelectContent position="popper">
                              <SelectItem value="sulawesi utara">Sulawesi Utara</SelectItem>
                            </SelectContent>
                          </Select>
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                </div>

                <div className="flex flex-col md:flex-row gap-2">
                  <FormField
                    control={form.control}
                    name="city_regency"
                    render={({ field }) => (
                      <FormItem className="w-full">
                        <FormControl>
                          <Select
                            value={field.value}
                            onValueChange={field.onChange}
                            disabled={!form.watch("state_province")}
                          >
                            <SelectTrigger className="w-full">
                              <SelectValue placeholder="Select City / Regency" />
                            </SelectTrigger>
                            <SelectContent position="popper">
                              <SelectItem value="kota manado">Kota Manado</SelectItem>
                              <SelectItem value="kota bitung">Kota Bitung</SelectItem>
                              <SelectItem value="kota kotamobagu">Kota Kotamobagu</SelectItem>
                              <SelectItem value="kota tomohon">Kota Tomohon</SelectItem>
                              <SelectItem value="bolaang mongondow">Kabupaten Bolaang Mongondow</SelectItem>
                              <SelectItem value="bolaang mongondow selatan">Kabupaten Bolaang Mongondow Selatan</SelectItem>
                              <SelectItem value="bolaang mongondow timur">Kabupaten Bolaang Mongondow Timur</SelectItem>
                              <SelectItem value="bolaang mongondow utara">Kabupaten Bolaang Mongondow Utara</SelectItem>
                              <SelectItem value="kepulauan sangihe">Kabupaten Kepulauan Sangihe</SelectItem>
                              <SelectItem value="kepulauan sitaro">Kabupaten Kepulauan Sitaro (Siau Tagulandang Biaro)</SelectItem>
                              <SelectItem value="kepulauan talaud">Kabupaten Kepulauan Talaud</SelectItem>
                              <SelectItem value="minahasa">Kabupaten Minahasa</SelectItem>
                              <SelectItem value="minahasa selatan">Kabupaten Minahasa Selatan</SelectItem>
                              <SelectItem value="minahasa utara">Kabupaten Minahasa Utara</SelectItem>
                              <SelectItem value="minahasa tenggara">Kabupaten Minahasa Tenggara</SelectItem>
                            </SelectContent>
                          </Select>
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />

                  <FormField
                    control={form.control}
                    name="postal_code"
                    render={({ field }) => (
                      <FormItem className="w-full">
                        <FormControl>
                          <Input
                            type="text"
                            placeholder="Postal code"
                            disabled={!form.watch("city_regency")}
                            {...field}
                          />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                </div>
              </div>
            </CardContent>

            {isDirtyComparedToSettingData && (
              <CardFooter className="mt-6 flex gap-2 justify-end border-t-2 pt-4">
                <Button
                  variant="outline"
                  size="sm"
                  type="button"
                  className="text-xs"
                  onClick={restoreChanges}
                >Cancel</Button>
                <Button
                  size="sm"
                  className="text-xs"
                >
                  {isPending && <Loader2Icon className="w-4 animate-spin mr-1" />}
                  Save
                </Button>
              </CardFooter>
            )}
          </form>
        </Form>
      </Card>
    </div>
  )
}
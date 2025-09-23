import {z} from "zod";
import {useForm} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import {Card, CardContent, CardFooter} from "@/components/ui/card";
import {Form, FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/components/ui/form";
import {Button} from "@/components/ui/button";
import {Input} from "@/components/ui/input";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select";
import {useStoreState} from "@/states/store-state";
import {useEffect, useMemo} from "react";
import {useUpdateSetting} from "@/hooks/use-store";
import {useQueryClient} from "@tanstack/react-query";
import {Loader2Icon} from "lucide-react";
import {toast} from "sonner";
import {isHTTPResponse} from "@/lib/api";

const ServiceTaxFormSchema = z.object({
  currency: z.string(),
  service_category: z.string(),
  service_rate: z.coerce.number(),
  tax_category: z.string(),
  tax_rate: z.coerce.number(),
})

type ServiceAndTaxRateErrorResponse = {
  service_rate?: string[]
  tax_rate?: string[]
}

export function ServiceAndTaxRateSection() {
  const {settings} =  useStoreState();
  const {mutate: update, isPending} = useUpdateSetting()
  const queryClient = useQueryClient();

  const getResetValues = (sett: typeof settings | null) => ({
    currency: sett?.currency.toLocaleLowerCase() ?? "",
    service_category: sett?.service_category ?? "",
    service_rate: sett?.service_rate !== undefined && sett?.service_rate !== null && sett?.service_rate !== ""
      ? Number(sett.service_rate) : 0,
    tax_category: sett?.tax_category ?? "",
    tax_rate: sett?.tax_rate !== undefined && sett?.tax_rate !== null && sett?.tax_rate !== ""
      ? Number(sett.tax_rate) : 0,
  });

  // Initialize form with live default values
  const form = useForm<z.infer<typeof ServiceTaxFormSchema>>({
    resolver: zodResolver(ServiceTaxFormSchema),
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

  function onSubmit(data: z.infer<typeof ServiceTaxFormSchema>) {
    let filteredValues = Object.fromEntries(
      Object.entries(data).filter(([_, value]) =>
        value !== undefined && value !== null && value !== ""
      )
    ) as Record<string, unknown>;

    if (filteredValues.currency) {
      filteredValues = {
        ...filteredValues,
        currency: data.currency.toUpperCase(),
      }
    }

    update(JSON.stringify(filteredValues), {
      onSuccess: async () => {
        await queryClient.invalidateQueries({ queryKey: ['store.settings'] })
        toast.success("Service and tax rate data updated successfully.");
      },
      onError: async (error) => {
        if (error && isHTTPResponse<null>(error)) {
          if (typeof error.data === "string") {
            toast.error(error.data);
            return;
          }

          if (typeof error.data === "object" && error.data !== null) {
            const data = error.data as ServiceAndTaxRateErrorResponse;

            if (data.service_rate && data.service_rate.length > 0) {
              form.setError("service_rate", {type: "manual", message: data.service_rate[0]})
            }

            if (data.tax_rate && data.tax_rate.length > 0) {
              form.setError("tax_rate", {type: "manual", message: data.tax_rate[0]})
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
                <div className="mt-2 space-y-6">
                  <FormField
                    control={form.control}
                    name="currency"
                    render={({ field }) => (
                      <FormItem className="w-full">
                        <FormLabel>Currency</FormLabel>
                        <FormControl>
                          <Select
                            value={field.value}
                            onValueChange={field.onChange}
                          >
                            <SelectTrigger className="w-full">
                              <SelectValue placeholder="Select Currency" />
                            </SelectTrigger>
                            <SelectContent position="popper">
                              <SelectItem value="idr">IDR - Indonesia Rupiah</SelectItem>
                            </SelectContent>
                          </Select>
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />

                  <div className="flex items-center space-x-2 text-sm">
                    <div className="flex-grow h-px bg-gray-200" />
                    <span className="whitespace-nowrap">Service Category & Rate</span>
                    <div className="flex-grow h-px bg-gray-200" />
                  </div>

                  <div className="flex flex-col md:flex-row gap-2">
                    <FormField
                      control={form.control}
                      name="service_category"
                      render={({ field }) => (
                        <FormItem className="w-full">
                          <FormControl>
                            <Select
                              value={field.value}
                              onValueChange={field.onChange}
                              disabled={!form.watch("currency")}
                            >
                              <SelectTrigger className="w-full">
                                <SelectValue placeholder="Select Service Category" />
                              </SelectTrigger>
                              <SelectContent position="popper">
                                <SelectItem value="standard">Standard – Percentage-based</SelectItem>
                                <SelectItem value="nominal">Nominal – Fixed amount</SelectItem>
                              </SelectContent>
                            </Select>
                          </FormControl>
                          <FormMessage />
                        </FormItem>
                      )}
                    />

                    <FormField
                      control={form.control}
                      name="service_rate"
                      render={({ field }) => {
                        const {ref, name, value, onChange, ...rest} = field

                        return (
                          <FormItem className="w-full">
                            <FormControl>
                              <div className="relative">
                                <Input
                                  type="number"
                                  placeholder="Service Rate"
                                  disabled={!form.watch("service_category")}
                                  value={value ?? ""}
                                  name={name}
                                  ref={ref}
                                  onChange={(e) => field.onChange(e.target.valueAsNumber)}
                                  {...rest}
                                />
                                <div className="absolute inset-y-0 right-0 flex items-center px-3 text-sm bg-muted rounded-r-md uppercase">
                                  {form.watch("service_category") === "standard" ? "% / TRX" : `${form.watch("currency")} / TRX`}
                                </div>
                              </div>

                            </FormControl>
                            <FormMessage />
                          </FormItem>
                        )
                      }}
                    />
                  </div>

                  <div className="flex items-center space-x-2 text-sm">
                    <div className="flex-grow h-px bg-gray-200" />
                    <span className="whitespace-nowrap">Tax Category & Rate</span>
                    <div className="flex-grow h-px bg-gray-200" />
                  </div>

                  <div className="flex flex-col md:flex-row gap-2">
                    <FormField
                      control={form.control}
                      name="tax_category"
                      render={({ field }) => (
                        <FormItem className="w-full">
                          <FormControl>
                            <Select
                              value={field.value}
                              onValueChange={field.onChange}
                              disabled={!form.watch("currency")}
                            >
                              <SelectTrigger className="w-full">
                                <SelectValue placeholder="Select Tax Category" />
                              </SelectTrigger>
                              <SelectContent position="popper">
                                <SelectItem value="standard">Standard – Percentage-based</SelectItem>
                                <SelectItem value="nominal">Nominal – Fixed amount</SelectItem>
                              </SelectContent>
                            </Select>
                          </FormControl>
                          <FormMessage />
                        </FormItem>
                      )}
                    />

                    <FormField
                      control={form.control}
                      name="tax_rate"
                      render={({ field }) => {
                        const {ref, name, value, onChange, ...rest} = field

                        return (
                          <FormItem className="w-full">
                            <FormControl>
                              <div className="relative">
                                <Input
                                  type="number"
                                  placeholder="Service Rate"
                                  disabled={!form.watch("tax_category")}
                                  value={value ?? ""}
                                  name={name}
                                  ref={ref}
                                  onChange={(e) => field.onChange(e.target.valueAsNumber)}
                                  {...rest}
                                />
                                <div className="absolute inset-y-0 right-0 flex items-center px-3 text-sm bg-muted rounded-r-md uppercase">
                                  {form.watch("tax_category") === "standard" ? "% / TRX" : `${form.watch("currency")} / TRX`}
                                </div>
                              </div>
                            </FormControl>
                            <FormMessage />
                          </FormItem>
                        )
                      }}
                    />
                  </div>
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
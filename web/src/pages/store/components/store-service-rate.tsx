import {Card, CardContent, CardDescription, CardHeader, CardTitle} from "@/components/ui/card.tsx";
import {Input} from "@/components/ui/input.tsx";
import {Button} from "@/components/ui/button.tsx";
import {Store} from "@/lib/types/store.ts";
import {z} from "zod";
import {useForm} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import {useEffect} from "react";
import {Form, FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/components/ui/form.tsx";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select.tsx";

interface StoreServiceRateProps {
  store?: Store | null;
}

const formSchema = z.object({
  currency: z.string(),
  currency_rate: z.string(),
  service_category: z.string(),
  service_rate: z.string(),
  tax_category: z.string(),
  tax_rate: z.string(),
})

export const StoreServiceRate = (
  {store}: StoreServiceRateProps
) => {
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      currency: "",
      currency_rate: "",
      service_category: "",
      service_rate: "",
      tax_category: "",
      tax_rate: "",
    },
  })

  useEffect(() => {
    if (store) {
      form.setValue("currency", store.currency)
      form.setValue("currency_rate", store.currency_rate)
      form.setValue("service_category", store.service_category)
      form.setValue("service_rate", store.service_rate)
      form.setValue("tax_category", store.tax_category)
      form.setValue("tax_rate", store.tax_rate)
    }
  }, [store])

  function onSubmit(values: z.infer<typeof formSchema>) {
    // Do something with the form values.
    // ✅ This will be type-safe and validated.
    console.log(values)
  }

  return (
    <Card className="w-full">
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)}>
          <div className="flex justify-between items-center">
            <CardHeader>
              <CardTitle>Service Rate</CardTitle>
              <CardDescription>
                Store Currency, Service and Tax Setting
              </CardDescription>
            </CardHeader>
            <Button type="submit" className="mr-6">Save Change</Button>
          </div>
          <CardContent className="space-y-2">
            <FormField
              control={form.control}
              name="currency"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Currency</FormLabel>
                  <FormControl>
                    <Input placeholder="Currency" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="currency_rate"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Currency Rate</FormLabel>
                  <FormControl>
                    <div className="relative">
                      <Input placeholder="Currency Rate" {...field} />
                      <div className="absolute inset-y-0 right-0 flex items-center px-6 bg-muted rounded-r-md font-semibold">
                        {store?.currency} / 1 USD
                      </div>
                    </div>
                  </FormControl>
                  <FormMessage/>
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="service_category"
              render={({field}) => (
                <FormItem>
                  <FormLabel>Service Category</FormLabel>
                  <Select
                    onValueChange={field.onChange}
                    defaultValue={store?.service_category}
                  >
                    <FormControl>
                      <SelectTrigger>
                        <SelectValue placeholder="Select a service category"/>
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      <SelectItem
                        key="standard"
                        value="standard"
                      >
                        Standard
                      </SelectItem>
                    </SelectContent>
                  </Select>
                  <FormMessage/>
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="service_rate"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Service Rate</FormLabel>
                  <FormControl>
                    <div className="relative">
                      <Input placeholder="Currency Rate" {...field} />
                      <div className="absolute inset-y-0 right-0 flex items-center px-6 bg-muted rounded-r-md font-semibold">
                        % / TRX
                      </div>
                    </div>
                  </FormControl>
                  <FormMessage/>
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="tax_category"
              render={({field}) => (
                <FormItem>
                  <FormLabel>Tax Category</FormLabel>
                  <Select
                    onValueChange={field.onChange}
                    defaultValue={store?.service_category}
                  >
                    <FormControl>
                      <SelectTrigger>
                        <SelectValue placeholder="Select a tax category"/>
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      <SelectItem
                        key="standard"
                        value="standard"
                      >
                        Standard
                      </SelectItem>
                    </SelectContent>
                  </Select>
                  <FormMessage/>
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="tax_rate"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Tax Rate</FormLabel>
                  <FormControl>
                    <div className="relative">
                      <Input placeholder="Currency Rate" {...field} />
                      <div className="absolute inset-y-0 right-0 flex items-center px-6 bg-muted rounded-r-md font-semibold">
                        % / TRX
                      </div>
                    </div>
                  </FormControl>
                  <FormMessage/>
                </FormItem>
              )}
            />
          </CardContent>
        </form>
      </Form>
    </Card>
  )
}
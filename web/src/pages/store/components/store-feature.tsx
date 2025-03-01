import {Card, CardContent, CardDescription, CardHeader, CardTitle} from "@/components/ui/card.tsx";
import {Button} from "@/components/ui/button.tsx";
import {Store} from "@/lib/types/store.ts";
import {z} from "zod";
import {useForm} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import {useEffect} from "react";
import {Form, FormControl, FormDescription, FormField, FormItem, FormLabel} from "@/components/ui/form.tsx";
import {Switch} from "@/components/ui/switch.tsx";

interface StoreFeatureProps {
  store?: Store | null;
}

const formSchema = z.object({
  feature_floor: z.boolean().default(false).optional(),
  feature_room: z.boolean().default(false).optional(),
  feature_table: z.boolean().default(false).optional(),
})

export const StoreFeature = (
  {store}: StoreFeatureProps
) => {

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      feature_floor: false,
      feature_room: false,
      feature_table: false,
    },
  })

  useEffect(() => {
    if (store) {
      form.setValue("feature_floor", store.feature_floor === "1")
      form.setValue("feature_room", store.feature_room === "1")
      form.setValue("feature_table", store.feature_table === "1")
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
              <CardTitle>Features</CardTitle>
              <CardDescription>
                Store Feature Enabled
              </CardDescription>
            </CardHeader>
            <Button type="submit" className="mr-6">Save Change</Button>
          </div>
          <CardContent className="space-y-2">
            <FormField
              control={form.control}
              name="feature_floor"
              render={({ field }) => (
                <FormItem className="flex flex-row items-center justify-between rounded-lg border p-4">
                  <div className="space-y-0.5">
                    <FormLabel className="text-base">
                      Floor
                    </FormLabel>
                    <FormDescription>
                      Enable Store Floor for better blueprint layout
                    </FormDescription>
                  </div>
                  <FormControl>
                    <Switch
                      checked={field.value}
                      onCheckedChange={field.onChange}
                    />
                  </FormControl>
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="feature_room"
              render={({ field }) => (
                <FormItem className="flex flex-row items-center justify-between rounded-lg border p-4">
                  <div className="space-y-0.5">
                    <FormLabel className="text-base">
                      Room
                    </FormLabel>
                    <FormDescription>
                      Enable Store Room for better blueprint layout
                    </FormDescription>
                  </div>
                  <FormControl>
                    <Switch
                      checked={field.value}
                      onCheckedChange={field.onChange}
                    />
                  </FormControl>
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="feature_table"
              render={({ field }) => (
                <FormItem className="flex flex-row items-center justify-between rounded-lg border p-4">
                  <div className="space-y-0.5">
                    <FormLabel className="text-base">
                      Table
                    </FormLabel>
                    <FormDescription>
                      Enable Store Table for better blueprint layout
                    </FormDescription>
                  </div>
                  <FormControl>
                    <Switch
                      checked={field.value}
                      onCheckedChange={field.onChange}
                    />
                  </FormControl>
                </FormItem>
              )}
            />
          </CardContent>
        </form>
      </Form>
    </Card>
  )
}
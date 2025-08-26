import {z} from "zod";
import {useForm} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import {Card, CardContent, CardFooter} from "@/components/ui/card";
import {Button} from "@/components/ui/button";
import {Form, FormControl, FormDescription, FormField, FormItem, FormLabel} from "@/components/ui/form";
import {Switch} from "@/components/ui/switch";
import {useStoreState} from "@/states/store-state";
import {useEffect} from "react";

const formSchema = z.object({
  feature_floor: z.boolean().default(false).optional(),
})

export function FeatureSection() {
  const {settings} =  useStoreState();

  const getResetValues = (sett: typeof settings | null) => ({
    feature_floor: sett?.feature_floor === "1",
  });

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: getResetValues(settings ?? null),
  })

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


  function onSubmit(values: z.infer<typeof formSchema>) {
    // Do something with the form values.
    // ✅ This will be type-safe and validated.
    console.log(values)
  }

  return (
    <div className="space-y-6 select-none">
      <Card>
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)}>
            <CardContent className="mb-6 space-y-3">
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
                        Enable floor plans to organize tables and rooms by level.
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

            <CardFooter className="flex gap-2 justify-end border-t-2 pt-4">
              <Button
                variant="outline"
                size="sm"
                type="button"
                className="text-xs cursor-pointer"
              >Cancel</Button>
              <Button
                size="sm"
                className="text-xs cursor-pointer"
              >
                {/*{isPending && <Loader2 className="w-4 animate-spin mr-1" />}*/}
                Save
              </Button>
            </CardFooter>
          </form>
        </Form>
      </Card>
    </div>
  )
}
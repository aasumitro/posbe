import {z} from "zod";
import {useForm} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import {Card, CardContent, CardFooter} from "@/components/ui/card";
import {Button} from "@/components/ui/button";
import {Form, FormControl, FormDescription, FormField, FormItem, FormLabel} from "@/components/ui/form";
import {Switch} from "@/components/ui/switch";
import {useStoreState} from "@/states/store-state";
import {useEffect, useMemo} from "react";
import {useUpdateSetting} from "@/hooks/use-store";
import {useQueryClient} from "@tanstack/react-query";
import {Loader2Icon} from "lucide-react";
import {isHTTPResponse} from "@/lib/api";
import {toast} from "sonner";

const formSchema = z.object({
  feature_floor: z.boolean().default(false).optional(),
})

export function FeatureSection() {
  const {settings} =  useStoreState();
  const {mutate: update, isPending} = useUpdateSetting()
  const queryClient = useQueryClient();

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

  function onSubmit(values: z.infer<typeof formSchema>) {
    const floorFlag = values.feature_floor ? "1" : "0"
    update(JSON.stringify({feature_floor: floorFlag,}), {
      onSuccess: async () => {
        await queryClient.invalidateQueries({ queryKey: ['store.settings'] })
        toast.success(`Floor feature ${floorFlag ? "enabled" : "disabled"}`);
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
    })
  }

  return (
    <div className="space-y-6 select-none">
      <Card>
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)}>
            <CardContent className="space-y-3">
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

            {isDirtyComparedToSettingData && (
              <CardFooter className="mt-6 flex gap-2 justify-end border-t-2 pt-4">
                <Button
                  variant="outline"
                  size="sm"
                  type="button"
                  className="text-xs cursor-pointer"
                  onClick={restoreChanges}
                >Cancel</Button>
                <Button
                  size="sm"
                  className="text-xs cursor-pointer"
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
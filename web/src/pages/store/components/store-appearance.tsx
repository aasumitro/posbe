import {Store} from "@/lib/types/store.ts";
import {Card, CardContent, CardDescription, CardHeader, CardTitle} from "@/components/ui/card.tsx";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage
} from "@/components/ui/form.tsx";
import {Button, buttonVariants} from "@/components/ui/button.tsx";
import {z} from "zod";
import {useForm} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import {useEffect} from "react";
import {RadioGroup, RadioGroupItem} from "@/components/ui/radio-group.tsx";
import {cn} from "@/lib/utils.ts";
import {ChevronDownIcon} from "lucide-react";

interface StoreAppSettingProps {
  store?: Store | null;
}

const appearanceFormSchema = z.object({
  fe_theme: z.enum(["light", "dark"], {
    required_error: "Please select a theme.",
  }),
  fe_lang: z.enum(["en_US", "id_ID"], {
    invalid_type_error: "Select a language",
    required_error: "Please select a language.",
  }),
  fe_locale: z.enum(["Asia/Singapore", "Asia/Makassar"], {
    invalid_type_error: "Select a timezone",
    required_error: "Please select a timezone.",
  }),
})

type AppearanceFormValues = z.infer<typeof appearanceFormSchema>

// This can come from your database or API.
const defaultValues: Partial<AppearanceFormValues> = {
  fe_theme: "light",
}

export const StoreAppearance = (
  {store}: StoreAppSettingProps
) => {
  const form = useForm<AppearanceFormValues>({
    resolver: zodResolver(appearanceFormSchema),
    defaultValues,
  })

  useEffect(() => {
    if (store) {
      // @ts-ignore
      form.setValue("fe_lang", store.fe_lang)
      // @ts-ignore
      form.setValue("fe_theme", store.fe_theme)
      // @ts-ignore
      form.setValue("fe_locale", store.fe_locale)
    }
  }, [store])

  function onSubmit(values: AppearanceFormValues) {
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
              <CardTitle>Appearance</CardTitle>
              <CardDescription>
                Customize the appearance of the app.
              </CardDescription>
            </CardHeader>
            <Button type="submit" className="mr-6">Save Change</Button>
          </div>
          <CardContent className="space-y-4">
            <div className="flex gap-8">
              <FormField
                control={form.control}
                name="fe_lang"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Language</FormLabel>
                    <div className="relative w-max">
                      <FormControl>
                        <select
                          className={cn(
                            buttonVariants({ variant: "outline" }),
                            "w-[200px] appearance-none font-normal"
                          )}
                          {...field}
                        >
                          <option value="en_US">English</option>
                          <option value="id_ID">Indonesia</option>
                        </select>
                      </FormControl>
                      <ChevronDownIcon className="absolute right-3 top-2.5 h-4 w-4 opacity-50" />
                    </div>
                    <FormDescription>
                      Set the language you want to use.
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name="fe_locale"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Timezone</FormLabel>
                    <div className="relative w-max">
                      <FormControl>
                        <select
                          className={cn(
                            buttonVariants({ variant: "outline" }),
                            "w-[200px] appearance-none font-normal"
                          )}
                          {...field}
                        >
                          <option value="Asia/Singapore">Asia/Singapore</option>
                          <option value="Asia/Makassar">Asia/Makassar</option>
                        </select>
                      </FormControl>
                      <ChevronDownIcon className="absolute right-3 top-2.5 h-4 w-4 opacity-50" />
                    </div>
                    <FormDescription>
                      Set the store timezone.
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </div>
            <FormField
              control={form.control}
              name="fe_theme"
              render={({ field }) => (
                <FormItem className="space-y-1">
                  <FormLabel>Theme</FormLabel>
                  <FormDescription>
                    Select the theme for the dashboard.
                  </FormDescription>
                  <FormMessage />
                  <RadioGroup
                    onValueChange={field.onChange}
                    defaultValue={field.value}
                    className="grid max-w-md grid-cols-2 gap-8 pt-2"
                  >
                    <FormItem>
                      <FormLabel className="[&:has([data-state=checked])>div]:border-primary">
                        <FormControl>
                          <RadioGroupItem value="light" className="sr-only" />
                        </FormControl>
                        <div className="items-center rounded-md border-2 border-muted p-1 hover:border-accent">
                          <div className="space-y-2 rounded-sm bg-[#ecedef] p-2">
                            <div className="space-y-2 rounded-md bg-white p-2 shadow-sm">
                              <div className="h-2 w-[80px] rounded-lg bg-[#ecedef]" />
                              <div className="h-2 w-[100px] rounded-lg bg-[#ecedef]" />
                            </div>
                            <div className="flex items-center space-x-2 rounded-md bg-white p-2 shadow-sm">
                              <div className="h-4 w-4 rounded-full bg-[#ecedef]" />
                              <div className="h-2 w-[100px] rounded-lg bg-[#ecedef]" />
                            </div>
                          </div>
                        </div>
                        <span className="block w-full p-2 text-center font-normal">
                      Light
                    </span>
                      </FormLabel>
                    </FormItem>
                    <FormItem>
                      <FormLabel className="[&:has([data-state=checked])>div]:border-primary">
                        <FormControl>
                          <RadioGroupItem value="dark" className="sr-only" />
                        </FormControl>
                        <div className="items-center rounded-md border-2 border-muted bg-popover p-1 hover:bg-accent hover:text-accent-foreground">
                          <div className="space-y-2 rounded-sm bg-slate-950 p-2">
                            <div className="space-y-2 rounded-md bg-slate-800 p-2 shadow-sm">
                              <div className="h-2 w-[80px] rounded-lg bg-slate-400" />
                              <div className="h-2 w-[100px] rounded-lg bg-slate-400" />
                            </div>
                            <div className="flex items-center space-x-2 rounded-md bg-slate-800 p-2 shadow-sm">
                              <div className="h-4 w-4 rounded-full bg-slate-400" />
                              <div className="h-2 w-[100px] rounded-lg bg-slate-400" />
                            </div>
                          </div>
                        </div>
                        <span className="block w-full p-2 text-center font-normal">
                      Dark
                    </span>
                      </FormLabel>
                    </FormItem>
                  </RadioGroup>
                </FormItem>
              )}
            />
          </CardContent>
        </form>
      </Form>
    </Card>
  )
}
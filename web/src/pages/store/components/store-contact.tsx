import {Card, CardContent, CardDescription, CardHeader, CardTitle} from "@/components/ui/card.tsx";
import {Input} from "@/components/ui/input.tsx";
import {Button} from "@/components/ui/button.tsx";
import {Store} from "@/lib/types/store.ts";
import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import { z } from "zod"
import {Form, FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/components/ui/form.tsx";
import {useEffect} from "react";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select.tsx";
import {capitalize} from "@/lib/str.ts";
import {Textarea} from "@/components/ui/textarea.tsx";

interface StoreContactProps {
  store?: Store | null;
}

const formSchema = z.object({
  name: z.string().min(5).max(16),
  phone:  z.string().refine((value) => /^[+]{1}(?:[0-9-()/.]\s?){6,15}[0-9]{1}$/.test(value), {
    message: "Invalid phone number"
  }),
  email:  z.string().email(),
  address: z.string().min(5).max(55),
  pos_type: z.string(),
})

const postType = ["none", "restaurant", "coffee_shop", "store", "karaoke"]

export const StoreContact = (
  {store}: StoreContactProps,
) => {
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: "",
      phone: "",
      email: "",
      address: "",
      pos_type: "",
    },
  })

  useEffect(() => {
    if (store) {
      form.setValue("name", store.name)
      form.setValue("phone", store.phone)
      form.setValue("email", store.email)
      form.setValue("address", store.address)
      form.setValue("pos_type", store.pos_type)
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
              <CardTitle>Contact</CardTitle>
              <CardDescription>
                Store Contact Information
              </CardDescription>
            </CardHeader>

            <Button type="submit" className="mr-6">Save Change</Button>
          </div>
          <CardContent className="space-y-2">
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Store Name</FormLabel>
                  <FormControl>
                    <Input placeholder="Name" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="phone"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Phone</FormLabel>
                  <FormControl>
                    <Input placeholder="Phone" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="email"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Email</FormLabel>
                  <FormControl>
                    <Input placeholder="Store Email" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="pos_type"
              render={({field}) => (
                <FormItem>
                  <FormLabel>Pos Type</FormLabel>
                  <Select
                    onValueChange={field.onChange}
                    defaultValue={store?.pos_type}
                  >
                    <FormControl>
                      <SelectTrigger>
                        <SelectValue placeholder="Select a pos type"/>
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      {postType?.map((type) => (
                        <SelectItem
                          key={type}
                          value={type}
                        >
                          {capitalize(type.replace("_", " "))}
                        </SelectItem>
                      ))}
                    </SelectContent>
                  </Select>
                  <FormMessage/>
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="address"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Address</FormLabel>
                  <FormControl>
                    <Textarea
                      placeholder="Store Address"
                      rows={5}
                      className="resize-none"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </CardContent>
        </form>
      </Form>
    </Card>
  )
}
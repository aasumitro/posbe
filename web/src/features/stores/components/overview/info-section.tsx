import {z} from "zod";
import {useForm} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import {Card, CardContent, CardFooter} from "@/components/ui/card";
import {Form, FormControl, FormField, FormItem, FormMessage} from "@/components/ui/form";
import {Button} from "@/components/ui/button";
import {Input} from "@/components/ui/input";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select";

const StoreInfoFormSchema = z.object({
  name: z.string().min(3).max(100),
  phone: z.string(),
  email: z.string(),
  type: z.string(),
  address_line1: z.string(),
  address_line2: z.string(),
  country: z.string(),
  state_province: z.string(),
  city_regency: z.string(),
  postal_code: z.string(),
})

export function InfoSection() {
  const getResetValues = () => ({
    name: "",
    phone: "",
    email: "",
    type: "",
    address_line1: "",
    address_line2: "",
    country: "",
    state_province: "",
    city_regency: "",
    postal_code: "",
  });

  // Initialize form with live default values
  const form = useForm<z.infer<typeof StoreInfoFormSchema>>({
    resolver: zodResolver(StoreInfoFormSchema),
    defaultValues: getResetValues(),
  });

  function onSubmit(data: z.infer<typeof StoreInfoFormSchema>) {
    console.log(data);
  }

  return (
    <div className="space-y-6">
      <Card>
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)}>
            <CardContent className="mb-6">
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

            <CardFooter className="flex gap-2 justify-end border-t-2 pt-4">
              <Button
                variant="outline"
                size="sm"
                type="button"
                className="text-xs"
              >Cancel</Button>
              <Button
                size="sm"
                className="text-xs"
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
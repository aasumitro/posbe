import {createFileRoute} from "@tanstack/react-router";
import { z } from "zod";
import {useForm} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage
} from "@/components/ui/form";
import {Input} from "@/components/ui/input";
import {Button} from "@/components/ui/button";
import {PasswordField} from "@/components/ui/password-input";

export const Route = createFileRoute("/login")({
  component: LoginPage,
})

const loginSchema = z.object({
  username:  z.string().min(5).max(10),
  password:  z.string().min(5).max(16),
});

function LoginPage() {
  const form = useForm<z.infer<typeof loginSchema>>({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      username: "",
      password: "",
    },
  });

  function onSubmit(values: z.infer<typeof loginSchema>) {
    console.log(values)
  }

  return (
    <section className="bg-gray-50 dark:bg-gray-900">
      <div className="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0">
        <div className="mb-8 text-center">
          <p className="text-2xl font-bold italic text-gray-900 dark:text-white">
            P<span className="font-mono">0</span>S
          </p>
          <span className="text-medium font-light mt-3 mb-6">Point of sales</span>
        </div>
        <div
          className="w-full bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700">
          <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
            <h1 className="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white">Sign in to your account</h1>
            <Form {...form}>
              <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4 flex flex-col">
                <FormField
                  control={form.control}
                  name="username"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Username</FormLabel>
                      <FormControl>
                        <Input
                          {...field}
                          placeholder="@posbe"
                        />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />

                <PasswordField
                  label="Password"
                  placeholder="* * * * * *"
                  {...form.register("password")}
                />

                <Button
                  type="submit"
                  className="ml-auto"
                  // disabled={signIn.isLoading}
                >
                  {/*{signIn.isLoading && <Loader2Icon className="w-4 h-4 mr-2 animate-spin"/>}*/}
                  Submit
                </Button>
              </form>
            </Form>
          </div>
        </div>
      </div>
    </section>
  )
}
import {createFileRoute, redirect} from "@tanstack/react-router";
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
import {useAuthStore} from "@/states/auth-state";
import {useLogin} from "@/hooks/use-auth";
import {Loader2Icon} from "lucide-react";
import type {User} from "@/types/user";
import {isHTTPResponse} from "@/lib/api";
import {toast} from "sonner";

export const Route = createFileRoute("/login")({
  beforeLoad: () => {
    const authState = useAuthStore.getState().auth;
    if (authState.accessToken) {
      const searchParams = new URLSearchParams(location.search);
      const redirectTo = searchParams.get("redirect") ?? "/";
      throw redirect({ to: redirectTo, replace: true });
    }
  },
  component: LoginPage,
})

const loginSchema = z.object({
  username:  z.string().min(5).max(10),
  password:  z.string().min(5).max(16),
});

function LoginPage() {
  const { mutate: login, isPending} = useLogin();
  const {auth} = useAuthStore();

  const form = useForm<z.infer<typeof loginSchema>>({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      username: "",
      password: "",
    },
  });

  function onSubmit(values: z.infer<typeof loginSchema>) {
    login({username: values.username, password: values.password}, {
      onSuccess: async (response) => {
        auth.setRefreshToken(response.data?.token?.refresh_token as string)
        auth.setAccessToken(response.data?.token?.access_token as string);
        auth.setUserId(response.data?.user?.id as number)
        auth.setUser(response.data?.user as User)

        const searchParams = new URLSearchParams(location.search);
        const redirectTo = searchParams.get("redirect") ?? "/";
        toast.success( `redirecting to ${redirectTo === "/" ? "home" : redirectTo} . . .`);
        window.location.href = redirectTo;
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
    <section className="hidden lg:flex bg-gray-50 dark:bg-gray-900">
      <div className="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0">
        <div className="mb-8 text-center">
          <p className="text-4xl font-bold italic text-gray-900 dark:text-white">
            P<span className="font-mono">0</span>S
          </p>
          <span className="text-medium font-light mt-3 mb-6">Point of sales</span>
        </div>
        <div
          className="min-w-[430px] w-full bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700">
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

                <div className="flex justify-between">
                  {/*<Button variant="link" className="p-0 cursor-pointer" type="button" disabled={isPending}>*/}
                  {/*  Forgot password?*/}
                  {/*</Button>*/}

                  <Button
                    type="submit"
                    className="ml-auto cursor-pointer"
                    disabled={isPending}
                  >
                    {isPending && (
                      <Loader2Icon
                        className="w-4 h-4 animate-spin"
                      />
                    )}
                    Submit
                  </Button>
                </div>
              </form>
            </Form>
          </div>
        </div>
      </div>
    </section>
  )
}
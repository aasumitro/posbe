import {useNavigate} from "react-router-dom";
import {useEffect} from "react";
import {KEYS} from "@/lib/keys.ts";
import {useAuth} from "@/hooks/use-auth.tsx";
import {useMutation} from "react-query";
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
} from "@/components/ui/form.tsx";
import {Input} from "@/components/ui/input.tsx";
import {Button} from "@/components/ui/button.tsx";
import {Loader2Icon} from "lucide-react";
import {ToastMessageContainer} from "@/components/toast-message-container.tsx";
import {useSessionStateStore} from "@/stores/session-state.ts";

const loginSchema = z.object({
  username:  z.string().min(5).max(10),
  password:  z.string().min(5).max(16),
});

export const LoginPage = () => {
  const {login} = useAuth();
  const {setUserToken, setUserData} =  useSessionStateStore();
  const navigate = useNavigate();
  const form = useForm<z.infer<typeof loginSchema>>({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      username: "",
      password: "",
    },
  });

  useEffect(() => {
    if (localStorage.getItem(KEYS.UI_AUTH_TOKEN)) {
      navigate("/");
      return;
    }
  }, []);

  const signIn = useMutation(login, {
    onSuccess: (resp) => {
      if (resp.code != 201) {
        if (resp.code == 404) {
          ToastMessageContainer("User Not Found", "We couldn't find a user with that username. Please check your username and try again.");
          return;
        }
        if (resp.code == 422) {
          ToastMessageContainer(resp.data, "The password you entered does not match the username. Please try again.");
          return;
        }
        ToastMessageContainer("Something went wrong", resp.data);
        return;
      }
      const token = resp.data?.token
      const profile = resp.data?.user
      if (token) {
        localStorage.setItem(KEYS.UI_AUTH_TOKEN, token);
        setUserToken(token);
      }
      if (profile) {
        localStorage.setItem(KEYS.UI_USER_PROFILE,
          JSON.stringify(profile));
        setUserData(profile);
      }
      navigate("/");
    },
  });

  const onSubmit = (values: z.infer<typeof loginSchema>)  => {
    signIn.mutate(JSON.stringify({username: values.username, password: values.password}));
  }

  return (
    <section className="bg-gray-50 dark:bg-gray-900">
      <div className="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0">
        <div className="mb-8 text-center">
          <p className="text-2xl font-bold italic text-gray-900 dark:text-white">POSBE</p>
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
                      <FormLabel>Your Username</FormLabel>
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
                <FormField
                  control={form.control}
                  name="password"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Your Password</FormLabel>
                      <FormControl>
                        <Input
                          {...field}
                          type="password"
                          placeholder="* * * * * *"
                        />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
                <Button
                  type="submit"
                  className="ml-auto"
                  disabled={signIn.isLoading}
                >
                  {signIn.isLoading && <Loader2Icon className="w-4 h-4 mr-2 animate-spin"/>}
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
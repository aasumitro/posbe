import {
  AlertDialog,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog"
import {useEffect, useState} from "react";
import {useGlobalStateStore} from "@/stores/global-state.ts";
import {Loader2Icon} from "lucide-react";
import {KEYS} from "@/lib/keys.ts";
import {useStorePageState} from "@/stores/store-state.ts";
import {z} from "zod";
import {useForm} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import {Form, FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/components/ui/form.tsx";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select.tsx";
import {Input} from "@/components/ui/input.tsx";
import {Button} from "@/components/ui/button.tsx";
import {useStoreUser} from "@/hooks/use-store-user.tsx";
import {HTTPStatusCode, isHttpResponse} from "@/lib/api.ts";
import {ToastMessageContainer} from "@/components/toast-message-container.tsx";
import {useMutation, useQueryClient} from "react-query";

const newUserSchema = z.object({
  role_id: z.string(),
  name:  z.string().min(5).max(25),
  username:  z.string().min(5).max(12),
  password: z.string().min(6).max(12),
  email: z.string().email().optional().or(z.literal('')),
  phone: z.string().optional().refine((value) =>
    !value || /^[+]{1}(?:[0-9-()/.]\s?){6,15}[0-9]{1}$/.test(value), {
      message: "Invalid phone number",
    }),
});

export function NewEmployeeModal() {
  const [open, setOpen] = useState(false)
  const {boolStates, setBoolState} = useGlobalStateStore();
  const {roles} = useStorePageState();
  const {newUser} = useStoreUser();
  const queryClient = useQueryClient()

  const form = useForm<z.infer<typeof newUserSchema>>({
    resolver: zodResolver(newUserSchema),
    defaultValues: {
      role_id: "",
      name: "",
      username: "",
      password: "",
      email: undefined,
      phone: undefined,
    },
  });

  const createUserAccount = useMutation(newUser, {
    onSuccess: async (resp) => {
      if (isHttpResponse(resp) && resp.code != HTTPStatusCode.Created) {
        console.log(resp)
        // ToastMessageContainer("Something went wrong", resp);
        return;
      }
      await queryClient.invalidateQueries('store.users')
      ToastMessageContainer("User Created", "New user account has been crated successfully.");
      onClose();
    },
  });

  useEffect(() => {
    if (boolStates[KEYS.DISPLAY_NEW_EMPLOYEE_MODAL]) {
      setOpen(true);
    }
  }, [boolStates])

  const onClose = () => {
    setOpen(!open);
    setBoolState(KEYS.DISPLAY_NEW_EMPLOYEE_MODAL, false);
    form.setValue("role_id", "")
    form.setValue("name", "")
    form.setValue("username", "")
    form.setValue("password", "")
    form.setValue("email", undefined)
    form.setValue("phone", undefined)
  }

  const onSubmit = (values: z.infer<typeof newUserSchema>) => {
    if (values.role_id == "") {
      form.setError("role_id", {
        type: 'manual',
        message: 'Please select an role',
      })
      return;
    }

    let body: Record<string, any> = {
      role_id: Number(values.role_id),
      name: values.name,
      username: values.username,
      password: values.password
    }

    if (values.email && values.email !== "") {
      body.email = values.email
    }

    if (values.phone && values.phone !== "") {
      body.phone = values.phone
    }

    createUserAccount.mutate(JSON.stringify(body))
  }

  return (
    <AlertDialog open={open} onOpenChange={onClose}>
      <AlertDialogTrigger asChild></AlertDialogTrigger>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Create new User</AlertDialogTitle>
          <AlertDialogDescription>
            You are about to create a new user account. Please make sure that all the provided details are accurate. This action will grant access to the system, and the user will be able to log in with their credentials.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-2" >
            <FormField
              control={form.control}
              name="role_id"
              render={({field}) => (
                <FormItem>
                  <FormLabel>Role (*)</FormLabel>
                  <Select onValueChange={field.onChange} defaultValue={field.value}>
                    <FormControl>
                      <SelectTrigger>
                        <SelectValue placeholder="Select a role"/>
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      {roles?.map((role) => (
                        <SelectItem
                          key={role.id}
                          value={`${role.id}`}
                        >
                          {role.name}
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
              name="name"
              render={({field}) => (
                <FormItem>
                  <FormLabel>Name (*)</FormLabel>
                  <FormControl>
                    <Input
                      {...field}
                      placeholder="Lorem Ipsum"
                    />
                  </FormControl>
                  <FormMessage/>
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="username"
              render={({field}) => (
                <FormItem>
                  <FormLabel>Username (*)</FormLabel>
                  <FormControl>
                    <Input
                      {...field}
                      placeholder="@lorem"
                    />
                  </FormControl>
                  <FormMessage/>
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="email"
              render={({field}) => (
                <FormItem>
                  <FormLabel>Email</FormLabel>
                  <FormControl>
                    <Input
                      {...field}
                      placeholder="lorem@mail.posbe"
                    />
                  </FormControl>
                  <FormMessage/>
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="phone"
              render={({field}) => (
                <FormItem>
                  <FormLabel>Phone</FormLabel>
                  <FormControl>
                    <Input
                      {...field}
                      placeholder="+62XXXXXXXXXX"
                    />
                  </FormControl>
                  <FormMessage/>
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="password"
              render={({field}) => (
                <FormItem>
                  <FormLabel>Password (*)</FormLabel>
                  <FormControl>
                    <Input
                      {...field}
                      type="password"
                      placeholder="* * * * * *"
                    />
                  </FormControl>
                  <FormMessage/>
                </FormItem>
              )}
            />
            <div className="pt-4 w-full space-x-2 text-right">
              <Button type="button" variant="outline" onClick={onClose}>Cancel</Button>
              <Button type="submit" disabled={createUserAccount.isLoading}>
                {createUserAccount.isLoading && <Loader2Icon className="w-4 h-4 animate-spin mr-2" />}
                Save changes
              </Button>
            </div>
          </form>
        </Form>
      </AlertDialogContent>
    </AlertDialog>
  )
}
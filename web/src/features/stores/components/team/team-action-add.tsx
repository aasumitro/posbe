import {z} from "zod";
import {useState} from "react";
import {useForm} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import {
  AlertDialog, AlertDialogCancel,
  AlertDialogContent, AlertDialogDescription, AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger
} from "@/components/ui/alert-dialog";
import {cn} from "@/lib/utils";
import {IconCheck, IconPlus, IconX} from "@tabler/icons-react";
import {Form, FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/components/ui/form";
import {Input} from "@/components/ui/input";
import {Button} from "@/components/ui/button";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select";
import {PasswordField} from "@/components/ui/password-input";
import {useUserState} from "@/states/user-state";
import {useNewUser} from "@/hooks/use-user";
import {toast} from "sonner";
import {isHTTPResponse} from "@/lib/api";
import {useQueryClient} from "@tanstack/react-query";
import {Loader2Icon} from "lucide-react";

const FormSchema = z.object({
  username: z.string(),
  name: z.string({
    required_error: "Please set a name to display.",
  }).min(1, "Name must be at least 1 character.").max(20, "Name must be at most 20 characters."),
  email: z.string().email(),
  role: z.string(),
  password: z
    .string({ required_error: "Password cannot be empty." })
    .min(8, { message: "Minimum 8 characters." })
    .max(12, { message: "Maximum 12 characters." })
    .regex(/(?=.*[A-Z])/, { message: "At least one uppercase character." })
    .regex(/(?=.*[a-z])/, { message: "At least one lowercase character." })
    .regex(/(?=.*\d)/, { message: "At least one digit." })
    .regex(/^[A-Za-z0-9._@]+$/, { message: "Only letters, numbers, '.', '_' and '@' are allowed." }),
})

export function TeamActionAdd() {
  const [open, setOpen] = useState(false)
  const { roles } =  useUserState();
  const { mutate: addUser, isPending} = useNewUser();
  const queryClient = useQueryClient();

  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
    defaultValues: {
      username: "",
      name: "",
      email: "",
      role: "",
      password: ""
    }
  })

  const hasUppercase = /[A-Z]/.test(form.watch("password") || "");
  const hasLowercase = /[a-z]/.test(form.watch("password") || "");
  const hasDigit = /\d/.test(form.watch("password") || "");
  const hasAllowedChars = /[._@]/.test(form.watch("password") || "");

  function onSubmit(data: z.infer<typeof FormSchema>) {
    addUser(JSON.stringify({
      role_id: Number(data.role),
      username: data.username,
      name: data.name,
      email: data.email,
      password: data.password,
    }), {
      onSuccess: async () => {
        await queryClient.invalidateQueries({ queryKey: ['users'] })
        toast.success("New user added successfully");
        onOpenChange(false);
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

  function onOpenChange(state: boolean) {
    setOpen(state);
    if (!open) {
      form.reset();
    }
  }

  return (
    <AlertDialog open={open} onOpenChange={onOpenChange}>
      <AlertDialogTrigger asChild>
        <div className={cn(
          "aspect-video rounded-xl border-2 border-dashed hover:bg-muted/50",
          "flex flex-col items-center justify-center group cursor-pointer min-h-56 max-h-56 w-96"
        )}>
          <IconPlus className="w-12 h-12 text-gray-400 group-hover:text-gray-500" />
        </div>
      </AlertDialogTrigger>
      <AlertDialogContent className="select-none">
        <AlertDialogHeader>
          <AlertDialogTitle>Add Team Member</AlertDialogTitle>
          <AlertDialogDescription>
            Fill in the details below to add a new member to your team.
          </AlertDialogDescription>
        </AlertDialogHeader>

        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)}>
            <section className="grid gap-4 pt-2 pb-8">
              <FormField
                control={form.control}
                name="username"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Username</FormLabel>
                    <FormControl>
                      <Input
                        className="w-full h-10"
                        type="text"
                        autoComplete="off"
                        placeholder="@zeros"
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="name"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Name</FormLabel>
                    <FormControl>
                      <Input
                        className="w-full h-10"
                        type="text"
                        placeholder="Zeros Mardigu"
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
                  <FormItem>
                    <FormLabel>Email address</FormLabel>
                    <FormControl>
                      <Input
                        className="w-full h-10"
                        type="email"
                        placeholder="zeros@member.pos"
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="role"
                render={({field}) => (
                  <FormItem>
                    <FormLabel>Role</FormLabel>
                    <Select
                      onValueChange={field.onChange}
                    >
                      <FormControl>
                        <SelectTrigger className="w-full h-10">
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

              <PasswordField
                label="New password"
                placeholder="Enter your new password"
                {...form.register("password")}
              />

              <div className="flex flex-col gap-2 ml-3 text-sm">
                Must contain:
                <span className={`flex gap-2 items-center ${hasUppercase ? "text-green-600" : "text-muted-foreground"}`}>
                  {hasUppercase ? <IconCheck className="w-4 h-4" /> : <IconX className="w-4 h-4" />}
                  At least one uppercase character.
                </span>
                <span className={`flex gap-2 items-center ${hasLowercase ? "text-green-600" : "text-muted-foreground"}`}>
                  {hasLowercase ? <IconCheck className="w-4 h-4" /> : <IconX className="w-4 h-4" />}
                  At least one lowercase character.
                </span>
                <span className={`flex gap-2 items-center ${hasDigit ? "text-green-600" : "text-muted-foreground"}`}>
                  {hasDigit ? <IconCheck className="w-4 h-4" /> : <IconX className="w-4 h-4" />}
                  At least one digit.
                </span>
                <span className={`flex gap-2 items-center ${hasAllowedChars ? "text-green-600" : "text-muted-foreground"}`}>
                  {hasAllowedChars ? (<IconCheck className="w-4 h-4" />) : (<IconX className="w-4 h-4" />)}
                  At least one special characters ("." | "_" | "@").
                </span>
              </div>

            </section>

            <AlertDialogFooter>
              <AlertDialogCancel className="cursor-pointer">Cancel</AlertDialogCancel>
              <Button
                type="submit"
                className="cursor-pointer"
                disabled={!form.formState.isDirty || !form.formState.isValid || isPending}
              >
                {isPending && <Loader2Icon className="w-4 animate-spin" />}
                Save
              </Button>
            </AlertDialogFooter>
          </form>
        </Form>

      </AlertDialogContent>
    </AlertDialog>
  )
}

import {
  Sheet,
  SheetContent,
  SheetDescription, SheetFooter,
  SheetHeader,
  SheetTitle,
} from "@/components/ui/sheet"
import {useActionState} from "@/states/action-state";
import {useEffect, useState} from "react";
import {Button} from "@/components/ui/button";
import {Input} from "@/components/ui/input";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select";
import {useUserState} from "@/states/user-state";
import {z} from "zod";
import {useForm} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import {Form, FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/components/ui/form";
import {useUpdateUser} from "@/hooks/use-user";
import {Loader2Icon} from "lucide-react";
import {toast} from "sonner";
import {isHTTPResponse} from "@/lib/api";
import {useQueryClient} from "@tanstack/react-query";
import {TeamConfirmDeleteAlertDialog} from "@/features/stores/components/team/team-confirm-delete";
import type {User} from "@/types/user";
import {Tabs, TabsContent, TabsList, TabsTrigger} from "@/components/ui/tabs";
import {PasswordField} from "@/components/ui/password-input";
import {IconCheck, IconX} from "@tabler/icons-react";

export const TeamDetailActionSheetState = "team_detail_action_sheet_state"

const FormSchema = z.object({
  username: z.string(),
  name: z.string({
    required_error: "Please set a name to display.",
  }).min(1, "Name must be at least 1 character.").max(20, "Name must be at most 20 characters."),
  email: z.string().email(),
  role: z.string(),
  password: z.union([
    z.string()
      .min(8, { message: "Minimum 8 characters." })
      .max(12, { message: "Maximum 12 characters." })
      .regex(/(?=.*[A-Z])/, { message: "At least one uppercase character." })
      .regex(/(?=.*[a-z])/, { message: "At least one lowercase character." })
      .regex(/(?=.*\d)/, { message: "At least one digit." })
      .regex(/^[A-Za-z0-9._@]+$/, { message: "Only letters, numbers, '.', '_' and '@' are allowed." }),
    z.literal("")
  ]).optional(),
  confirm_password: z.union([z.string(), z.literal("")]).optional(),
}).superRefine(({ password, confirm_password }, ctx) => {
  // only validate if password is set and not empty
  if (password && password !== "" && password !== confirm_password) {
    ctx.addIssue({
      code: "custom",
      path: ["confirm_password"],
      message: "Passwords do not match.",
    });
  }
});

export function TeamDetailActionSheet() {
  const [activeTab, setActiveTab] = useState<"profile" | "password">("profile");
  const [open, setOpen] = useState(false);
  const { bool, setBoolState } = useActionState();
  const { roles, selectedUser } =  useUserState();
  const { mutate: update, isPending} = useUpdateUser();
  const queryClient = useQueryClient();

  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
    defaultValues: {
      username: "",
      name: "",
      email: "",
      role: "",
      password: "",
      confirm_password: ""
    }
  })

  useEffect(() => {
    if (bool[TeamDetailActionSheetState]) setOpen(bool[TeamDetailActionSheetState])

    if (selectedUser) reset(selectedUser)
  }, [bool, selectedUser]);

  const handleTabChange = (tab: string) => {
    if (selectedUser) reset(selectedUser);
    setActiveTab(tab as "profile" | "password");
  }

  const reset = (selectedUser: User) => {
    const id = setTimeout(() => {
      form.reset({
        username: selectedUser.username,
        name: selectedUser.name,
        email: selectedUser.email,
        role: `${selectedUser.role?.id}`,
        password: "",
        confirm_password: "",
      });
    }, 100);
    return () => clearTimeout(id);
  }

  const hasUppercase = /[A-Z]/.test(form.watch("password") || "");
  const hasLowercase = /[a-z]/.test(form.watch("password") || "");
  const hasDigit = /\d/.test(form.watch("password") || "");
  const hasAllowedChars = /[._@]/.test(form.watch("password") || "");

  const watchAll = form.watch(); // watch all fields

  const isChanged =
    watchAll.password && watchAll.password?.length > 0
      && watchAll.confirm_password && watchAll.confirm_password?.length > 0
      ? true : (
        watchAll.name !== selectedUser?.name ||
        watchAll.username !== selectedUser?.username ||
        watchAll.email !== selectedUser?.email ||
        Number(watchAll.role) !== selectedUser?.role?.id
      );

  function onSubmit(data: z.infer<typeof FormSchema>) {
    if (!selectedUser) return;

    let filteredValues = Object.fromEntries(
      Object.entries(data).filter(([_, value]) =>
        value !== undefined && value !== null && value !== ""
      )
    ) as Record<string, unknown>;

    // convert role (string) → role_id (number)
    if (filteredValues.role) {
      filteredValues = {
        ...filteredValues,
        role_id: Number(filteredValues.role),
      };
      delete filteredValues.role;
    }

    if (filteredValues.confirm_password) {
      delete filteredValues.confirm_password;
    }

    update({
      id: selectedUser.id,
      body: JSON.stringify(filteredValues)
    }, {
      onSuccess: async () => {
        await queryClient.invalidateQueries({ queryKey: ['users'] })
        toast.success("Update user successfully");
        onOpenChange(false);
      },
      onError: async (error) => {
        if (error && isHTTPResponse<null>(error)) {
          reset(selectedUser as User)
          toast.error(error.data);
          return;
        }
        if (error instanceof Error) {
          reset(selectedUser as User)
          const clientError = error as Error
          toast.error(clientError.message);
        }
      }
    })
  }

  function onOpenChange(newOpen: boolean) {
    setOpen(newOpen);
    if (!newOpen) {
      setBoolState(TeamDetailActionSheetState, false);
    }
  }

  return (
    <Sheet open={open} onOpenChange={onOpenChange}>
      <Form {...form}>
          <SheetContent>
            <SheetHeader>
              <SheetTitle>Team Member Info</SheetTitle>
              <SheetDescription>
                Preview team member details or remove them from the team.
              </SheetDescription>
            </SheetHeader>

            <form id="form-update" onSubmit={form.handleSubmit(onSubmit)}>
              <Tabs
                value={activeTab}
                className="h-full w-full px-4"
                onValueChange={handleTabChange}
              >
                <TabsList>
                  <TabsTrigger value="profile" className="cursor-pointer">
                    Profile
                  </TabsTrigger>
                  <TabsTrigger value="password" className="cursor-pointer">
                    Password
                  </TabsTrigger>
                </TabsList>

                <TabsContent value="profile" className="px-2">
                  <section className="grid gap-4 pt-2">
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
                          <Select onValueChange={field.onChange} value={field.value}>
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
                  </section>
                </TabsContent>

                <TabsContent value="password" className="px-2 pt-2 gap-4 flex flex-col">
                  <PasswordField
                    label="New password"
                    placeholder="Enter your current password"
                    {...form.register("password")}
                  />

                  <PasswordField
                    label="Confirm password"
                    placeholder="Enter your new password"
                    {...form.register("confirm_password")}
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
                </TabsContent>
              </Tabs>
            </form>

            <SheetFooter className="flex flex-row items-center justify-between">
              <TeamConfirmDeleteAlertDialog action={onOpenChange} />
              <div className="flex justify-end items-center space-x-2">
                <Button
                  variant="outline"
                  type="button"
                  className="cursor-pointer"
                  onClick={() => onOpenChange(false)}
                >Cancel</Button>
                <Button
                  type="submit"
                  className="cursor-pointer"
                  form="form-update"
                  disabled={!isChanged || isPending}
                >
                  {isPending && <Loader2Icon className="w-4 animate-spin" />}
                  Save
                </Button>
              </div>
            </SheetFooter>
          </SheetContent>
      </Form>
    </Sheet>
  )
}
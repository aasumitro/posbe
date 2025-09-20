import {
  Sheet,
  SheetContent,
  SheetDescription,
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

export const TeamDetailActionSheetState = "team_detail_action_sheet_state"

const FormSchema = z.object({
  username: z.string(),
  name: z.string({
    required_error: "Please set a name to display.",
  }).min(1, "Name must be at least 1 character.").max(20, "Name must be at most 20 characters."),
  email: z.string().email(),
  role: z.string(),
})

export function TeamDetailActionSheet() {
  const [open, setOpen] = useState(false);
  const { bool, setBoolState } = useActionState();
  const { roles, selectedUser } =  useUserState();
  const { mutate: update, isPending} = useUpdateUser(selectedUser?.id);
  const queryClient = useQueryClient();

  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
    defaultValues: {
      username: "",
      name: "",
      email: "",
      role: "",
    }
  })

  useEffect(() => {
    if (bool[TeamDetailActionSheetState]){
      setOpen(bool[TeamDetailActionSheetState])
    }

    if (selectedUser) {
      reset(selectedUser)
    }
  }, [bool, selectedUser]);

  const reset = (selectedUser: User) => {
    const id = setTimeout(() => {
      form.reset({
        username: selectedUser.username,
        name: selectedUser.name,
        email: selectedUser.email,
        role: `${selectedUser.role?.id}`
      });
    }, 100);
    return () => clearTimeout(id);
  }

  function onSubmit(data: z.infer<typeof FormSchema>) {
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

    update(JSON.stringify(filteredValues), {
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
      <SheetContent>
        <SheetHeader>
          <SheetTitle>Team Member Info</SheetTitle>
          <SheetDescription>
            Preview team member details or remove them from the team.
          </SheetDescription>
        </SheetHeader>

        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)}>
            <section className="grid gap-4 px-4 pt-2 pb-8">
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

            <div className="flex justify-end items-center px-4 mt-2 space-x-2">
              <Button
                variant="outline"
                type="button"
                onClick={() => onOpenChange(false)}
              >Cancel</Button>
              <Button
                type="submit"
                disabled={!form.formState.isDirty || !form.formState.isValid || isPending}
              >
                {isPending && <Loader2Icon className="w-4 animate-spin" />}
                Save
              </Button>
            </div>
          </form>
        </Form>

        {/*TODO: fix the style*/}
        <div className="absolute top-[39%]">
          <TeamConfirmDeleteAlertDialog action={onOpenChange} />
        </div>
      </SheetContent>
    </Sheet>
  )
}
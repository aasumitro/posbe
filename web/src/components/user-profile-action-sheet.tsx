import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
} from "@/components/ui/sheet"
import {useEffect, useState} from "react";
import { useActionState } from "@/states/action-state";
import {z} from "zod";
import {useForm} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import {Form, FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/components/ui/form";
import {Button} from "@/components/ui/button";
import {Input} from "@/components/ui/input";
import {useAuthStore} from "@/states/auth-state";
import {useUpdateProfile} from "@/hooks/use-user";
import {Loader2Icon} from "lucide-react";
import {toast} from "sonner";
import {isHTTPResponse} from "@/lib/api";
import {useQueryClient} from "@tanstack/react-query";

export const ProfileSheetState = "profile_sheet_state"

const UserProfileFormSchema = z.object({
  username: z.string(),
  name: z.string(),
  email: z.string(),
  role: z.string(),
});

export function UserProfileActionSheet() {
  const [openProfileSheet, setProfileSheetOpen] = useState(false);
  const { bool, setBoolState } = useActionState();
  const {auth} = useAuthStore();
  const { mutate: update, isPending} = useUpdateProfile();
  const queryClient = useQueryClient();

  useEffect(() => {
    if (bool[ProfileSheetState]){
      setProfileSheetOpen(bool[ProfileSheetState])
    }
  }, [bool]);

  const getResetValues = (data: typeof auth.user) => ({
    username: data?.username ?? "",
    name: data?.name ?? "",
    email: data?.email ?? "",
    role: data?.role?.name ?? "",
  });

  const form = useForm<z.infer<typeof UserProfileFormSchema>>({
    resolver: zodResolver(UserProfileFormSchema),
    defaultValues: getResetValues(auth.user),
    mode: "onChange"
  });

  useEffect(() => {
    if (!auth?.user) return;
    const id = setTimeout(() => {
      form.reset(getResetValues(auth?.user));
    }, 100);
    return () => clearTimeout(id);
  }, [auth?.user]);

  function onSubmit(data: z.infer<typeof UserProfileFormSchema>) {
    const filteredValues = Object.fromEntries(
      // eslint-disable-next-line @typescript-eslint/no-unused-vars
      Object.entries(data).filter(([_, value]) =>
        value !== undefined && value !== null && value !== "")
    );

    update(JSON.stringify(filteredValues), {
      onSuccess: async (resp) => {
        await queryClient.invalidateQueries({ queryKey: ['user', resp?.data?.id] })
        toast.success("Update profile successfully");
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

  function onOpenChange(newOpen: boolean) {
    setProfileSheetOpen(newOpen);
    if (!newOpen) {
      form.reset();
      setBoolState(ProfileSheetState, false);
    }
  }

  return (
    <Sheet open={openProfileSheet} onOpenChange={onOpenChange}>
      <SheetContent>
        <SheetHeader>
          <SheetTitle>Edit Profile</SheetTitle>
          <SheetDescription>
            Update your account information below.
          </SheetDescription>
        </SheetHeader>

        <Form {...form}>
          <form
            onSubmit={form.handleSubmit(onSubmit)}
            className="px-4 gap-4 flex flex-col"
          >
            <FormField
              control={form.control}
              name="username"
              render={({field}) => (
                <FormItem>
                  <FormLabel>Username</FormLabel>
                  <FormControl>
                    <Input
                      {...field}
                      placeholder="@zeros"
                    />
                  </FormControl>
                  <FormMessage/>
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="name"
              render={({field}) => (
                <FormItem>
                  <FormLabel>Name</FormLabel>
                  <FormControl>
                    <Input
                      {...field}
                      placeholder="Zeros Mardigu"
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
                  <FormLabel>Email address</FormLabel>
                  <FormControl>
                    <Input
                      {...field}
                      placeholder="zeros@member.pos"
                    />
                  </FormControl>
                  <FormMessage/>
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="role"
              render={({field}) => (
                <FormItem>
                  <FormLabel>Role</FormLabel>
                  <FormControl>
                    <Input
                      {...field}
                      placeholder="-"
                      disabled
                    />
                  </FormControl>
                  <FormMessage/>
                </FormItem>
              )}
            />

            <Button
              type="submit"
              className="ml-auto mt-2"
              disabled={!form.formState.isDirty || !form.formState.isValid || isPending}
            >
              {isPending && <Loader2Icon className="w-4 animate-spin" />}
              Save
            </Button>
          </form>
        </Form>
      </SheetContent>
    </Sheet>
  )
}
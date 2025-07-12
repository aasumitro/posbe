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
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select";

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

  useEffect(() => {
    if (bool[ProfileSheetState]){
      setProfileSheetOpen(bool[ProfileSheetState])
    }
  }, [bool]);

  const getResetValues = () => ({
    username: "",
    name: "",
    email: "",
    role: "",
  });

  const form = useForm<z.infer<typeof UserProfileFormSchema>>({
    resolver: zodResolver(UserProfileFormSchema),
    defaultValues: getResetValues(),
    mode: "onChange"
  });

  function onSubmit(data: z.infer<typeof UserProfileFormSchema>) {
    console.log(data);
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
                  <Select
                    onValueChange={field.onChange}
                  >
                    <FormControl>
                      <SelectTrigger className="w-full">
                        <SelectValue placeholder="Select a role"/>
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      {[
                        {id: 1, name: "admin"},
                        {id: 2, name: "cashier"},
                        {id: 3, name: "waiter"},
                      ].map((role) => (
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

            <Button
              type="submit"
              className="ml-auto mt-2"
              disabled={!form.formState.isDirty || !form.formState.isValid}
            >Save</Button>
          </form>
        </Form>
      </SheetContent>
    </Sheet>
  )
}
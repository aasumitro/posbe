"use client"

import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "@/components/ui/sheet"
import {useEffect, useState} from "react";
import {useGlobalStateStore} from "@/stores/global-state.ts";
import {KEYS} from "@/lib/keys.ts";
import {useStorePageState} from "@/stores/store-state.ts";
import {z} from "zod";
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
import {useMutation, useQuery, useQueryClient} from "react-query";
import {HttpResponse} from "@/lib/types/http-response.ts";
import {Role, User as UType} from "@/lib/types/user.ts";
import {useStoreUser} from "@/hooks/use-store-user.tsx";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select.tsx";
import {cn} from "@/lib/utils.ts";
import {EyeIcon, Loader2Icon, UserPen} from "lucide-react";
import {Toggle} from "@/components/ui/toggle.tsx";
import {useSessionStateStore} from "@/stores/session-state.ts";
import {ToastMessageContainer} from "@/components/toast-message-container.tsx";
import {HTTPStatusCode, isHttpResponse} from "@/lib/api.ts";

const userSchema = z.object({
  name:  z.string().min(5).max(16),
  username:  z.string().min(5).max(10),
  email:  z.string().email(),
  phone:  z.string().refine((value) => /^[+]{1}(?:[0-9-()/.]\s?){6,15}[0-9]{1}$/.test(value), {
    message: "Invalid phone number"
  }),
  role: z.string()
});

export const UserDetailSheet = ()=> {
  const form = useForm<z.infer<typeof userSchema>>({
    resolver: zodResolver(userSchema),
    defaultValues: {
      name: "",
      username: "",
      email: "",
      phone: "",
      role: "0"
    },
  });

  const [open, setOpen] = useState(false)
  const {boolStates, setBoolState} = useGlobalStateStore();
  const {user: access, setUserData} = useSessionStateStore();
  const {user, users, setUserDetail, setUserList, roles, setUserRoleList} = useStorePageState();
  const {editUser, getRoles, getUsers} = useStoreUser();
  const [editMode, setEditMode] = useState(false)
  const queryClient = useQueryClient()

  useQuery<HttpResponse<Role[]>>(
    "user.roles", getRoles, {
      onSuccess: (resp) => setUserRoleList(resp.data),
      onError: (error) => console.log(error),
      retry: false,
    });

  useQuery<HttpResponse<UType[]>>(
    "store.users", getUsers, {
      onSuccess: (resp) => setUserList(resp.data),
      onError: (error) => console.log(error),
      retry: false,
    });


  const updateUserAccount = useMutation(editUser, {
    onSuccess: async (resp) => {
      if (isHttpResponse(resp) && resp.code != HTTPStatusCode.OK) {
        ToastMessageContainer("User Updated", resp.data);
        return;
      }
      await queryClient.invalidateQueries('store.users')
      ToastMessageContainer("User Updated", "The request to update the account has been successfully processed and completed.");
      if (access?.id == user?.id) {
        localStorage.setItem(KEYS.UI_USER_PROFILE,
          JSON.stringify(resp.data));
        setUserData(resp.data);
      }
      onClose();
    },
  });

  useEffect(() => {
    if (boolStates[KEYS.DISPLAY_USER_DETAIL_SHEET]) {
      setOpen(true);
    }
    if (user) {
      form.setValue("name", user.name)
      form.setValue("username", user.username)
      form.setValue("email", user.email)
      form.setValue("phone", user.phone)
      form.setValue("role", `${user.role.id}`)
    }
  }, [boolStates, user])

  const onClose = () => {
    setOpen(!open);
    setUserDetail(null);
    setEditMode(false);
    setBoolState(KEYS.DISPLAY_USER_DETAIL_SHEET, false);
  }

  const onSubmit = (values: z.infer<typeof userSchema>) => {
    let self = false
    if (access?.id == user?.id) {
      self = true
    }

    if (user?.id == null) {
      ToastMessageContainer("Update User", "User id is required!");
      onClose();
      return
    }

    let body: Record<string, any> = {}

    if (values.name && values.name !== "") {
      body.name = values.name
    }

    if (values.username && values.username !== "") {
      body.username = values.username
    }

    if (values.email && values.email !== "") {
      body.email = values.email
    }

    if (values.phone && values.phone !== "") {
      body.phone = values.phone
    }

    if (values.role && (values.role !== "" && values.role !== "0")) {
      // bug in another page (when direct access profile), when users not fetched
      if (users && users?.filter((user) => user.role.name === "admin")
        .length === 1 && user?.role.name === "admin" && Number(values.role) !== user?.role.id) {
        ToastMessageContainer("Action Not Allowed",
          "Sorry, you can't change this user role, since admin account only one.");
        return;
      }

      body.role_id = Number(values.role)
    }

    updateUserAccount.mutate({self, userId: user?.id, body: JSON.stringify(body)})
  }

  const onModeChange = () => {
    if (access?.role.name !== "admin" && access?.id !== user?.id) {
      ToastMessageContainer("Not Authorized",
        "Sorry, you don't have authorization to access this menu!");
      return;
    }
    setEditMode(!editMode)
  }

  const ableToUpdateUsername = () =>  access?.id === user?.id &&
    user?.role.name.toLocaleLowerCase() !== "admin"

  return (
    <Sheet open={open} onOpenChange={onClose}>
      <SheetTrigger asChild></SheetTrigger>
      <SheetContent side="right">
        <SheetHeader>
          <SheetTitle>
            {editMode ? "Edit" : ""} User profile
            <Toggle
              className="mt-4 ml-4"
              size="sm"
              onClick={onModeChange}
            >
              {!editMode ? <>
                <EyeIcon className="mr-2 h-4 w-4"/>
                View Mode
              </> : <>
                <UserPen className="mr-2 h-4 w-4"/>
                Edit Mode
              </>}
            </Toggle>
          </SheetTitle>
          <SheetDescription>
            {editMode
              ? "Make changes to your profile here. Click save when you're done."
              : "Current selected user data (credentials & contact)."
            }
          </SheetDescription>
        </SheetHeader>

        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)}>
            <div className="space-y-2 py-4">
              <FormField
                control={form.control}
                name="name"
                render={({field}) => (
                  <FormItem>
                    <FormLabel>Name</FormLabel>
                    <FormControl>
                      <Input
                        {...field}
                        placeholder="Lorem Ipsum"
                        disabled={!editMode}
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
                    <FormLabel>Username</FormLabel>
                    <FormControl>
                      <Input
                        {...field}
                        placeholder="@lorem"
                        disabled={!editMode || ableToUpdateUsername()}
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
                        disabled={!editMode}
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
                        disabled={!editMode}
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
                      defaultValue={`${user?.role.id}`}
                      disabled={!editMode || access?.role.name !== "admin"}
                    >
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
            </div>

            <div className={cn(
              {"hidden": !editMode},
              "w-full text-right"
            )}>
              <Button type="submit" disabled={updateUserAccount.isLoading}>
                {updateUserAccount.isLoading && <Loader2Icon className="w-4 h-4 animate-spin mr-2" />}
                Save changes
              </Button>
            </div>
          </form>
        </Form>
      </SheetContent>
    </Sheet>
  )
}

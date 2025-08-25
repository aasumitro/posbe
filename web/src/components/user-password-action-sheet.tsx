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
import {Form} from "@/components/ui/form";
import {PasswordField} from "@/components/ui/password-input";
import {Button} from "@/components/ui/button";
import {IconCheck, IconX} from "@tabler/icons-react";
import {useAuthStore} from "@/states/auth-state";
import {useUpdatePassword} from "@/hooks/use-user";
import {toast} from "sonner";
import {isHTTPResponse} from "@/lib/api";
import {Loader2Icon} from "lucide-react";
import {useNavigate} from "@tanstack/react-router";

export const PasswordSheetState = "password_sheet_state"

type PasswordErrorResponse = {
  new_password?: string[]
}

const UserPasswordFormSchema = z.object({
  old_pwd: z.string().min(6).max(12),
  new_pwd: z
    .string({ required_error: "Password cannot be empty." })
    .min(8, { message: "Minimum 8 characters." })
    .max(12, { message: "Maximum 12 characters." })
    .regex(/(?=.*[A-Z])/, { message: "At least one uppercase character." })
    .regex(/(?=.*[a-z])/, { message: "At least one lowercase character." })
    .regex(/(?=.*\d)/, { message: "At least one digit." })
    .regex(/^[A-Za-z0-9._@]+$/, { message: "Only letters, numbers, '.', '_' and '@' are allowed." }),
});

export function UserPasswordActionSheet() {
  const [open, setOpen] = useState(false);
  const { bool, setBoolState } = useActionState();
  const {auth} = useAuthStore();
  const { mutate: updatePassword, isPending} = useUpdatePassword();
  const navigate = useNavigate();

  useEffect(() => {
    if (bool[PasswordSheetState]){
      setOpen(bool[PasswordSheetState])
    }
  }, [bool]);

  const getResetValues = () => ({
    old_pwd: "",
    new_pwd: "",
  });

  // Initialize form with live default values
  const form = useForm<z.infer<typeof UserPasswordFormSchema>>({
    resolver: zodResolver(UserPasswordFormSchema),
    defaultValues: getResetValues(),
  });

  const hasUppercase = /[A-Z]/.test(form.watch("new_pwd") || "");
  const hasLowercase = /[a-z]/.test(form.watch("new_pwd") || "");
  const hasDigit = /\d/.test(form.watch("new_pwd") || "");
  const hasAllowedChars = /[._@]/.test(form.watch("new_pwd"));
  const oldPwd = form.watch("old_pwd") || "";
  const isOldPwdValid = oldPwd.length >= 6 && oldPwd.length <= 12;
  const isFormValid = isOldPwdValid && hasUppercase && hasLowercase && hasDigit && hasAllowedChars;

  function onSubmit(data: z.infer<typeof UserPasswordFormSchema>) {
    updatePassword(JSON.stringify({
      password: data.old_pwd,
      new_password: data.new_pwd,
    }), {
      onError: async (error) => {
        if (error && isHTTPResponse<null>(error)) {
          if (typeof error.data === "string") {
            if (error.data === "Password updated successfully. Please log in again.") {
              // it should not return success inside an error branch.
              // may need to adjust the API or handle it in onSuccess.
              // but it's ok for now
              toast.success("Password update successfully, please re-login to continue!");
              auth.reset();
              onOpenChange(false);
              await navigate({to: "/login", replace: true});
              return;
            }

            if (error.data === "invalid username or password") {
              toast.error("The current password you entered is incorrect.");
              return;
            }
          }

          if (typeof error.data === "object" && error.data !== null) {
            const data = error.data as PasswordErrorResponse;
            if (data.new_password && data.new_password.length > 0) {
              toast.error(data.new_password[0]);
            }
          }

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
    setOpen(newOpen);
    if (!newOpen) {
      setBoolState(PasswordSheetState, false);
      form.reset();
    }
  }

  return (
    <Sheet open={open} onOpenChange={onOpenChange}>
      <SheetContent>
        <SheetHeader>
          <SheetTitle>
            Change Password
          </SheetTitle>
          <SheetDescription>
            Enter your current password and choose a new one that meets the security requirements.
          </SheetDescription>
        </SheetHeader>
        <Form {...form}>
          <form
            onSubmit={form.handleSubmit(onSubmit)}
            className="px-4 gap-4 flex flex-col"
          >
            <PasswordField
              label="Current password"
              placeholder="Enter your current password"
              {...form.register("old_pwd")}
            />

            <PasswordField
              label="New password"
              placeholder="Enter your new password"
              {...form.register("new_pwd")}
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

            <Button
              type="submit"
              className="ml-auto"
              disabled={!isFormValid || isPending}
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
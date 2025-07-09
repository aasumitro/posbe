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

export const PasswordSheetState = "password_sheet_state"

const UserPasswordFormSchema = z.object({
  old_pwd: z.string().min(6).max(12),
  new_pwd: z
    .string({ required_error: "Password cannot be empty." })
    .min(6, { message: "Minimum 6 characters." })
    .max(12, { message: "Maximum 12 characters." })
    .regex(/(?=.*[A-Z])/, { message: "At least one uppercase character." })
    .regex(/(?=.*[a-z])/, { message: "At least one lowercase character." })
    .regex(/(?=.*\d)/, { message: "At least one digit." }),
});

export function UserPasswordActionSheet() {
  const [open, setOpen] = useState(false);
  const { bool, setBoolState } = useActionState();

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
  const oldPwd = form.watch("old_pwd") || "";
  const isOldPwdValid = oldPwd.length >= 6 && oldPwd.length <= 12;
  const isFormValid = isOldPwdValid && hasUppercase && hasLowercase && hasDigit;

  function onSubmit(data: z.infer<typeof UserPasswordFormSchema>) {
    console.log(data);
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
            </div>

            <Button type="submit" className="ml-auto" disabled={!isFormValid}>Save</Button>
          </form>
        </Form>
      </SheetContent>
    </Sheet>
  )
}
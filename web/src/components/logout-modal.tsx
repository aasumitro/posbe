import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog"
import {useEffect, useState} from "react";
import {useGlobalStateStore} from "@/stores/global-state.ts";
import {Loader2Icon} from "lucide-react";
import {KEYS} from "@/lib/keys.ts";
import {useNavigate} from "react-router-dom";
import {useMutation} from "react-query";
import {useAuth} from "@/hooks/use-auth.tsx";
import {ToastMessageContainer} from "@/components/toast-message-container.tsx";
import {isHttpResponse} from "@/lib/api.ts";

export function LogoutDialog() {
  const [open, setOpen] = useState(false)
  const {boolStates, setBoolState} = useGlobalStateStore();
  const navigate = useNavigate();
  const {logout} = useAuth();

  useEffect(() => {
    if (boolStates[KEYS.DISPLAY_LOGOUT_MODAL]) {
      setOpen(true);
    }
  }, [boolStates])

  const onClose = () => {
    setOpen(!open);
    setBoolState(KEYS.DISPLAY_LOGOUT_MODAL, false);
  }

  const signOut = useMutation(logout, {
    onSuccess: (resp) => {
      if (isHttpResponse(resp) && resp.code != 200) {
        ToastMessageContainer("Something went wrong", resp);
        return;
      }
      localStorage.removeItem(KEYS.UI_AUTH_TOKEN);
      localStorage.removeItem(KEYS.UI_USER_PROFILE);
      navigate("/login");
    },
  });

  const fireLogout = () =>  signOut.mutate();

  return (
    <AlertDialog open={open} onOpenChange={onClose}>
      <AlertDialogTrigger asChild></AlertDialogTrigger>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Sign Out Confirmation</AlertDialogTitle>
          <AlertDialogDescription>
            Are you sure you want to log out? This action will end your current session and you will need to log in again to access your account. Your data will remain intact.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel>Cancel</AlertDialogCancel>
          <AlertDialogAction
            className="bg-red-500 hover:bg-red-600 dark:text-white"
            onClick={fireLogout}
            disabled={signOut.isLoading}
          >
            {signOut.isLoading && <Loader2Icon className="w-4 h-4 animate-spin mr-2" />}
            Log Out
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  )
}
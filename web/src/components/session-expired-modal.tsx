import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog"
import {useEffect, useState} from "react";
import {useGlobalStateStore} from "@/stores/global-state.ts";
import {KEYS} from "@/lib/keys.ts";
import {useNavigate} from "react-router-dom";

export function SessionExpiredModal() {
  const [open, setOpen] = useState(false)
  const {boolStates, setBoolState} = useGlobalStateStore();
  const navigate = useNavigate();

  useEffect(() => {
    if (boolStates[KEYS.DISPLAY_SESSION_EXPIRED_MODAL]) {
      setOpen(true);
    }
  }, [boolStates])

  const onClose = () => {
    setOpen(!open);
    setBoolState(KEYS.DISPLAY_SESSION_EXPIRED_MODAL, false);
    localStorage.removeItem(KEYS.UI_AUTH_TOKEN);
    localStorage.removeItem(KEYS.UI_USER_PROFILE);
    navigate("/login");
  }

  return (
    <AlertDialog open={open}>
      <AlertDialogTrigger asChild></AlertDialogTrigger>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Session Expired</AlertDialogTitle>
          <AlertDialogDescription>
            Your session has expired. Please log in again to continue.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogAction
            onClick={onClose}
          >OK</AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  )
}
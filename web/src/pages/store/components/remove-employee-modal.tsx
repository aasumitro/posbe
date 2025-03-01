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
import {useStorePageState} from "@/stores/store-state.ts";
import {useStoreUser} from "@/hooks/use-store-user.tsx";
import {useMutation, useQueryClient} from "react-query";
import {HTTPStatusCode, isHttpResponse} from "@/lib/api.ts";
import {ToastMessageContainer} from "@/components/toast-message-container.tsx";

export function RemoveEmployeeModal() {
  const [open, setOpen] = useState(false)
  const {boolStates, setBoolState} = useGlobalStateStore();
  const {user} = useStorePageState();
  const {deleteUser} = useStoreUser();
  const queryClient = useQueryClient()

  useEffect(() => {
    if (boolStates[KEYS.DISPLAY_REMOVE_EMPLOYEE_CONFIRMATION_MODAL]) {
      setOpen(true);
    }
  }, [boolStates])

  const onClose = () => {
    setOpen(!open);
    setBoolState(KEYS.DISPLAY_REMOVE_EMPLOYEE_CONFIRMATION_MODAL, false);
  }

  const destroyUserAccount = useMutation(deleteUser, {
    onSuccess: async (resp) => {
      if (isHttpResponse(resp) && resp.code != HTTPStatusCode.NoContent) {
        ToastMessageContainer("Something went wrong", resp);
        return;
      }
      await queryClient.invalidateQueries('store.users')
      ToastMessageContainer("User Removed", "The user has been removed from the record.");
    },
  });

  const fireDeleteAction = (userId: string) =>  destroyUserAccount.mutate(userId);

  return (
    <AlertDialog open={open} onOpenChange={onClose}>
      <AlertDialogTrigger asChild></AlertDialogTrigger>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Confirm User Deletion</AlertDialogTitle>
          <AlertDialogDescription>
            Are you sure you want to delete this account ({user?.name})? This action is irreversible and will permanently remove all user data associated with this account.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel>Cancel</AlertDialogCancel>
          <AlertDialogAction
            className="bg-red-500 hover:bg-red-600 dark:text-white"
            onClick={() => fireDeleteAction(`${user!.id}`)}
            disabled={destroyUserAccount.isLoading}
          >
            {destroyUserAccount.isLoading && <Loader2Icon className="w-4 h-4 animate-spin mr-2" />}
            Yes, Destroy!
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  )
}
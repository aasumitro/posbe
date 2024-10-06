import {Outlet,  useNavigate} from "react-router-dom";
import {Toolbar} from "@/components/toolbar";
import {useEffect} from "react";
import {KEYS} from "@/lib/keys.ts";
import {useSessionStateStore} from "@/stores/session-state.ts";
import {LogoutDialog} from "@/components/logout-modal.tsx";
import {SessionExpiredModal} from "@/components/session-expired-modal.tsx";
import {UserDetailSheet} from "@/components/user-detail-sheet.tsx";

export const BackofficeLayout = () => {
  const navigate = useNavigate();
  const {setUserToken, setUserData} = useSessionStateStore();

  useEffect(() => {
    const token = localStorage.getItem(KEYS.UI_AUTH_TOKEN)
    const profile = localStorage.getItem(KEYS.UI_USER_PROFILE)
    if (!token) {
      navigate("/login");
      return;
    }
    setUserToken(token);
    if (profile) setUserData(JSON.parse(profile));
  }, []);

  return (
    <div className="flex flex-col overflow-y-hidden bg-background">
      <Toolbar />
      <main className="w-full h-[calc(100vh-0px)] overflow-y-auto bg-muted/40">
        <Outlet/>
        <UserDetailSheet />
        <LogoutDialog />
        <SessionExpiredModal />
      </main>
    </div>
  )
}

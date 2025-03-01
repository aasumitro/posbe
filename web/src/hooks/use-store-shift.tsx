import {useSessionStateStore} from "@/stores/session-state.ts";
import {useGlobalStateStore} from "@/stores/global-state.ts";
import {HttpResponse} from "@/lib/types/http-response.ts";
import {Endpoint, HTTPStatusCode} from "@/lib/api.ts";
import {KEYS} from "@/lib/keys.ts";
import {Shift} from "@/lib/types/shift.ts";

export function useStoreShift() {
  const {token} = useSessionStateStore();
  const {setBoolState} = useGlobalStateStore();

  const get = async (): Promise<HttpResponse<Shift[]>> => {
    const response = await fetch(Endpoint.Shift.Base, {
      method: "GET",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
    });

    if (!response.ok && response.status === HTTPStatusCode.Unauthorized) {
      setBoolState(KEYS.DISPLAY_SESSION_EXPIRED_MODAL, true)
    }

    return await response.json();
  };

  return {get}
}
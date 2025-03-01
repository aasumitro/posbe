import {useSessionStateStore} from "@/stores/session-state.ts";
import {Endpoint, HTTPStatusCode} from "@/lib/api.ts";
import {HttpResponse} from "@/lib/types/http-response.ts";
import {Store} from "@/lib/types/store.ts";
import {useGlobalStateStore} from "@/stores/global-state.ts";
import {KEYS} from "@/lib/keys.ts";

export function useStoreSetting() {
  const {token} = useSessionStateStore();
  const {setBoolState} = useGlobalStateStore();

  const get = async (): Promise<HttpResponse<Store>> => {
    const response = await fetch(Endpoint.Store.Pref, {
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

  const edit = async (body: string): Promise<HttpResponse<Store>> => {
    const response = await fetch(Endpoint.Auth.SignOut, {
      method: "PUT",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body,
    });

    if (!response.ok && response.status === HTTPStatusCode.Unauthorized) {
      setBoolState(KEYS.DISPLAY_SESSION_EXPIRED_MODAL, true)
    }

    return await response.json();
  };

  return {get, edit}
}
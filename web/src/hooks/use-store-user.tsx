import {useSessionStateStore} from "@/stores/session-state.ts";
import {Endpoint, HTTPStatusCode} from "@/lib/api.ts";
import {HttpResponse} from "@/lib/types/http-response.ts";
import {Role, User} from "@/lib/types/user.ts";
import {KEYS} from "@/lib/keys.ts";
import {useGlobalStateStore} from "@/stores/global-state.ts";

export function useStoreUser() {
  const {token} = useSessionStateStore();
  const {setBoolState} = useGlobalStateStore();

  const getUsers = async (): Promise<HttpResponse<User[]>> => {
    const response = await fetch(Endpoint.User.Base, {
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

  const newUser = async (body: string): Promise<HttpResponse<User>> => {
    const response = await fetch(`${Endpoint.User.Base}`, {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body
    });

    if (!response.ok && response.status === HTTPStatusCode.Unauthorized) {
      setBoolState(KEYS.DISPLAY_SESSION_EXPIRED_MODAL, true)
    }

    return response.json();
  };

  const editUser = async ({
   self, userId, body,
  }:{
    self: boolean;
    userId: number;
    body: string;
  }): Promise<HttpResponse<User>> => {
    let uri = Endpoint.User.Base
    if (!self) uri += `/${userId}`

    const response = await fetch(uri, {
      method: "PUT",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body
    });

    if (!response.ok && response.status === HTTPStatusCode.Unauthorized) {
      setBoolState(KEYS.DISPLAY_SESSION_EXPIRED_MODAL, true)
    }

    return response.json();
  };

  const deleteUser = async (userId: string): Promise<HttpResponse<string>> => {
    const response = await fetch(`${Endpoint.User.Base}/${userId}`, {
      method: "DELETE",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
    });

    if (!response.ok && response.status === HTTPStatusCode.Unauthorized) {
      setBoolState(KEYS.DISPLAY_SESSION_EXPIRED_MODAL, true)
    }

    return { data: "" } as HttpResponse<string>;
  };

  const getRoles = async (): Promise<HttpResponse<Role[]>> => {
    const response = await fetch(Endpoint.User.Role, {
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

  return {getUsers, newUser, editUser, deleteUser, getRoles}
}
import {useSessionStateStore} from "@/stores/session-state.ts";
import {Endpoint} from "@/lib/api.ts";
import {HttpResponse} from "@/lib/types/http-response.ts";
import {Access} from "@/lib/types/user.ts";

export function useAuth() {
  const {token} = useSessionStateStore();

  const login = async (body: string): Promise<HttpResponse<Access>> => {
    const response = await fetch(Endpoint.Auth.SignIn, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
      body,
    });

    return await response.json();
  };

  const logout = async (): Promise<HttpResponse<string> | string> => {
    const response = await fetch(Endpoint.Auth.SignOut, {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
    });

    return await response.json();
  };

  return {login, logout}
}
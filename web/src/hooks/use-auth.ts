import type {HTTPResponse} from "@/types/http-response";
import {api, API_PATH, AuthPath, catchHTTPError} from "@/lib/api";
import {useMutation} from "@tanstack/react-query";
import type {AccountAccess} from "@/types/user";

export function useLogin() {
  const login = async (
    data: { username: string, password: string }
  ): Promise<HTTPResponse<AccountAccess>> => {
    try {
      const url = API_PATH.ACCOUNT.AUTH(AuthPath.SIGN_IN)
      const response =
        await api.post<HTTPResponse<AccountAccess>>(url, data);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useMutation({ mutationFn: login })
}

export function useLogout() {
  const logout = async (): Promise<HTTPResponse<null>> => {
    try {
      const url = API_PATH.ACCOUNT.AUTH(AuthPath.SIGN_OUT)
      const response =
        await api.post<HTTPResponse<null>>(url, null);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useMutation({ mutationFn: logout })
}
import type {HTTPResponse} from "@/types/http-response";
import {api, API_PATH, catchHTTPError} from "@/lib/api";
import {useMutation, useSuspenseQuery} from "@tanstack/react-query";
import type {StoreSetting} from "@/types/store";
import type {User} from "@/types/user";

export function useStoresSetting() {
  const settings = async (): Promise<HTTPResponse<StoreSetting>> => {
    try {
      const url = API_PATH.STORE.SETTINGS
      const response =
        await api.get<HTTPResponse<StoreSetting>>(url);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useSuspenseQuery({ queryKey: ['store.settings'], queryFn: settings })
}

export function useUpdateSetting() {
  const setting = async (
    body: string
  ): Promise<HTTPResponse<null>> => {
    try {
      const url = API_PATH.STORE.SETTINGS
      const response =
        await api.put<HTTPResponse<null>>(url, body);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useMutation({ mutationFn: profile })
}
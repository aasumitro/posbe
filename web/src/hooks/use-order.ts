import type {HTTPResponse} from "@/types/http-response";
import {api, API_PATH, catchHTTPError} from "@/lib/api";
import {useMutation, useSuspenseQuery} from "@tanstack/react-query";
import type {ActiveShift} from "@/types/shift";

export function useActiveShift() {
  const activeShift = async (): Promise<HTTPResponse<ActiveShift>> => {
    try {
      const url = API_PATH.ORDERS.SHIFTS
      const response =
        await api.get<HTTPResponse<ActiveShift>>(url);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useSuspenseQuery({ queryKey: ['active.shift'], queryFn: activeShift })
}

export function useShiftAction() {
  const newActiveShift = async (
    body: string
  ): Promise<HTTPResponse<ActiveShift>> => {
    try {
      const url = API_PATH.ORDERS.SHIFTS
      const response =
        await api.post<HTTPResponse<ActiveShift>>(url, body);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useMutation({ mutationFn: newActiveShift })
}

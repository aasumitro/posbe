import type {HTTPResponse} from "@/types/http-response";
import {api, API_PATH, catchHTTPError} from "@/lib/api";
import {useMutation, useQuery, useSuspenseQuery} from "@tanstack/react-query";
import type {StoreSetting} from "@/types/store";
import type {StoreShift} from "@/types/shift";

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
        await api.patch<HTTPResponse<null>>(url, body);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useMutation({ mutationFn: setting })
}

export function useStoreShifts() {
  const shifts = async (): Promise<HTTPResponse<StoreShift[]>> => {
    try {
      const url = API_PATH.STORE.SHIFTS
      const response =
        await api.get<HTTPResponse<StoreShift[]>>(url);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useSuspenseQuery({ queryKey: ['store.shifts'], queryFn: shifts })
}

export function useStoreShiftDetail(id?: number) {
  const shift = async (): Promise<HTTPResponse<StoreShift>> => {
    try {
      const url = `${API_PATH.STORE.SHIFTS}/${id}`
      const response =
        await api.get<HTTPResponse<StoreShift>>(url);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useQuery({ queryKey: ["store.shift", id], queryFn: shift,  enabled: !!id, })
}

export function useNewStoreShift() {
  const newShift = async (
    body: string
  ): Promise<HTTPResponse<StoreShift>> => {
    try {
      const url = API_PATH.STORE.SHIFTS
      const response =
        await api.post<HTTPResponse<StoreShift>>(url, body);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useMutation({ mutationFn: newShift })
}

export function useEditStoreShift() {
  const editShift = async (
    {id, body}: {id: number, body: string}
  ): Promise<HTTPResponse<StoreShift>> => {
    try {
      const url = `${API_PATH.STORE.SHIFTS}/${id}`
      const response =
        await api.patch<HTTPResponse<StoreShift>>(url, body);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useMutation({ mutationFn: editShift })
}

export function useDeleteShift() {
  const deleteShift = async (
    id?: number
  ): Promise<HTTPResponse<null>> => {
    try {
      const url = `${API_PATH.STORE.SHIFTS}/${id}`
      const response =
        await api.delete<HTTPResponse<null>>(url);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useMutation({ mutationFn: deleteShift })
}
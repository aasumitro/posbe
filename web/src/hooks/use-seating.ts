import type {HTTPResponse} from "@/types/http-response";
import {api, API_PATH, catchHTTPError} from "@/lib/api";
import {useMutation, useQuery, useSuspenseQuery} from "@tanstack/react-query";
import type {Floor, Table} from "@/types/seating";

export function useFloorList() {
  const floors = async (): Promise<HTTPResponse<Floor[]>> => {
    try {
      const url = API_PATH.STORE.FLOORS
      const response =
        await api.get<HTTPResponse<Floor[]>>(url);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useSuspenseQuery({ queryKey: ['floors'], queryFn: floors })
}

export function useFloorDetail(id?: number) {
  const floor = async (): Promise<HTTPResponse<Floor>> => {
    try {
      const url = `${API_PATH.STORE.FLOORS}/${id}`
      const response =
        await api.get<HTTPResponse<Floor>>(url);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useQuery({ queryKey: ["floor", id], queryFn: floor,  enabled: !!id, })
}

export function useFloorTableList(id?: number) {
  const tables = async (): Promise<HTTPResponse<Table[]>> => {
    try {
      const url = `${API_PATH.STORE.FLOORS}/${id}/tables`
      const response =
        await api.get<HTTPResponse<Table[]>>(url);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useQuery({ queryKey: ["floor.tables", id], queryFn: tables,  enabled: !!id, })
}

export function useNewFloor() {
  const newFloor = async (
    body: string
  ): Promise<HTTPResponse<Floor>> => {
    try {
      const url = API_PATH.STORE.FLOORS
      const response =
        await api.post<HTTPResponse<Floor>>(url, body);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useMutation({ mutationFn: newFloor })
}

export function useUpdateFloor(id?: number) {
  const editFloor = async (
    body: string
  ): Promise<HTTPResponse<Floor>> => {
    try {
      const url = `${API_PATH.STORE.FLOORS}/${id}`
      const response =
        await api.patch<HTTPResponse<Floor>>(url, body);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useMutation({ mutationFn: editFloor })
}

export function useDeleteFloor() {
  const deleteFloor = async (
    id?: number
  ): Promise<HTTPResponse<null>> => {
    try {
      const url = `${API_PATH.STORE.FLOORS}/${id}`
      const response =
        await api.delete<HTTPResponse<null>>(url);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useMutation({ mutationFn: deleteFloor })
}

export function useNewTable() {
  const newTable = async (
    body: string
  ): Promise<HTTPResponse<Table>> => {
    try {
      const url = API_PATH.STORE.TABLES
      const response =
        await api.post<HTTPResponse<Table>>(url, body);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useMutation({ mutationFn: newTable })
}

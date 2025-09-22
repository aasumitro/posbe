import type {HTTPResponse} from "@/types/http-response";
import {api, API_PATH, catchHTTPError} from "@/lib/api";
import {useMutation, useSuspenseQuery} from "@tanstack/react-query";
import type {Unit} from "@/types/unit";
import type {Category} from "@/types/category";

export function useUnitList() {
  const units = async (): Promise<HTTPResponse<Unit[]>> => {
    try {
      const url = API_PATH.CATALOG.ATTRIBUTES.UNITS
      const response =
        await api.get<HTTPResponse<Unit[]>>(url);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useSuspenseQuery({ queryKey: ['units'], queryFn: units })
}

export function useNewUnit() {
  const newUnit = async (
    body: string
  ): Promise<HTTPResponse<Unit>> => {
    try {
      const url = API_PATH.CATALOG.ATTRIBUTES.UNITS
      const response =
        await api.post<HTTPResponse<Unit>>(url, body);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useMutation({ mutationFn: newUnit })
}

export function useUpdateUnit() {
  const editUnit = async (
    {id, body}: {id?: number, body: string}
  ): Promise<HTTPResponse<Unit>> => {
    try {
      const url = `${API_PATH.CATALOG.ATTRIBUTES.UNITS}/${id}`
      const response =
        await api.patch<HTTPResponse<Unit>>(url, body);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useMutation({ mutationFn: editUnit })
}

export function useDeleteUnit() {
  const deleteUnit = async (
    id?: number
  ): Promise<HTTPResponse<null>> => {
    try {
      const url = `${ API_PATH.CATALOG.ATTRIBUTES.UNITS}/${id}`
      const response =
        await api.delete<HTTPResponse<null>>(url);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useMutation({ mutationFn: deleteUnit })
}

export function useCategoryList() {
  const categories = async (): Promise<HTTPResponse<Category[]>> => {
    try {
      const url = API_PATH.CATALOG.ATTRIBUTES.CATEGORIES
      const response =
        await api.get<HTTPResponse<Category[]>>(url);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useSuspenseQuery({ queryKey: ['categories'], queryFn: categories })
}

export function useNewCategory() {
  const newCategory = async (
    body: string
  ): Promise<HTTPResponse<Category>> => {
    try {
      const url = API_PATH.CATALOG.ATTRIBUTES.CATEGORIES
      const response =
        await api.post<HTTPResponse<Category>>(url, body);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useMutation({ mutationFn: newCategory })
}

export function useDeleteCategory() {
  const deleteCategory = async (
    id?: number
  ): Promise<HTTPResponse<null>> => {
    try {
      const url = `${ API_PATH.CATALOG.ATTRIBUTES.CATEGORIES}/${id}`
      const response =
        await api.delete<HTTPResponse<null>>(url);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useMutation({ mutationFn: deleteCategory })
}

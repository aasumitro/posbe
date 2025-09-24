import type {HTTPResponse} from "@/types/http-response";
import {api, API_PATH, catchHTTPError} from "@/lib/api";
import {useMutation, useSuspenseQuery} from "@tanstack/react-query";
import type {ProductAddon} from "@/types/product";

export function useProductAddonList() {
  const addons = async (): Promise<HTTPResponse<ProductAddon[]>> => {
    try {
      const url = API_PATH.CATALOG.PRODUCTS.ADDONS
      const response =
        await api.get<HTTPResponse<ProductAddon[]>>(url);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useSuspenseQuery({ queryKey: ['addons'], queryFn: addons })
}

export function useNewProductAddon() {
  const newAddon = async (
    body: string
  ): Promise<HTTPResponse<ProductAddon>> => {
    try {
      const url = API_PATH.CATALOG.PRODUCTS.ADDONS
      const response =
        await api.post<HTTPResponse<ProductAddon>>(url, body);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useMutation({ mutationFn: newAddon })
}

export function useUpdateProductAddon() {
  const editAddon = async (
    {id, body}: {id?: number, body: string}
  ): Promise<HTTPResponse<ProductAddon>> => {
    try {
      const url = `${API_PATH.CATALOG.PRODUCTS.ADDONS}/${id}`
      const response =
        await api.patch<HTTPResponse<ProductAddon>>(url, body);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useMutation({ mutationFn: editAddon })
}

export function useDeleteProductAddon() {
  const deleteAddon = async (
    id?: number
  ): Promise<HTTPResponse<null>> => {
    try {
      const url = `${ API_PATH.CATALOG.PRODUCTS.ADDONS}/${id}`
      const response =
        await api.delete<HTTPResponse<null>>(url);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useMutation({ mutationFn: deleteAddon })
}

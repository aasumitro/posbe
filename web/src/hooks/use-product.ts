import type {HTTPResponse} from "@/types/http-response";
import {api, API_PATH, catchHTTPError} from "@/lib/api";
import {useMutation, useQuery, useSuspenseQuery} from "@tanstack/react-query";
import type {Product, ProductAddon} from "@/types/product";

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

export function useProductList() {
  const products = async (): Promise<HTTPResponse<Product[]>> => {
    try {
      const url = API_PATH.CATALOG.PRODUCTS.BASE
      const response =
        await api.get<HTTPResponse<Product[]>>(url);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useSuspenseQuery({ queryKey: ['products'], queryFn: products })
}

export function useProductDetail(id?: number) {
  const product = async (): Promise<HTTPResponse<Product>> => {
    try {
      const url = `${API_PATH.CATALOG.PRODUCTS.BASE}/${id}`
      const response =
        await api.get<HTTPResponse<Product>>(url);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useQuery({ queryKey: ["product", id], queryFn: product,  enabled: !!id, })
}

export function useNewProduct() {
  const newProduct = async (
    body: string
  ): Promise<HTTPResponse<Product>> => {
    try {
      const url = API_PATH.CATALOG.PRODUCTS.BASE
      const response =
        await api.post<HTTPResponse<Product>>(url, body);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useMutation({ mutationFn: newProduct })
}

export function useUpdateProduct() {
  const editProduct = async (
    {id, body}: {id?: number, body: string}
  ): Promise<HTTPResponse<Product>> => {
    try {
      const url = `${API_PATH.CATALOG.PRODUCTS.BASE}/${id}`
      const response =
        await api.patch<HTTPResponse<Product>>(url, body);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useMutation({ mutationFn: editProduct })
}

export function useDeleteProduct() {
  const deleteProduct = async (
    id?: number
  ): Promise<HTTPResponse<null>> => {
    try {
      const url = `${ API_PATH.CATALOG.PRODUCTS.BASE}/${id}`
      const response =
        await api.delete<HTTPResponse<null>>(url);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useMutation({ mutationFn: deleteProduct })
}

export function useDeleteProductVariant() {
  const deleteVariant = async (
    {pid, vid}: {pid?: number, vid: number}
  ): Promise<HTTPResponse<null>> => {
    try {
      const url = `${ API_PATH.CATALOG.PRODUCTS.BASE}/${pid}/variants/${vid}`
      const response =
        await api.delete<HTTPResponse<null>>(url);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useMutation({ mutationFn: deleteVariant })
}

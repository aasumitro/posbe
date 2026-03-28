import type {HTTPResponse} from "@/types/http-response";
import type {Customer} from "@/types/customer";
import {api, API_PATH, catchHTTPError} from "@/lib/api";
import {useMutation, useSuspenseQuery} from "@tanstack/react-query";

export function useCustomerList() {
  const customer = async (): Promise<HTTPResponse<Customer[]>> => {
    try {
      const url = API_PATH.CUSTOMER.BASE;
      const response =
        await api.get<HTTPResponse<Customer[]>>(url);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error)
    }
  }

  return useSuspenseQuery(({ queryKey: ["customers"], queryFn: customer}))
}

export function useNewCustomer() {
  const newCustomer = async (
    body: string
  ): Promise<HTTPResponse<null>> => {
    try {
      const url = API_PATH.CUSTOMER.BASE
      const response =
        await api.post<HTTPResponse<null>>(url, body);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  }

  return useMutation({ mutationFn: newCustomer })
}

export function useUpdateCustomer() {
  const editCustomer = async (
    {id, body}: {id?: number, body: string}
  ): Promise<HTTPResponse<null>> => {
    try {
      const url = `${API_PATH.CUSTOMER.BASE}/${id}`
      const response =
        await api.patch<HTTPResponse<null>>(url, body);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  }

  return useMutation({ mutationFn: editCustomer })
}

export function useDeleteCustomer() {
  const deleteCustomer = async (
    id?: number
  ): Promise<HTTPResponse<null>> => {
    try {
      const url = `${API_PATH.CUSTOMER.BASE}/${id}`
      const response =
        await api.delete<HTTPResponse<null>>(url);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  }

  return useMutation({ mutationFn: deleteCustomer })
}
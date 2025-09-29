import type {HTTPResponse} from "@/types/http-response";
import type {User} from "@/types/user";
import {api, API_PATH, catchHTTPError} from "@/lib/api";
import {useMutation, useQuery, useSuspenseQuery} from "@tanstack/react-query";
import type {Role} from "@/types/role";

export function useRoleList() {
  const roles = async (): Promise<HTTPResponse<Role[]>> => {
    try {
      const url = API_PATH.ACCOUNT.ROLE
      const response =
        await api.get<HTTPResponse<Role[]>>(url);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useSuspenseQuery({ queryKey: ['roles'], queryFn: roles })
}

export function useUserList() {
  const users = async (): Promise<HTTPResponse<User[]>> => {
    try {
      const url = API_PATH.ACCOUNT.USER
      const response =
        await api.get<HTTPResponse<User[]>>(url);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useSuspenseQuery({ queryKey: ['users'], queryFn: users })
}

export function useUserDetail(id?: number) {
  const users = async (): Promise<HTTPResponse<User>> => {
    try {
      const url = `${API_PATH.ACCOUNT.USER}/${id}`
      const response =
        await api.get<HTTPResponse<User>>(url);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useQuery({ queryKey: ["user", id], queryFn: users,  enabled: !!id, })
}

export function useUpdateProfile() {
  const profile = async (
    body: string
  ): Promise<HTTPResponse<User>> => {
    try {
      const url = API_PATH.ACCOUNT.USER
      const response =
        await api.put<HTTPResponse<User>>(url, body);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useMutation({ mutationFn: profile })
}

export function useUpdatePassword() {
  const password = async (
    body: string
  ): Promise<HTTPResponse<any>> => {
    try {
      const url = API_PATH.ACCOUNT.USER
      const response =
        await api.patch<HTTPResponse<any>>(url, body);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useMutation({ mutationFn: password })
}

export function useNewUser() {
  const newUser = async (
    body: string
  ): Promise<HTTPResponse<User>> => {
    try {
      const url = API_PATH.ACCOUNT.USER
      const response =
        await api.post<HTTPResponse<User>>(url, body);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useMutation({ mutationFn: newUser })
}

export function useUpdateUser() {
  const editUser = async (
    {id, body}: {id:number, body: string}
  ): Promise<HTTPResponse<User>> => {
    try {
      const url = `${API_PATH.ACCOUNT.USER}/${id}`
      const response =
        await api.put<HTTPResponse<User>>(url, body);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useMutation({ mutationFn: editUser })
}

export function useDeleteUser() {
  const deleteUser = async (
    id?: number
  ): Promise<HTTPResponse<null>> => {
    try {
      const url = `${API_PATH.ACCOUNT.USER}/${id}`
      const response =
        await api.delete<HTTPResponse<null>>(url);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useMutation({ mutationFn: deleteUser })
}
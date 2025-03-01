import {HttpResponse} from "@/lib/types/http-response.ts";

export  const API_URL = `http://localhost:8000/api/v1`

export const Endpoint = {
  Auth: {
    SignIn: `${API_URL}/login`,
    SignOut: `${API_URL}/logout`,
  },
  User: {
    Base: `${API_URL}/users`,
    Role: `${API_URL}/roles`,
  },
  Catalog: {
    Addon: "addons"
  },
  Store: {
    Pref: `${API_URL}/store/prefs`,
  },
  Shift: {
    Base: `${API_URL}/shifts`,
  }
}

export enum HTTPStatusCode {
  OK = 200,
  Created = 201,
  NoContent = 204,
  BadRequest = 400,
  Unauthorized = 401,
  NotFound = 404,
  UnprocessableEntity = 422,
  ToManyRequest = 429,
  InternalServerError = 500,
}

export function isHttpResponse(response: any): response is HttpResponse<string> {
  return typeof response === "object" && response !== null && "code" in response;
}
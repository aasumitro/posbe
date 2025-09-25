import axios, {AxiosError} from "axios";
import {useAuthStore} from "@/states/auth-state";
import {isJWTExpired} from "@/lib/jwt";
import type {HTTPResponse} from "@/types/http-response";

const host =  process.env.NODE_ENV === 'production' ? window.location.host : "localhost:8000"
const SERVER_URL = `${window.location.protocol}//${host}/api/v1`

export const AuthPath = {
  SIGN_IN: "login",
  REFRESH_TOKEN: "refresh",
  SIGN_OUT: "logout"
} as const;

export type AuthPath = typeof AuthPath[keyof typeof AuthPath];

export const API_PATH = {
  ACCOUNT: {
    AUTH: (path: AuthPath) => `/auth/${path}`,
    USER: "/users",
    ROLE: "/roles"
  },
  STORE: {
    SETTINGS: "/store/settings",
    SHIFTS: "/store/shifts",
    FLOORS: "/store/floors",
    TABLES: "/store/tables",
  },
  CATALOG: {
    ATTRIBUTES: {
      UNITS: "/units",
      CATEGORIES: "/categories",
    },
    PRODUCTS: {
      BASE: "/products",
      ADDONS: "/product-addons",
    }
  }
}

export const HTTP_STATUS_CODE = {
  OK: 200,
  CREATED: 201,
  BAD_REQUEST: 400,
  UNAUTHORIZED: 401,
  FORBIDDEN: 403,
  UNPROCESSABLE_ENTITY: 422,
  TO_MANY_REQUEST: 429,
  INTERNAL_SERVER_ERROR: 500
}

export const api = axios.create({
  baseURL: SERVER_URL,
  timeout: process.env.NODE_ENV === 'production' ? 3000 : 10000,
  headers: {"Content-Type": "application/json"},
});

// Automatically add header to requests also check for refresh token
api.interceptors.request.use(async  (config) => {
  const auth = useAuthStore.getState().auth;
  const shouldRefresh = auth.accessToken && isJWTExpired(auth.accessToken)
    && auth.refreshToken && !isJWTExpired(auth.refreshToken);

  // Attach headers
  if (auth.accessToken) config.headers.Authorization = `Bearer ${auth.accessToken}`;
  if (auth.firebaseToken) config.headers["X-FIREBASE-TOKEN"] = auth.firebaseToken;

  // Handle refresh
  if (shouldRefresh) {
    try {
      const refreshTokenURL = SERVER_URL+API_PATH.ACCOUNT.AUTH(AuthPath.REFRESH_TOKEN)
      const response = await axios.post(refreshTokenURL, {},
        {headers: {"X-REFRESH-TOKEN": auth.refreshToken}});

      if (response.status === HTTP_STATUS_CODE.CREATED) {
        const { access_token } = response.data?.data?.token;
        auth.setAccessToken(access_token);
        api.defaults.headers.common["Authorization"] = `Bearer ${access_token}`;
        config.headers.Authorization = `Bearer ${access_token}`;
      }
    } catch (error) {
      auth.reset();
      window.location.href = "/login";
    }
  }

  return config;
});

export const catchHTTPError = (error: unknown) => {
  if (error instanceof AxiosError && error.response) {
    const errorData = error.response.data;

    if (errorData && typeof errorData === "object") {
      throw errorData;
    }

    throw error
  }

  if (axios.isAxiosError(error) && !error.response) {
    throw new Error("Unable to connect to the server. Please check your internet and try again. If the issue persists, try again later.");
  }

  throw new Error("Unexpected error occurred");
}

export function isHTTPResponse<T>(obj: unknown): obj is HTTPResponse<T> {
  return typeof obj === "object" && obj !== null && "status" in obj;
}
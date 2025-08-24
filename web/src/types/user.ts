import type {Role} from "@/types/role";

export interface AccountAccess {
  token: Token;
  user: User;
}

export interface Token {
  access_token: string
  refresh_token: string
}

export interface User {
  id: number,
  name: string,
  username: string,
  email: string,
  Role?: Role
}
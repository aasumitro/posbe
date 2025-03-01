import {Time} from "@/lib/types/common.ts";

export interface Access {
  token: string;
  user: User;
}

export interface User {
  id: number;
  name: string;
  username: string;
  email: string;
  phone: string;
  role: Role;
  created_at: Time
}

export interface Role {
  id: number;
  name: string;
  description: string;
}


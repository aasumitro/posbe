import {Time} from "@/lib/types/common.ts";

export interface Floor {
  id: number
  name: string
  total_tables: number
  created_at: Time
  updated_at: Time
}

export interface Room {
  id: number
  floor_id: number
  name: string
  x_pos: number
  y_pos: number
  w_size: number
  h_size: number
  capacity: number
  price: number
  created_at: Time
  updated_at: Time
}

export interface Table {
  id: number
  floor_id: number
  name: string
  x_pos: number
  y_pos: number
  w_size: number
  h_size: number
  capacity: number
  type: string
  created_at: Time
  updated_at: Time
}


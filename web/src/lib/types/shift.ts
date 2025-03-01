import {Time} from "@/lib/types/common.ts";

export interface Shift {
  id: number,
  name: string,
  start_time: number,
  end_time: number,
  total_usage: number,
  total_transaction: number,
  current_shift: StoreShift | null,
}

export interface StoreShift {
  id: number,
  open_at: Time,
  close_at: Time,
}
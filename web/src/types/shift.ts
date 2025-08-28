export interface StoreShift {
  id: number;
  name: string;
  start_time: number;
  end_time: number;
  total_usage: number;
  total_transaction: number;
  active_shift: ActiveShift | null;
}

export interface ActiveShift {
  id: number;
  shift_id: number;
}
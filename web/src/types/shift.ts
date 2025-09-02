export interface StoreShift {
  id: number;
  name: string;
  start_time: number;
  end_time: number;
  total_usage: number;
  total_transaction: number;
  total_surplus: number;
  total_deficit: number;
  profit: number;
  loss: number;
  net: number;
  active: ActiveShift | null;
  last: ActiveShift | null;
  users_id: number[] | null;
}

export interface ActiveShift {
  id: number;
  shift_id: number;
}
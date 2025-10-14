export interface Floor {
  id: number,
  name: string,
  total_tables?: number
  tables?: Table[]
}

export interface Table {
  id: number,
  floor_id: number,
  name: string,
  x_pos: number,
  y_pos: number,
  w_size: number,
  h_size: number,
  d_size: number,
  capacity: number,
  type: string,
  status: string,
  has_order: boolean,
  has_history: boolean,
}
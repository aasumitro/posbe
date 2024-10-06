export interface Unit {
  id: number
  magnitude: string
  name: string
  symbol: string
}

export interface Category {
  id: number
  name: string
}

export interface Subcategory {
  id: number
  category_id: number
  name: string
}

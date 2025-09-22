export interface Category {
  id: number,
  name: string,
  usage: number,
  subcategories?: Subcategory[],
}

export interface Subcategory {
  id: number,
  category_id:  number,
  name: string,
  usage: number,
}
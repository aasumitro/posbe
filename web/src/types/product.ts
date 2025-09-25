import type {Category, Subcategory} from "@/types/category";
import type {Unit} from "@/types/unit";

export interface ProductAddon {
  id: number,
  price: number,
  name: string,
  description: string
  usage: number
}

export interface Product {
  id: number,
  category_id: number,
  subcategory_id: number,
  sku: string,
  image: string,
  name: string,
  description: string,
  category: Category | null,
  subcategory: Subcategory | null,
  variants: ProductVariant[] | null,
  sales_today: number
  sales_this_week: number
  sales_this_year: number
}

export interface ProductVariant {
  id: number,
  product_id: number,
  unit_id: number,
  unit_size: number,
  type: string,
  name: string,
  description: string,
  price: number,
  unit: Unit | null,
}
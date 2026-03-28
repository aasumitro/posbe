import type {User} from "@/types/user";
import type {ActiveShift} from "@/types/shift";
import type {Table} from "@/types/seating";
import type {Product, ProductAddon, ProductVariant} from "@/types/product";
import type {Category, Subcategory} from "@/types/category";

export interface Order {
  id: number;
  cashier_id: number;
  shift_id: number;
  table_id: number;
  time_open: number;
  time_close: number;
  // customer_id: number;

  // main data
  gross: number;
  discount: number;
  net: number;
  tax: number;
  total: number;
  payment_method: string;
  payment_amount: number;
  change: number;
  notes: string;
  cancel_reason: string;
  status: string;

  // relation/extend data
  cashier?: User | null;
  shift?: ActiveShift | null;
  // customer?: Customer | null;
  table?: Table | null;
  order_products?: OrderProduct[] | null;
  order_product_addons?: OrderProductAddon[] | null;
  order_product_options?: OrderProductOption[] | null;
}

export interface OrderProduct {
  id: number;
  order_id: number;
  product_id: number;
  category_id: number;
  subcategory_id: number;

  // main data
  name: string;
  quantity: number;
  price: number;
  net: number;
  notes: string;

  // relation
  order?: Order | null;
  product?: Product | null;
  category?: Category | null;
  subcategory?: Subcategory | null;
  order_product_addons?: OrderProductAddon[] | null;
  order_product_options?: OrderProductOption[] | null;
}

export interface OrderProductOption {
  id: number;
  order_id: number;
  order_product_id: number;
  variant_id: number;

  name: string;
  value: string;
  price: number;

  order?: Order | null;
  order_product?: OrderProduct | null;
  variant?: ProductVariant | null;
}

export interface OrderProductAddon {
 id: number;
 order_id: number;
 order_product_id: number;
 addon_id: number;

 name: string;
 quantity: number;
 price: number;
 net: number;
 notes: string;

 order?: Order | null;
 order_product?: OrderProduct | null;
 addon?: ProductAddon | null;
}
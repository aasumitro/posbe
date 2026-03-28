import {create} from "zustand";
import type {ActiveShift} from "@/types/shift";
import Cookies from "js-cookie";
import type {Order} from "@/types/order";
import type {Product} from "@/types/product";

const DEFAULT_FLOOR = 'posbe-order-default-floor'

interface States {
  activeShift: ActiveShift | null;
  defaultFloorId: number;
  orders: Order[] | null;
  selectedOrder: Order | null;
  products: Product[] | null;
  selectedProduct: Product | null;
}

interface Actions {
  setActiveShift(activeShift: ActiveShift | null): void;
  setDefaultFloor: (defaultFloorId: number) => void;
  setOrdes: (orders: Order[] | null) => void;
  setSelectedOrder(selectedOrder: Order | null): void;
  setProducts(products: Product[] | null): void;
  setSelectedProduct(selectedProduct: Product | null): void;
}

export const useOrderState = create<States & Actions>((set) => {
  const defaultFloorCookieState = Cookies.get(DEFAULT_FLOOR)
  const initDefaultFloor = defaultFloorCookieState ? Number(defaultFloorCookieState) : 0

  return {
    defaultFloorId: initDefaultFloor,
    activeShift: null,
    orders: null,
    selectedOrder: null,
    products: null,
    selectedProduct: null,

    setActiveShift: (activeShift: ActiveShift | null) => set({activeShift}),
    setDefaultFloor: (defaultFloorId: number) =>
      set(() => {
        Cookies.set(DEFAULT_FLOOR, String(defaultFloorId));
        return { defaultFloorId };
      }),
    setOrdes: (orders: Order[] | null) => set({orders}),
    setSelectedOrder: (selectedOrder: Order | null) => set({selectedOrder}),
    setProducts: (products: Product[] | null) => set({products}),
    setSelectedProduct: (selectedProduct: Product | null) => set({selectedProduct})
  }
})
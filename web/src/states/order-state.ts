import {create} from "zustand";
import type {ActiveShift} from "@/types/shift";
import Cookies from "js-cookie";
import type {Order} from "@/types/order";

const DEFAULT_FLOOR = 'posbe-order-default-floor'

interface States {
  activeShift: ActiveShift | null;
  defaultFloorId: number;
  orders: Order[] | null;
}

interface Actions {
  setActiveShift(activeShift: ActiveShift | null): void;
  setDefaultFloor: (defaultFloorId: number) => void;
  setOrdes: (orders: Order[] | null) => void;
}

export const useOrderState = create<States & Actions>((set) => {
  const defaultFloorCookieState = Cookies.get(DEFAULT_FLOOR)
  const initDefaultFloor = defaultFloorCookieState ? Number(defaultFloorCookieState) : 0

  return {
    defaultFloorId: initDefaultFloor,
    activeShift: null,
    orders: null,
    setActiveShift: (activeShift: ActiveShift | null) => set({activeShift}),
    setDefaultFloor: (defaultFloorId: number) =>
      set(() => {
        Cookies.set(DEFAULT_FLOOR, String(defaultFloorId));
        return { defaultFloorId };
      }),
    setOrdes: (orders: Order[] | null) => set({orders}),
  }
})
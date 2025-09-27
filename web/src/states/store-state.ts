import type {StoreSetting} from "@/types/store";
import {create} from "zustand";
import type {StoreShift} from "@/types/shift";
import Cookies from "js-cookie";

const STORE_SETTINGS = 'posbe-store-settings'

interface States {
  settings: StoreSetting | null;
  shifts: StoreShift[] | null;
  selectedShift: StoreShift | null;
}

interface Actions {
  setSettings(settings: StoreSetting): void;
  setShifts(shifts: StoreShift[]): void;
  setSelectedShifts(selectedShift: StoreShift | null): void;
}

export const useStoreState = create<States & Actions>((set) => {
  const storeSettingCookieState = Cookies.get(STORE_SETTINGS)
  const initStoreSetting = storeSettingCookieState ? JSON.parse(storeSettingCookieState) : ''

  return {
    settings: initStoreSetting,
    shifts: null,
    selectedShift: null,
    setSettings: (settings: StoreSetting)  => set((state) => {
      Cookies.set(STORE_SETTINGS, JSON.stringify(settings))
      return { ...state, settings }
    }),
    setShifts: (shifts: StoreShift[]) => set((state) => {
      let selectedShift = state.selectedShift;

      if (selectedShift) {
        const found = shifts.find((s) =>
          s.id === selectedShift?.id);
        selectedShift =found ?? null;
      }

      return { shifts, selectedShift };
    }),
    setSelectedShifts: (selectedShift: StoreShift | null) => set({selectedShift}),
  }
})
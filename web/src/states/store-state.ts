import type {StoreSetting} from "@/types/store";
import {create} from "zustand";
import type {StoreShift} from "@/types/shift";

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
  return {
    settings: null,
    shifts: null,
    selectedShift: null,
    setSettings: (settings: StoreSetting)  => set({settings}),
    setShifts: (shifts: StoreShift[]) => set({shifts}),
    setSelectedShifts: (selectedShift: StoreShift | null) => set({selectedShift}),
  }
})
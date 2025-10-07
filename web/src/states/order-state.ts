import {create} from "zustand";
import type {ActiveShift} from "@/types/shift";

interface States {
  activeShift: ActiveShift | null;
  defaultFloorId: number;
}

interface Actions {
  setActiveShift(activeShift: ActiveShift | null): void;
  setDefaultFloor: (defaultFloorId: number) => void;
}

export const useOrderState = create<States & Actions>((set) => {
  return {
    defaultFloorId: 0,
    activeShift: null,
    setDefaultFloor: (defaultFloorId: number) => set({defaultFloorId}),
    setActiveShift: (activeShift: ActiveShift | null) => set({activeShift}),
  }
})
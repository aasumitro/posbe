import type {StoreSetting} from "@/types/store";
import {create} from "zustand";

interface States {
  settings: StoreSetting | null;
}

interface Actions {
  setSettings(settings: StoreSetting): void;
}

export const useStoreState = create<States & Actions>((set) => {
  return {
    settings: null,
    setSettings: (settings: StoreSetting)  => set({settings}),
  }
})
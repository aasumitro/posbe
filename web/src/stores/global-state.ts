import { create } from "zustand";

interface BooleanState {
  [key: string]: boolean;
}

interface StringState {
  [key: string]: string;
}

interface States {
  boolStates: BooleanState;
  strStates: StringState;
}

interface Actions {
  setBoolState: (key: string, value: boolean) => void;
  setStrState: (key: string, value: string) => void;
  cleanBoolState: () => void;
  cleanStrState: () => void;
  cleanupState: () => void;
}

export const useGlobalStateStore = create<States & Actions>((set) => ({
  boolStates: {},
  strStates: {},
  setBoolState: (key: string, value: boolean) =>
    set((state) => {
      const newBoolStates = { ...state.boolStates };
      if (value) {
        newBoolStates[key] = value;
      } else {
        delete newBoolStates[key];
      }
      return { boolStates: newBoolStates };
    }),
  setStrState: (key: string, value: string) =>
    set((state) => {
      const newStrStates = { ...state.strStates };
      if (value) {
        newStrStates[key] = value;
      } else {
        delete newStrStates[key];
      }
      return { strStates: newStrStates };
    }),
  cleanBoolState: () => set({ boolStates: {} }),
  cleanStrState: () => set({ strStates: {} }),
  cleanupState: () => set({ boolStates: {}, strStates: {} }),
}));
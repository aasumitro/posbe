import { create } from "zustand";

interface BoolState {
  [key: string]: boolean;
}

interface StrState {
  [key: string]: string;
}

interface States {
  bool: BoolState;
  str: StrState;
}

interface Actions {
  setBoolState: (key: string, value: boolean) => void;
  setStrState: (key: string, value: string) => void;
  resetState: () => void;
}

export const useActionState = create<States & Actions>((set) => ({
  bool: {},
  str: {},

  setBoolState: (key: string, value: boolean) =>
    set((state) => {
      const newBool = { ...state.bool };
      if (value) {
        newBool[key] = value;
      } else {
        delete newBool[key];
      }
      return { bool: newBool };
    }),

  setStrState: (key: string, value: string) =>
    set((state) => {
      const newStr = { ...state.str };
      if (value) {
        newStr[key] = value;
      } else {
        delete newStr[key];
      }
      return { str: newStr };
    }),

  resetState: () => set({ bool: {}, str: {} }),
}));
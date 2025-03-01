import { create } from "zustand";
import {User} from "@/lib/types/user.ts";

interface States {
  token: string | null;
  user: User | null;
}

interface Actions {
  setUserToken: (token: string | null) => void;
  setUserData: (user: User | null) => void;
}

export const useSessionStateStore = create<States & Actions>((set) => {
  return {
    token: null,
    user: null,
    setUserToken: (token: string | null) => set({ token }),
    setUserData: (user: User | null) => set({ user }),
  }
})
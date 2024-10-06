import { create } from "zustand";
import {Role, User} from "@/lib/types/user.ts";
import {Store} from "@/lib/types/store.ts";

interface States {
  store: Store | null;
  users: User[] | null;
  roles: Role[] | null;
  user: User | null;
}

interface Actions {
  setStoreData: (store: Store | null) => void;
  setUserList: (users: User[] | null) => void;
  setUserRoleList: (roles: Role[] | null) => void;
  setUserDetail: (user: User | null) => void
}

export const useStorePageState = create<States & Actions>((set) => {
  return {
    store: null,
    users: null,
    user: null,
    roles: null,
    setStoreData: (store: Store | null) => set({ store }),
    setUserList: (users: User[] | null) => set({ users }),
    setUserRoleList: (roles: Role[] | null) => set({ roles }),
    setUserDetail: (user: User | null) => set({ user }),
  }
})
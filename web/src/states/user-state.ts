import {create} from "zustand";
import type {Role} from "@/types/role";
import type {User} from "@/types/user";

interface States {
  roles: Role[] | null
  users: User[] | null
  selectedUser: User | null
}

interface Actions {
  setRoles: (roles: Role[] | null) => void
  setUsers: (users: User[] | null) => void
  setSelectedUser: (user: User | null) => void
}

export const useUserState = create<States & Actions>((set) => {
  return {
    roles: null,
    users: null,
    selectedUser: null,
    setRoles: (roles: Role[] | null) => set({roles}),
    setUsers: (users: User[] | null) => set({users}),
    setSelectedUser: (selectedUser: User | null) => set({selectedUser})
  }
})

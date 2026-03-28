import {create} from "zustand";
import type {Customer} from "@/types/customer";

interface States {
  customers: Customer[] | null
  selectedCustomer: Customer | null
}

interface Actions {
  setCustomers: (customers: Customer[] | null) => void
  setSelectedCustomer: (selectedCustomer: Customer | null) => void
}
export const useCustomerState = create<States & Actions>((set) => {
  return {
    customers: null,
    selectedCustomer: null,
    setCustomers: (customers: Customer[] | null) => set({customers}),
    setSelectedCustomer: (selectedCustomer: Customer | null) => set({selectedCustomer}),
  }
})
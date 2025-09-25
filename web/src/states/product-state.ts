import type {Product, ProductAddon} from "@/types/product";
import {create} from "zustand";

interface States {
  addons: ProductAddon[] | null;
  selectedAddon: ProductAddon | null;
  products: Product[] | null;
}

interface Actions {
  setAddons(addons: ProductAddon[] | null): void;
  setSelectedAddon(selectedAddon: ProductAddon | null): void;
  setProducts(products: Product[] | null): void;
}

export const useProductState = create<States & Actions>((set) => {
  return {
    addons: null,
    selectedAddon: null,
    products: null,
    setAddons: (addons: ProductAddon[] | null) => set({addons}),
    setSelectedAddon: (selectedAddon: ProductAddon | null) => set({selectedAddon}),
    setProducts: (products: Product[] | null) => set({products}),
  }
})
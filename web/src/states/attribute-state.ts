import type {Unit} from "@/types/unit";
import type {Category, Subcategory} from "@/types/category";
import {create} from "zustand";

interface States {
  units: Unit[] | null;
  selectedUnit: Unit | null;
  categories: Category[] | null;
  selectedCategory: Category | null;
  subcategories: Subcategory[] | null;
}

interface Actions {
  setUnits(units: Unit[]): void;
  setSelectedUnit(selectedUnit: Unit | null): void;
  setCategories(categories: Category[]): void;
  setSelectedCategory(selectedCategory: Category | null): void;
  setSubcategories(subcategories: Subcategory[]): void;
}

export const useAttributeState = create<States & Actions>((set) => {
  return {
    units: null,
    selectedUnit: null,
    categories: null,
    selectedCategory: null,
    subcategories: null,
    setUnits: (units: Unit[]) => set({units}),
    setSelectedUnit: (selectedUnit: Unit | null) => set({selectedUnit}),
    setCategories: (categories: Category[])  => set((state) => {
      let selectedCategory = state.selectedCategory;

      if (selectedCategory) {
        const found = categories.find((c) =>
          c.id === selectedCategory?.id);
        selectedCategory = found ?? null;
      }

      // set the subcategory
      const subcategories = categories.flatMap((c) => c.subcategories || []);

      return {categories, selectedCategory, subcategories}
    }),
    setSelectedCategory: (selectedCategory: Category | null) => set({selectedCategory}),
    setSubcategories: (subcategories: Subcategory[] | null) => set({subcategories}),
  }
})
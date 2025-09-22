import type {Unit} from "@/types/unit";
import type {Category} from "@/types/category";
import {create} from "zustand";

interface States {
  units: Unit[] | null;
  selectedUnit: Unit | null;
  categories: Category[] | null;
  selectedCategory: Category | null;
}

interface Actions {
  setUnits(units: Unit[]): void;
  setSelectedUnit(selectedUnit: Unit | null): void;
  setCategories(categories: Category[]): void;
  setSelectedCategory(selectedCategory: Category | null): void;
}

export const useAttributeState = create<States & Actions>((set) => {
  return {
    units: null,
    selectedUnit: null,
    categories: null,
    selectedCategory: null,
    setUnits: (units: Unit[]) => set({units}),
    setSelectedUnit: (selectedUnit: Unit | null) => set({selectedUnit}),
    setCategories: (categories: Category[])  => set({categories}),
    setSelectedCategory: (selectedCategory: Category | null) => set({selectedCategory}),
  }
})
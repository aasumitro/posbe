import type {Floor, Table} from "@/types/seating";
import {create} from "zustand";

interface States {
  floors: Floor[] | null;
  tables: Table[] | null;
  selectedFloor: Floor | null;
  selectedTableId: number | null;
}

interface Actions {
  setFloors: (floors: Floor[] | null) => void;
  setTables: (tables: Table[] | null) => void;
  setSelectedFloor: (selectedFloor: Floor | null) => void;
  setSelectedTableId: (selectedTableId: number | null) => void;
}

export const useSeatingState = create<States & Actions>((set) => {
  return {
    floors: null,
    tables: null,
    selectedFloor: null,
    selectedTableId: null,
    setFloors: (floors: Floor[] | null) => set({floors}),
    setTables: (tables: Table[] | null) => set({tables}),
    setSelectedFloor: (selectedFloor: Floor | null) => set({selectedFloor}),
    setSelectedTableId: (selectedTableId: number | null) => set({selectedTableId}),
  }
})
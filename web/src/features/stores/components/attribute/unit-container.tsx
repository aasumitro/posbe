import {UnitCard} from "@/features/stores/components/attribute/unit-card";
import type {Unit} from "@/types/unit";
import {UnitActionAdd} from "@/features/stores/components/attribute/unit-action-add";

export function UnitContainer() {
  const units : Unit[] = [
    {id: 1, magnitude: "mass", symbol: "mg", name: "milligram"},
    {id: 2, magnitude: "mass", symbol: "g", name: "gram"},
    {id: 3, magnitude: "mass", symbol: "kg", name: "kilogram"},
    {id: 4, magnitude: "volume", symbol: "ml", name: "milliliter"},
    {id: 5, magnitude: "volume", symbol: "l", name: "liter"},
    {id: 6, magnitude: "volume", symbol: "c", name: "cup"},
    {id: 7, magnitude: "count", symbol: "prt", name: "portion"},
    {id: 8, magnitude: "count", symbol: "box", name: "box"},
  ]

  return (
    <div className="grid auto-rows-min gap-4 grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4 py-4">
      <UnitActionAdd />

      {units.map((item) => (
          <UnitCard key={item.id} {...item} />
      ))}
    </div>
  )
}
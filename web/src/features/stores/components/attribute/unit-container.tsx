import {UnitCard} from "@/features/stores/components/attribute/unit-card";
import {UnitActionAdd} from "@/features/stores/components/attribute/unit-action-add";
import {useAttributeState} from "@/states/attribute-state";
import {UnitActionDelete} from "@/features/stores/components/attribute/unit-action-delete";
import {UnitActionEdit} from "@/features/stores/components/attribute/unit-action-edit";

interface UnitContainerProps {
  sort?: Record<string, "asc" | "desc">;
}

export function UnitContainer({sort}: UnitContainerProps) {
  const {units} =  useAttributeState();

  let sortedUnits = [...(units ?? [])];

  if (sort) {
    const [field, order] = Object.entries(sort)[0] ?? [];
    if (field && order) {
      sortedUnits.sort((a, b) => {
        const aValue = a[field as keyof typeof a];
        const bValue = b[field as keyof typeof b];

        if (typeof aValue === "string" && typeof bValue === "string") {
          return order === "asc"
            ? aValue.localeCompare(bValue)
            : bValue.localeCompare(aValue);
        }

        if (typeof aValue === "number" && typeof bValue === "number") {
          return order === "asc" ? aValue - bValue : bValue - aValue;
        }

        return 0;
      });
    }
  }

  return (
    <div className="grid auto-rows-min gap-4 grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4 py-4">
      <UnitActionAdd />

      {sortedUnits?.map((item) => (
          <UnitCard key={item.id} {...item} />
      ))}

      <UnitActionEdit />
      <UnitActionDelete />
    </div>
  )
}
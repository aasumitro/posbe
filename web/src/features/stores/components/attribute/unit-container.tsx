import {UnitCard} from "@/features/stores/components/attribute/unit-card";
import {UnitActionAdd} from "@/features/stores/components/attribute/unit-action-add";
import {useAttributeState} from "@/states/attribute-state";
import {UnitActionDelete} from "@/features/stores/components/attribute/unit-action-delete";
import {UnitActionEdit} from "@/features/stores/components/attribute/unit-action-edit";

export function UnitContainer() {
  const {units} =  useAttributeState();

  return (
    <div className="grid auto-rows-min gap-4 grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4 py-4">
      <UnitActionAdd />

      {units?.map((item) => (
          <UnitCard key={item.id} {...item} />
      ))}

      <UnitActionEdit />
      <UnitActionDelete />
    </div>
  )
}
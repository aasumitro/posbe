import {CategoryCard} from "@/features/stores/components/attribute/category-card";
import {CategoryActionAdd} from "@/features/stores/components/attribute/category-action-add";
import {useAttributeState} from "@/states/attribute-state";
import {CategoryActionDelete} from "@/features/stores/components/attribute/category-action-delete";

export function CategoryContainer() {
  const {categories} =  useAttributeState();

  return (
    <div className="grid auto-rows-fr gap-4 grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4 py-4">
      <CategoryActionAdd />

      {categories?.map((item) => (
        <CategoryCard key={item.id} {...item} />
      ))}

      <CategoryActionDelete />
    </div>
  )
}
import {CategoryCard} from "@/features/stores/components/attribute/category-card";
import {CategoryActionAdd} from "@/features/stores/components/attribute/category-action-add";
import {useAttributeState} from "@/states/attribute-state";
import {CategoryActionDelete} from "@/features/stores/components/attribute/category-action-delete";
import {CategoryActionEdit} from "@/features/stores/components/attribute/category-action-edit";

interface CategoryContainerProps {
  sort?: Record<string, "asc" | "desc">;
}

export function CategoryContainer({sort}: CategoryContainerProps) {
  const {categories} =  useAttributeState();

  let sortedCategories = [...(categories ?? [])];

  if (sort) {
    const [field, order] = Object.entries(sort)[0] ?? [];
    if (field && order) {
      sortedCategories.sort((a, b) => {
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
    <div className="grid auto-rows-fr gap-4 grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4 py-4">
      <CategoryActionAdd />

      {sortedCategories?.map((item) => (
        <CategoryCard key={item.id} {...item} />
      ))}

      <CategoryActionEdit />
      <CategoryActionDelete />
    </div>
  )
}
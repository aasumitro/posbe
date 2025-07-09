import type {Category, Subcategory} from "@/types/category";
import {useEffect, useState} from "react";
import {CategoryCard} from "@/features/stores/components/attribute/category-card";
import {CategoryActionAdd} from "@/features/stores/components/attribute/category-action-add";

export function CategoryContainer() {
  const [categoryList, setCategoryList] = useState<Category[]>([]);

  const categories: Category[] = [
    {id: 1, name: "foods"},
    {id: 2, name: "beverages"},
  ]

  const subcategories: Subcategory[] = [
    {id: 1, category_id: 1, name: "meat"},
    {id: 2, category_id: 1, name: "seafood"},
    {id: 3, category_id: 2, name: "coffee"},
    {id: 4, category_id: 2, name: "juice"},
  ]

  useEffect(() => {
    const data = categories.map(category => ({
      ...category,
      subcategories: subcategories.filter(sub => sub.category_id === category.id)
    }));
    setCategoryList(data)
  }, []);

  return (
    <div className="grid auto-rows-min gap-4 grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4 py-4">
      <CategoryActionAdd />

      {categoryList.map((item) => (
        <CategoryCard key={item.id} {...item} />
      ))}
    </div>
  )
}
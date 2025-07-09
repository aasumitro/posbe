import { Grid } from "lucide-react"

const categories = [
  { icon: Grid, label: "All", items: "2 Items", active: true },
  { label: "Beef", items: "1 Items" },
  { label: "Tea", items: "1 Items" },
  { label: "Chicken", items: "0 Items" },
  { label: "Lamb", items: "0 Items" },
  { label: "Fish", items: "0 Items" },
  { label: "Shrimp", items: "0 Items" },
  { label: "Crab", items: "0 Items" },
  { label: "Clam", items: "0 Items" },
  { label: "Juice", items: "0 Items" },
  { label: "Coffee", items: "0 Items" },
  { label: "Chocolate", items: "0 Items" },
  { label: "Milk", items: "0 Items" },
]

export function SubcategoryFilter() {
  return (
    <div className="flex gap-3 overflow-x-auto pb-4 px-5 ml-[-20px] mr-[-12px]">
      {categories.map((category, index) => (
        <div
          key={index}
          className={`flex flex-col items-center justify-center p-3 rounded-xl min-w-[100px] ${
            category.active ? "bg-green-50 text-green-600" : "bg-white"
          } border cursor-pointer hover:bg-green-50`}
        >
          {category.icon && <category.icon className="h-6 w-6 mb-1"/>}
          <span className="text-sm font-medium">{category.label}</span>
          <span className="text-xs text-gray-500">{category.items}</span>
        </div>
      ))}
    </div>
  )
}

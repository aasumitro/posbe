import {PlusIcon} from "lucide-react";

export const NewProductCard = () => {
  return (
    <button
      className="flex flex-col items-center justify-center w-full lg:w-96 min-h-60 border-2 border-dashed rounded-lg hover:bg-gray-100 gap-4">
      <PlusIcon/>
      <section>
        <h3 className="text-2xl font-semibold leading-none tracking-tight">New Product</h3>
        <p className="text-sm text-muted-foreground">Easily add a new product to your inventory.</p>
      </section>
    </button>
  )
}
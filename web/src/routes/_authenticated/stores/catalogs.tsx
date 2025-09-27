import {createFileRoute, useSearch} from "@tanstack/react-router"
import { ProductCatalogPage } from "@/features/stores/catalogs";
import {z} from "zod";

const catalogSearchSchema = z.object({
  tab: z.enum(['products', 'addons']).catch('products'),
  sort: z.string().optional()
    .transform((val) => {
      if (!val) return undefined;
      const [key, order] = val.split(":");
      if (!key || !order) return undefined;
      if (order !== "asc" && order !== "desc") return undefined;
      return `${key}:${order}` as const;  // keep as string "name:asc" for the URL
    }),
  status: z.enum(["draft", "active", "inactive"]).optional(),
  category: z.string().optional()
})

export const Route = createFileRoute("/_authenticated/stores/catalogs")({
  validateSearch: (search) => catalogSearchSchema.parse(search),
  component: CatalogPageRoute,
})

function CatalogPageRoute() {
  const { tab, sort, status, category } = useSearch({from: '/_authenticated/stores/catalogs'})

  let sortObj: Record<string, "asc" | "desc"> = {};
  if (sort) {
    const [key, direction] = sort.split(":");
    if (key && (direction === "asc" || direction === "desc")) {
      sortObj[key] = direction;
    }
  }

  return <ProductCatalogPage query={{ tab, sort: sortObj, status, category }}  />
}
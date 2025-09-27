import {createFileRoute, useSearch} from "@tanstack/react-router"
import { MasterDataPage } from "@/features/stores/attributes";
import { z } from 'zod'

const attributeSearchSchema = z.object({
  tab: z.enum(['units', 'categories']).catch('units'),
  sort: z.string().optional()
    .transform((val) => {
      if (!val) return undefined;
      const [key, order] = val.split(":");
      if (!key || !order) return undefined;
      if (order !== "asc" && order !== "desc") return undefined;
      return `${key}:${order}` as const;  // keep as string "name:asc" for the URL
    }),
})

export const Route = createFileRoute("/_authenticated/stores/attributes")({
  validateSearch: (search) => attributeSearchSchema.parse(search),
  component: AttributePageRoute,
})

function AttributePageRoute() {
  const { tab, sort } = useSearch({from: '/_authenticated/stores/attributes'})

  let sortObj: Record<string, "asc" | "desc"> = {};
  if (sort) {
    const [key, direction] = sort.split(":");
    if (key && (direction === "asc" || direction === "desc")) {
      sortObj[key] = direction;
    }
  }

  return <MasterDataPage query={{ tab, sort: sortObj }} />
}

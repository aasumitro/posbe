import { createFileRoute } from "@tanstack/react-router"
import { ProductCatalogPage } from "@/features/stores/catalogs";

export const Route = createFileRoute("/_authenticated/stores/catalogs")({
  component: ProductCatalogPage,
})
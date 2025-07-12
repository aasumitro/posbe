import { createFileRoute } from "@tanstack/react-router"
import { MasterDataPage } from "@/features/stores/attributes";

export const Route = createFileRoute("/_authenticated/stores/attributes")({
  component: MasterDataPage,
})
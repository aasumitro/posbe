import { createLazyFileRoute } from "@tanstack/react-router"
import { FloorPlanPage } from "@/features/stores/floors";

export const Route = createLazyFileRoute("/_authenticated/stores/floors")({
  component: FloorPlanPage,
})
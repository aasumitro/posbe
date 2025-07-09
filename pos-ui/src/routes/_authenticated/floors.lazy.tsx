import { createLazyFileRoute } from "@tanstack/react-router"
import { FloorPage } from "@/features/floors";

export const Route = createLazyFileRoute("/_authenticated/floors")({
  component: FloorPage,
})
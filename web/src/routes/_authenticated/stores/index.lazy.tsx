import { createLazyFileRoute } from "@tanstack/react-router"
import { StorePage } from "@/features/stores";

export const Route = createLazyFileRoute("/_authenticated/stores/")({
  component: StorePage,
})
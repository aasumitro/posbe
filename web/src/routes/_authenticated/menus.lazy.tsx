import { createLazyFileRoute } from "@tanstack/react-router"
import { MenuPage } from "@/features/menus";

export const Route = createLazyFileRoute("/_authenticated/menus")({
  component: MenuPage,
})
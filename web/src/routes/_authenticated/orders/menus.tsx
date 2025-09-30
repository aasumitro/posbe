import { createFileRoute } from '@tanstack/react-router'
import { MenuPage } from "@/features/orders/menus";

export const Route = createFileRoute('/_authenticated/orders/menus')({
  component: MenuPage,
})

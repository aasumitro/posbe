import { createFileRoute } from '@tanstack/react-router'
import { FloorPage } from "@/features/orders/floors";

export const Route = createFileRoute('/_authenticated/orders/floors')({
  component: FloorPage,
})

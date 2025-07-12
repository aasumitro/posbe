import { createFileRoute } from "@tanstack/react-router"
import { CustomersPage } from "@/features/stores/customers";

export const Route = createFileRoute("/_authenticated/stores/customers")({
  component: CustomersPage,
})
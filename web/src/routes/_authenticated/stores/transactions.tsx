import { createFileRoute } from "@tanstack/react-router"
import { TransactionPage } from "@/features/stores/transactions";

export const Route = createFileRoute("/_authenticated/stores/transactions")({
  component: TransactionPage,
})
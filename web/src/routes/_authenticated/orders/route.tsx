import {createFileRoute} from '@tanstack/react-router'
import {OrderLayout} from "@/features/orders/layout";

export const Route = createFileRoute('/_authenticated/orders')({
  component: OrderLayout,
})


import { createLazyFileRoute } from '@tanstack/react-router'
import {NewProductPage} from "@/features/stores/product-add";

export const Route = createLazyFileRoute('/_authenticated/stores/products/')({
  component: NewProductPage,
})

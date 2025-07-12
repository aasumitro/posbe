import { createLazyFileRoute } from '@tanstack/react-router'
import {EditProductPage} from "@/features/stores/product-edit";

export const Route = createLazyFileRoute('/_authenticated/stores/products/$id')({
  component: EditProductPage,
})
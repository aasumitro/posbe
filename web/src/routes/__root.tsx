import { createRootRouteWithContext } from '@tanstack/react-router'
import MainLayout from '@/layouts/main'

export const Route = createRootRouteWithContext()({
  component: MainLayout,
})
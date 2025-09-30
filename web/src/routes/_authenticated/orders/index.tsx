import {createFileRoute, redirect} from '@tanstack/react-router'

export const Route = createFileRoute('/_authenticated/orders/')({
  beforeLoad: () => { throw redirect({to: '/orders/menus'}); },
})

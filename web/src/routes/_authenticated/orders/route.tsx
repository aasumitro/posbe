import {createFileRoute, Outlet} from '@tanstack/react-router'
import {useEffect} from "react";
// import {useCategoryList} from "@/hooks/use-attribute";

export const Route = createFileRoute('/_authenticated/orders')({
  // beforeLoad: () => {},
  component: OrderLayout,
})

function OrderLayout() {
  // TODO:
  // 1. load attributes - units, categories & subcategories
  // 2. load shift check if theres active shift or not
  // 3. load floors & its own table also subscribe (real time) to table status
  // 4. load active orders
  // 5. load products
  // 6. load customers (api wip)

  // const {data: categories} = useCategoryList()

  useEffect(() => {

  }, []);

  return (
    <Outlet />
  )
}

import {createFileRoute, Outlet, useRouterState} from '@tanstack/react-router'
import {Fragment, useEffect} from "react";
import {useActionState} from "@/states/action-state";
import {OpenShiftModalState, ShiftActionOpenAlertDialog} from "@/components/shift-action-open-alert-dialog";
import {ShiftActionCloseAlertDialog} from "@/components/shift-action-close-alert-dialog";
import {useStoreShifts} from "@/hooks/use-store";
import {useStoreState} from "@/states/store-state";
import {useActiveShift} from "@/hooks/use-order";
import {useOrderState} from "@/states/order-state";
// import {useCategoryList} from "@/hooks/use-attribute";

export const Route = createFileRoute('/_authenticated/orders')({
  component: OrderLayout,
})

function OrderLayout() {
  const { setBoolState } = useActionState();
  const { location } = useRouterState();
  const pathname = location.pathname;

  // TODO:
  // 1. load attributes - units, categories & subcategories
  // 2. load shift check if theres active shift or not
  // 3. load floors & its own table also subscribe (real time) to table status
  // 4. load active orders
  // 5. load products
  // 6. load customers (api wip)

  const { data: activeShift, isFetching, isSuccess } = useActiveShift();
  const {setActiveShift} =  useOrderState();
  useEffect(() => {
    if (!isFetching && isSuccess) {
      if (activeShift?.data) setActiveShift(activeShift.data);
      else setBoolState(OpenShiftModalState, true);
    }
  }, [activeShift?.data, isFetching, isSuccess]);

  const {data: shifts} = useStoreShifts();
  const {setShifts} =  useStoreState();
  useEffect(() => {
    if (shifts?.data) setShifts(shifts.data);
  }, [shifts?.data]);

  // const {data: categories} = useCategoryList()


  return (
    <Fragment>
      <Outlet />
      {pathname.startsWith('/orders/') && (
        <>
          <ShiftActionOpenAlertDialog />
          <ShiftActionCloseAlertDialog />
        </>
      )}
    </Fragment>
  )
}

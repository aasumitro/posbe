import {useActionState} from "@/states/action-state";
import {Outlet, useRouterState} from "@tanstack/react-router";
import {useActiveShift} from "@/hooks/use-order";
import {useOrderState} from "@/states/order-state";
import {Fragment, useEffect, useMemo} from "react";
import {OpenShiftModalState, ShiftActionOpenAlertDialog} from "@/components/shift-action-open-alert-dialog";
import {useStoreShifts} from "@/hooks/use-store";
import {useStoreState} from "@/states/store-state";
import {useCategoryList} from "@/hooks/use-attribute";
import {useAttributeState} from "@/states/attribute-state";
import {useFloorList} from "@/hooks/use-seating";
import {useSeatingState} from "@/states/seating-state";
import {useProductList} from "@/hooks/use-product";
import {OrderListSheet} from "@/features/orders/components/order-list-sheet";
import {ShiftActionCloseAlertDialog} from "@/components/shift-action-close-alert-dialog";
import {AddToCartModal} from "@/features/orders/components/add-to-cart-modal";

export function OrderLayout() {
  const { setBoolState } = useActionState();
  const { location } = useRouterState();
  const pathname = location.pathname;

  // TODO:
  // 1. load attributes - units, categories & subcategories ✅
  // 2. load shift check if theres active shift or not ✅
  // 3. load floors & its own table also subscribe (real time) to table status ✅
  // 4. load active orders
  // 5. load products ✅
  // 6. load customers (api wip)

  // --- Validate active shift ---
  // maybe this will be moved to top level and only the validation can be here
  // because open and close shift also can be accessed from dashboard
  const { data: activeShift, isFetching, isSuccess } = useActiveShift();
  const {setActiveShift, defaultFloorId, setDefaultFloor, setProducts} =  useOrderState();
  useEffect(() => {
    if (!isFetching && isSuccess) {
      if (activeShift?.data) setActiveShift(activeShift.data);
      else setBoolState(OpenShiftModalState, false);
    }
  }, [activeShift?.data, isFetching, isSuccess]);

  // --- Shifts Reference ---
  const {data: shifts} = useStoreShifts();
  const {setShifts} =  useStoreState();
  const memoizedShifts = useMemo(() =>
    shifts?.data ?? [], [shifts?.data]);
  useEffect(() => {
    if (!memoizedShifts.length) return;
    setShifts(memoizedShifts);
  }, [memoizedShifts]);

  // --- Categories Reference ---
  const {data: categories} = useCategoryList({ product_status: 'active' })
  const {setCategories} =  useAttributeState();
  const memoizedCategories = useMemo(() =>
    categories?.data ?? [], [categories?.data]);
  useEffect(() => {
    if (!memoizedCategories.length) return;
    setCategories(memoizedCategories);
  }, [memoizedCategories]);

  // --- Floors Reference ---
  const {data: floors} = useFloorList();
  const {setFloors} =  useSeatingState();
  const memoizedFloors = useMemo(() =>
    floors?.data ?? [], [floors?.data]);
  useEffect(() => {
    if (!memoizedFloors.length) return;
    setFloors(memoizedFloors);
    if (defaultFloorId === 0) {
      const validFloors = memoizedFloors.filter(
        f => (f.total_tables ?? 0) > 0);
      if (validFloors.length > 0) {
        const lowestId = Math.min(
          ...validFloors.map(f => f.id));
        setDefaultFloor(lowestId);
      }
    }
  }, [memoizedFloors, defaultFloorId]);

  // --- Products ---
  const {data: products} = useProductList({ status: 'active' });
  const memoizedProducts = useMemo(() =>
    products?.data ?? [], [products?.data]);
  useEffect(() => {
    if (!memoizedProducts.length) return;
    setProducts(memoizedProducts);
  }, [memoizedProducts]);

  return (
    <Fragment>
      <Outlet />

      <OrderListSheet />

      {pathname.startsWith('/orders/') && (
        <>
          <ShiftActionOpenAlertDialog />
          {/* TODO: move close shift to top level so its can be access from everywhere */}
          <ShiftActionCloseAlertDialog />
          <AddToCartModal />
        </>
      )}
    </Fragment>
  )
}

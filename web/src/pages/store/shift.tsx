import {useQuery} from "react-query";
import {HttpResponse} from "@/lib/types/http-response.ts";
import {useStoreShift} from "@/hooks/use-store-shift.tsx";
import {useStorePageState} from "@/stores/store-state.ts";
import {Shift as SType} from "@/lib/types/shift.ts";
import {HandleRequestLoading} from "@/components/request-loading.tsx";
import {NewShiftModal} from "@/pages/store/components/new-shift-modal.tsx";
import {ShiftTable} from "@/pages/store/components/shift-table.tsx";
import {ShiftDetailSheet} from "@/pages/store/components/shift-detail-sheet.tsx";

export const Shift = () => {
  const {get} = useStoreShift();
  const {setShiftList} = useStorePageState();

  const shiftList = useQuery<HttpResponse<SType[]>>(
    "store.shift", get, {
      onSuccess: (resp) => setShiftList(resp.data),
      onError: (error) => console.log(error),
      retry: false,
    });

  if (shiftList.isLoading) {
    return <HandleRequestLoading />
  }

  return (
    <>
      <ShiftTable />
      <NewShiftModal />
      <ShiftDetailSheet />
    </>
  )
}
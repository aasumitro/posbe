import {useStoreSetting} from "@/hooks/use-store-setting.tsx";
import {Store} from "@/lib/types/store.ts";
import {HttpResponse} from "@/lib/types/http-response.ts";
import {useQuery} from "react-query";
import {HandleRequestLoading} from "@/components/request-loading.tsx";
import {useStorePageState} from "@/stores/store-state.ts";
import {StoreContact} from "@/pages/store/components/store-contact.tsx";
import {StoreServiceRate} from "@/pages/store/components/store-service-rate.tsx";
import {StoreFeature} from "@/pages/store/components/store-feature.tsx";
import {StoreAppearance} from "@/pages/store/components/store-appearance.tsx";

export const General = () => {
  const { get } = useStoreSetting();
  const {store, setStoreData} = useStorePageState();

  const setting = useQuery<HttpResponse<Store>>(
    "store.setting", get, {
      onSuccess: (resp) => setStoreData(resp.data),
      onError: (error) =>console.log(error),
      retry: false,
    });

  if (setting.isLoading) {
    return <HandleRequestLoading />
  }

  return (
    <div className="space-y-6">
      <div className="flex flex-col lg:flex-row gap-4">
        <StoreContact store={store}/>

        <StoreServiceRate store={store}/>
      </div>

      <div className="flex flex-col lg:flex-row gap-4">
        <StoreFeature store={store}/>

        <StoreAppearance store={store}/>
      </div>
    </div>
  )
}
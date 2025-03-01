import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import {BrowserRouter, Navigate, Route, Routes} from "react-router-dom";
import {Home} from "@/pages/home.tsx";
import {BackofficeLayout} from "@/layouts/backoffice.tsx";
import {LoginPage} from "@/pages/login.tsx";
import {QueryClient, QueryClientProvider} from "react-query";
import {Sonner} from "@/components/ui/sonner.tsx";
import {StorePage} from "@/pages/store";
import {TransactionPage} from "@/pages/transaction";
import {OldLayoutPage} from "@/pages/layout/old.tsx";
import {CatalogPage} from "@/pages/catalog";
import {TooltipProvider} from "@/components/ui/tooltip.tsx";
import {LayoutBlueprintPage} from "@/pages/layout/ref";
import {StoreLayoutPage} from "@/pages/layout";

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchOnWindowFocus: false,
      cacheTime: 1000 * 60 * 2
    },
  },
});

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <TooltipProvider>
        <BrowserRouter>
          <Routes>
            <Route path="/login" element={<LoginPage/>}/>
            <Route element={<BackofficeLayout />}>
              <Route path="/" element={<Navigate to="home"/>}/>
              <Route path="/home" element={<Home/>}/>
              <Route path="/layouts" element={<StoreLayoutPage/>}/>
              <Route path="/layouts/ref" element={<LayoutBlueprintPage/>}/>
              <Route path="/layouts/old" element={<OldLayoutPage/>}/>
              <Route path="/transactions" element={<TransactionPage/>}/>
              <Route path="/catalogs" element={<CatalogPage />}/>
              <Route path="/store/*" element={<StorePage/>}/>
            </Route>
          </Routes>
        </BrowserRouter>
        <Sonner/>
      </TooltipProvider>
    </QueryClientProvider>
  </React.StrictMode>,
)

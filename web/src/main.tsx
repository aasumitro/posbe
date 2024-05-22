import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import { BrowserRouter, Route, Routes } from "react-router-dom";
import DefaultLayout from "@/layouts/default.tsx";
import { Transaction } from "@/pages/transaction.tsx";
import { Product } from "@/pages/product.tsx";
import { Home } from "@/pages/home.tsx";
import { User } from "@/pages/user.tsx";
import { Setting } from "@/pages/setting.tsx";
import { Map } from "@/pages/map.tsx";

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <BrowserRouter basename="fe">
      <DefaultLayout>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/maps" element={<Map />} />
          <Route path="/transactions" element={<Transaction />} />
          <Route path="/products" element={<Product />} />
          <Route path="/users" element={<User />} />
          <Route path="/settings" element={<Setting />} />
        </Routes>
      </DefaultLayout>
    </BrowserRouter>
  </React.StrictMode>,
)

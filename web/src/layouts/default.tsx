import * as React from 'react'
import { WithLauncher } from '@/components/launcher';
import { useEffect, useState } from 'react';
import { useLocation } from 'react-router-dom';
import { Sidebar, Header } from "@/components/navigation.tsx";
import { Toaster } from "@/components/ui/sonner"

interface Props {
  children: React.ReactNode
}
const DefaultLayout: React.FC<Props> = (props) => {
  const [isLogin, setLogin] = useState(false);
  const location = useLocation();

  useEffect(() => {
      setLogin(true)
  }, [setLogin, location])

  return (
    <>
      <div className=" flex-col md:flex min-h-screen w-full bg-muted/40">
        {(isLogin) && <Sidebar />}
        <div className="flex flex-col sm:gap-4 sm:py-4 sm:pl-14">
            {(isLogin) && <Header />}
            <div className="flex min-h-screen w-full flex-col">
                {props.children}
            </div>
        </div>
        <Toaster/>
      </div>
    </>
  )
}

export default WithLauncher(DefaultLayout);

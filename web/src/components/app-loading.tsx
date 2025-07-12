"use client"

import {useRouterState} from "@tanstack/react-router";
import {cn} from "@/lib/utils";

export const AppLoading = () => {
  const { location } = useRouterState();
  const pathname = location.pathname;

  return (
    <section className={cn(
      "relative flex flex-col items-center justify-center z-999",
      pathname.startsWith('/stores/') ? " w-full h-screen" : "w-screen h-screen"
    )}>
      <div className="loader"></div>
    </section>
  )
}
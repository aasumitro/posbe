import { MonitorXIcon } from "lucide-react";
import { Outlet } from "@tanstack/react-router";
import { Toaster } from "@/components/ui/sonner"

function MainLayout() {
  return (
    <>
      <div className="lg:hidden text-center py-10">
        <MonitorXIcon className="my-8 w-24 h-24 mx-auto"/>
        <h4 className="text-primary text-xl font-bold tracking-tight">
          Screen size not supported.
        </h4>
        <p className="text-secondary-foreground text-md font-normal">
          This application only supports screen sizes of 1024px and above.
        </p>
      </div>

      <Outlet />

      <Toaster position="bottom-center" />
    </>
  )
}

export default MainLayout

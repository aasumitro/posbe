import {StoreNavigation} from "@/pages/store/components/navigation.tsx";
import {Navigate, Route, Routes} from "react-router-dom";
import {General} from "@/pages/store/general.tsx";
import {Shift} from "@/pages/store/shift.tsx";
import {User} from "@/pages/store/user.tsx";
import {Separator} from "@/components/ui/separator.tsx";

export const StorePage = () => {
  return (
    <div className="hidden space-y-6 p-10 pb-16 md:block">
      <div className="space-y-0.5">
        <h2 className="text-2xl font-bold tracking-tight">Settings</h2>
        <p className="text-muted-foreground">
          Manage your store settings and preferences.
        </p>
      </div>
      <Separator className="my-6"/>
      <div className="flex flex-col space-y-8 lg:flex-row lg:space-x-12 lg:space-y-0">
        <aside className="-mx-4 lg:w-1/5">
          <StoreNavigation/>
        </aside>
        <div className="flex-1 ">
          <Routes>
            <Route path="/" element={<Navigate replace to="/store/general"/>}/>
            <Route path="/general" element={<General/>}/>
            <Route path="/shifts" element={<Shift/>}/>
            <Route path="/users" element={<User/>}/>
            <Route path="/*" element={<Navigate replace to="/store/general"/>}/>
          </Routes>
        </div>
      </div>
    </div>
)
}
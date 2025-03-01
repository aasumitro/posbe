import {useNavigate} from "react-router-dom";
import {cn} from "@/lib/utils.ts";
import {isCurrentRoute} from "@/lib/path-validation.ts";
import {buttonVariants} from "@/components/ui/button.tsx";

const sidebarNavItems = [
  {
    title: "General",
    href: "/store/general",
  },
  {
    title: "Shifts",
    href: "/store/shifts",
  },
  {
    title: "Users",
    href: "/store/users",
  },
]

export const StoreNavigation = () => {
  const navigate = useNavigate();

  return (
    <nav className={cn("flex space-x-2 lg:flex-col lg:space-x-0 lg:space-y-1")}>
      {sidebarNavItems.map((item) => (
        <button
          key={item.href}
          onClick={() => navigate(item.href)}
          className={cn(
            buttonVariants({variant: "ghost"}),
            isCurrentRoute(item.href)
              ? "bg-muted hover:bg-muted"
              : "hover:bg-transparent hover:underline",
            "justify-start"
          )}
        >
          {item.title}
        </button>
      ))}
    </nav>
  )
}
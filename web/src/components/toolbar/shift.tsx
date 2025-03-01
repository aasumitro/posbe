import { StoreIcon } from "lucide-react";
import {useEffect, useState} from "react";
import {KEYS} from "@/lib/keys.ts";
import {useNavigate} from "react-router-dom";
import {cn} from "@/lib/utils.ts";
import {isCurrentRoute} from "@/lib/path-validation.ts";

export const ShiftSwitcherNavigation = () => {
  const [currentTime, setCurrentTime] = useState(new Date());
  const [role, setRole] = useState("");
  const navigate = useNavigate();

  useEffect(() => {
    const interval = setInterval(() => {
      setCurrentTime(new Date());
    }, 1000);
    return () => clearInterval(interval);
  }, []);

  useEffect(() => {
    const key = KEYS.UI_USER_PROFILE;
    const profile = localStorage.getItem(key);
    if (profile) {
      const user = JSON.parse(profile);
      setRole(user.role.name);
    }
  }, []);

  const formattedDate = currentTime.toLocaleDateString("en-US", {
    day: "2-digit",
    month: "short",
    year: "numeric",
  });

  const formattedTime = currentTime.toLocaleTimeString("en-US", {
    hour: "2-digit",
    minute: "2-digit",
    second: "2-digit",
    hour12: false,
  });

  const navigateStore = () => {
    if (role !== "admin") return;
    return navigate("/store");
  }

  return (
    <button
      className={cn("p-2 flex items-center space-x-3 rtl:space-x-reverse w-52", {
        "rtl:space-x-reverse md:flex-row md:mt-0 md:border-0 ":
          role === "admin",
        "rounded  border border-gray-100 bg-gray-100": isCurrentRoute("store"),
      })}
      onClick={navigateStore}
    >
      <StoreIcon className="w-6 h-6"/>
      <span className="text-md text-left font-semibold whitespace-nowrap dark:text-white">
          Lorem Store
          <span className="text-xs block">{formattedDate} - {formattedTime}</span>
      </span>
    </button>
  )
}
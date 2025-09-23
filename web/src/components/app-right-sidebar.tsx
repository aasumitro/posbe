import { useEffect, useState } from "react";
import type { ComponentProps } from "react"
import { motion, AnimatePresence } from "framer-motion";
import { Sidebar } from "@/components/ui/sidebar"
import { useRouterState } from "@tanstack/react-router";
import { MenuSidebarSection } from "@/features/menus/components/sidebar";
import { useActionState } from "@/states/action-state";

export const ActionVisibleRightSidebarKey = "action_visible_right_sidebar";

export function AppRightSidebar({
  ...props
}: ComponentProps<typeof Sidebar>) {
  const { bool } = useActionState();
  const { location } = useRouterState();
  const pathname = location.pathname
  const [isVisible, setVisible] = useState(false);

  useEffect(() => {
    setVisible(bool[ActionVisibleRightSidebarKey] ?? false)
  }, [bool]);

  return (
    <AnimatePresence>
      {isVisible && (
        <motion.div
          key="right-sidebar"
          initial={{ x: "100%", opacity: 0 }}
          animate={{
            x: 0, opacity: 1,
            transition: { duration: 0.4, ease: "easeOut" }
          }}
          exit={{
            x: "100%", opacity: 0,
            transition: { duration: 0.4, ease: "easeIn" }
          }}
        >
          {pathname === "/menus" && <MenuSidebarSection {...props} />}
        </motion.div>
      )}
    </AnimatePresence>
  )
}

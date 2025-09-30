"use client"

import { useState } from "react"
import { ChevronUp } from "lucide-react"
import { cn } from "@/lib/utils"
import {TABLE_STATUSES, getLabelBgColorByStatus} from "@/lib/table";
import { AnimatePresence, motion } from "framer-motion"

export default function TableStatusPanel() {
  const [open, setOpen] = useState(false)

  return (
    <div
      className={cn(
        "absolute bottom-2 right-2 select-none",
        "p-4 rounded-lg text-black bg-gray-100 w-52",
      )}
    >
      {/* Header with toggle */}
      <div
        onClick={() => setOpen(!open)}
        className={cn(
          "flex items-center justify-between cursor-pointer",
          open && "border-b pb-2 mb-2"
        )}
        role="button"
        aria-expanded={open}
        tabIndex={0}
      >
        <h4 className="text-sm font-semibold">
          {open ? "Table status" : "Table & chair status"}
        </h4>
        <motion.div
          animate={{ rotate: open ? 180 : 0 }}
          transition={{ duration: 0.25 }}
        >
          <ChevronUp className="w-4 h-4" />
        </motion.div>
      </div>

      {/* Collapsible content */}
      <AnimatePresence>
        {open && (
          <motion.div
            key="table-status-content"
            initial={{ height: 0, opacity: 0 }}
            animate={{ height: "auto", opacity: 1 }}
            exit={{ height: 0, opacity: 0 }}
            transition={{ duration: 0.25, ease: "easeInOut" }}
            className="overflow-hidden"
          >
            <div className="space-y-2 mb-4 mt-2">
              {TABLE_STATUSES.map((status, index) => (
                <div
                  className="flex flex-row gap-2 items-center text-xs"
                  key={index}
                >
                  <div
                    className="w-4 h-4 rounded-sm"
                    style={{ background: getLabelBgColorByStatus(status) }}
                  />
                  {status.replace("-", " ").toUpperCase()}
                </div>
              ))}
            </div>

            <h4 className="text-sm font-semibold border-y py-2 mb-2">Chair status</h4>
            <div className="space-y-2">
              {[
                { color: "#7a411c", title: "VACANT" },
                { color: "#593d35", title: "OCCUPIED" },
              ].map((status, index) => (
                <div
                  className="flex flex-row gap-2 items-center text-xs"
                  key={index}
                >
                  <div
                    className="w-4 h-4 rounded-sm"
                    style={{ background: status.color }}
                  />
                  {status.title}
                </div>
              ))}
            </div>
          </motion.div>
        )}
      </AnimatePresence>
    </div>
  )
}

import {cn} from "@/lib/utils";
import {ChevronDown} from "lucide-react";
import {useOrderState} from "@/states/order-state";
import {useSeatingState} from "@/states/seating-state";
import {AnimatePresence, motion} from "framer-motion";
import {useState} from "react";

export function FloorPanel() {
  const [open, setOpen] = useState(false)
  const {defaultFloorId, setDefaultFloor} =  useOrderState();
  const {floors} =  useSeatingState();
  const selectedFloor = floors?.find(
    (f) => f.id === defaultFloorId)

  return (
    <div
      className={cn(
        "absolute top-4 left-4 select-none py-2 px-4 w-44 cursor-pointer",
        "rounded-lg border bg-background shadow-xs hover:bg-accent hover:text-accent-foreground",
        open && "bg-accent",
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
          {open ? "Floor List" : `Floor: ${selectedFloor?.name}`}
        </h4>
        <motion.div
          animate={{ rotate: open ? 180 : 0 }}
          transition={{ duration: 0.25 }}
        >
          <ChevronDown className="w-4 h-4" />
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
            <div className="space-y-2 mt-2">
              {floors?.map((floor) => (
                <div
                  className={cn(
                    "relative flex flex-row gap-2 items-center text-xs ",
                    "hover:bg-black/50 hover:text-white cursor-pointer rounded-md p-2",
                    selectedFloor?.id === floor.id && "bg-black/50  text-white"
                  )}
                  key={floor.id}
                  onClick={(e) => {
                    e.preventDefault();
                    setDefaultFloor(floor.id);
                    setOpen(false);
                  }}
                >
                  {floor.name} {floor?.total_tables ? `(${floor.total_tables})` : ''}
                </div>
              ))}
            </div>
          </motion.div>
        )}
      </AnimatePresence>
    </div>
  )
}
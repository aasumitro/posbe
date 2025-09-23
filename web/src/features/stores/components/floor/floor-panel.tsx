import {cn} from "@/lib/utils";
import {ChevronUp} from "lucide-react";
import {useState} from "react";
import {AnimatePresence, motion} from "framer-motion";
import {useSeatingState} from "@/states/seating-state";
import {FloorActionAdd} from "@/features/stores/components/floor/floor-action-new";
import {FloorSheet} from "@/features/stores/components/floor/floor-sheet";

export function FloorPanel() {
  const [open, setOpen] = useState(false)
  const {selectedFloor, floors, setSelectedFloor} =  useSeatingState();

  return (
    <div
      className={cn(
        "absolute bottom-4 right-4 select-none",
        "py-3 px-4 rounded-lg text-black bg-white border dark:bg-input/30 dark:border-input w-56",
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
          {open ? "Floor List" : selectedFloor?.name}
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
                    setSelectedFloor(floor);
                    setOpen(false);
                  }}
                >
                  {floor.name} {floor?.total_tables ? `(${floor.total_tables})` : ''}
                </div>
              ))}
            </div>

            <h4 className="text-sm font-semibold border-t border-b py-2 my-4">
              Floor Actions
            </h4>

           <div className="space-y-2">
             <FloorActionAdd />
             <FloorSheet />
           </div>
          </motion.div>
        )}
      </AnimatePresence>
    </div>
  )
}
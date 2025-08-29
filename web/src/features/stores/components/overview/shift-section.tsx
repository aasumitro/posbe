import {Table, TableBody, TableCell, TableFooter, TableHead, TableHeader, TableRow} from "@/components/ui/table";
import {IconEye} from "@tabler/icons-react";
import {Button} from "@/components/ui/button";
import {ShiftActionAdd} from "@/features/stores/components/overview/shift-action-add";
import {useStoreState} from "@/states/store-state";
import {intToTime} from "@/lib/time";
import {ShiftDetailActionSheet, ShiftDetailActionSheetState} from "@/features/stores/components/overview/shift-detail-sheet";
import {useActionState} from "@/states/action-state";

export function ShiftTableSection() {
  const {shifts, setSelectedShifts} = useStoreState();
  const { setBoolState } = useActionState();

  return (
    <div className="overflow-hidden border rounded-lg">
      <Table className="w-full">
        <TableHeader className="bg-gray-100 dark:bg-input/30">
          <TableRow className="h-14">
            <TableHead className="w-[100px]"></TableHead>
            <TableHead>Name</TableHead>
            <TableHead>Work time</TableHead>
            <TableHead>Total Used</TableHead>
            <TableHead>Last Action (Open & Close)</TableHead>
            <TableHead className="w-[100px]"></TableHead>
          </TableRow>
        </TableHeader>

        <TableBody>
          {shifts?.map((shift, index) => (
            <TableRow key={index}>
              <TableCell />

              <TableCell>
                {shift.name.toUpperCase()}
              </TableCell>
              <TableCell className="font-light text-xs">
                {shift.start_time && intToTime(shift.start_time)} -
                {shift.end_time && intToTime(shift.end_time)}
              </TableCell>
              <TableCell className="font-light text-xs text-center">{shift.total_usage}</TableCell>

              {!shift?.active_shift ? (
                <TableCell className="font-light text-xs text-center">no data</TableCell>
              ) : (
                <TableCell className="font-light text-xs flex flex-col">
                  <span>29 Jan 2025</span>
                  <span>08:30 AM (O) - 10:30 PM (C)</span>
                </TableCell>
              )}

              <TableCell className="space-x-2">
                <Button
                  variant="outline"
                  className="p-1 h-8 w-8 cursor-pointer"
                  onClick={(e) => {
                    e.preventDefault();
                    setSelectedShifts(shift);
                    setBoolState(ShiftDetailActionSheetState, true);
                  }}
                >
                  <IconEye className="w-4 h-4 text-gray-500" />
                </Button>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>

        <TableFooter className="bg-gray-100 hover:bg-gray-100/50 dark:bg-input/30">
          <TableRow>
            <TableCell colSpan={6} className="py-4 text-xs font-light pl-4 select-none cursor-pointer">
              <ShiftActionAdd />
            </TableCell>
          </TableRow>
        </TableFooter>
      </Table>

      <ShiftDetailActionSheet />
    </div>
  )
}
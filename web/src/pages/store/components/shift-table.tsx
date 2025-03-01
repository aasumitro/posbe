import {Card, CardContent, CardDescription, CardHeader, CardTitle} from "@/components/ui/card.tsx";
import {Button} from "@/components/ui/button.tsx";
import {MoreHorizontal, PlusIcon} from "lucide-react";
import {Table, TableBody, TableCell, TableHead, TableHeader, TableRow} from "@/components/ui/table.tsx";
import {capitalize} from "@/lib/str.ts";
import {intToTime, sqlTimeToFormattedTime} from "@/lib/time.ts";
import {
  DropdownMenu,
  DropdownMenuContent, DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuTrigger
} from "@/components/ui/dropdown-menu.tsx";
import {KEYS} from "@/lib/keys.ts";
import {useGlobalStateStore} from "@/stores/global-state.ts";
import {useStorePageState} from "@/stores/store-state.ts";

export function ShiftTable() {
  const {shifts} = useStorePageState();

  const {setBoolState} = useGlobalStateStore();

  const openNewShiftModal = () =>
    setBoolState(KEYS.DISPLAY_NEW_SHIFT_MODAL, true)

  const openShiftDetailSheet = () =>
    setBoolState(KEYS.DISPLAY_SHIFT_DETAIL_SHEET, true)

  return (
    <Card>
      <div className="flex justify-between items-center">
        <CardHeader>
          <CardTitle>Shifts</CardTitle>
          <CardDescription>
            Manage your store shifts.
          </CardDescription>
        </CardHeader>
        <Button type="submit" className="mr-6" onClick={openNewShiftModal} >
          <PlusIcon className="h-4 w-4 mr-2" />Add new Shift
        </Button>
      </div>
      <CardContent>
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Name</TableHead>
              <TableHead className="hidden md:table-cell">
                Work time
              </TableHead>
              <TableHead className="hidden md:table-cell">
                Total Orders, & Usages
              </TableHead>
              <TableHead className="hidden md:table-cell">
                Open At (Last Action)
              </TableHead>
              <TableHead className="hidden md:table-cell">
                Close At (Last Action)
              </TableHead>
              <TableHead>
                <span className="sr-only">Actions</span>
              </TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {shifts && shifts?.map((shift) => (
              <TableRow key={shift.id}>
                <TableCell className="font-medium">
                  {capitalize(shift.name)}
                </TableCell>
                <TableCell className="hidden md:table-cell">
                  {intToTime(shift.start_time)} - {intToTime(shift.end_time)}
                </TableCell>
                <TableCell className="hidden md:table-cell">
                  {shift.total_transaction}x | {shift.total_usage}x
                </TableCell>
                <TableCell className="hidden md:table-cell">
                  {shift.current_shift ?
                    sqlTimeToFormattedTime(shift.current_shift.open_at)
                    : "-"}
                </TableCell>
                <TableCell className="hidden md:table-cell">
                  {shift.current_shift ?
                    sqlTimeToFormattedTime(shift.current_shift.close_at)
                    : "-"}
                </TableCell>
                <TableCell>
                  <DropdownMenu>
                    <DropdownMenuTrigger asChild>
                      <Button
                        aria-haspopup="true"
                        size="icon"
                        variant="ghost"
                      >
                        <MoreHorizontal className="h-4 w-4"/>
                        <span className="sr-only">Toggle menu</span>
                      </Button>
                    </DropdownMenuTrigger>
                    <DropdownMenuContent align="end">
                      <DropdownMenuLabel>Actions</DropdownMenuLabel>
                      <DropdownMenuItem onClick={openShiftDetailSheet}>Detail</DropdownMenuItem>
                      <DropdownMenuItem>Delete</DropdownMenuItem>
                    </DropdownMenuContent>
                  </DropdownMenu>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </CardContent>
    </Card>
  )
}
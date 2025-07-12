import {Table, TableBody, TableCell, TableFooter, TableHead, TableHeader, TableRow} from "@/components/ui/table";
import {IconEye, IconPlus} from "@tabler/icons-react";
import {Button} from "@/components/ui/button";
import {useEffect, useState} from "react";

export function ShiftTableSection() {
  const [total, setTotal] = useState(0)

  useEffect(() => setTotal(2), []);

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
          {Array.from({ length: total }).map((_, i) => (
            <TableRow key={i}>
              <TableCell />

              <TableCell>
                Shift {i+1}
              </TableCell>
              <TableCell className="font-light text-xs">
                {i === 0 ? "08:00AM - 03:00PM" : "03:01PM - 10:00PM"}
              </TableCell>
              <TableCell className="font-light text-xs text-center">{(i+1) * 2}</TableCell>
              <TableCell className="font-light text-xs flex flex-col">
                <span>29 Jan 2025</span>
                <span>08:30 AM (O) - 10:30 PM (C)</span>
              </TableCell>

              <TableCell className="space-x-2">
                <Button variant="outline" className="p-1 h-8 w-8">
                  <IconEye className="w-4 h-4 text-gray-500" />
                </Button>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>

        <TableFooter className="bg-gray-100 hover:bg-gray-100/50 dark:bg-input/30">
          <TableRow>
            <TableCell colSpan={6} className="py-4 text-xs font-light pl-4 select-none cursor-pointer">
              <div className="flex items-center justify-center gap-1">
                <IconPlus className="w-4 h-4" />
                Add new Shift
              </div>
            </TableCell>
          </TableRow>
        </TableFooter>
      </Table>
    </div>
  )
}
import {Card, CardContent, CardDescription, CardHeader, CardTitle} from "@/components/ui/card.tsx";
import {Table, TableBody, TableCell, TableHead, TableHeader, TableRow} from "@/components/ui/table.tsx";
import {
  DropdownMenu,
  DropdownMenuContent, DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuTrigger
} from "@/components/ui/dropdown-menu.tsx";
import {Button} from "@/components/ui/button.tsx";
import {MoreHorizontal} from "lucide-react";
import {Badge} from "@/components/ui/badge.tsx";

export const Shift = () => {
  return (
    <Card>
      <CardHeader>
        <CardTitle>Shifts</CardTitle>
        <CardDescription>
          Manage your store shifts.
        </CardDescription>
      </CardHeader>
      <CardContent>
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Name</TableHead>
              <TableHead className="hidden md:table-cell">
                Status
              </TableHead>
              <TableHead className="hidden md:table-cell">
                Total Transaction & Usages
              </TableHead>
              <TableHead className="hidden md:table-cell">
                Work time
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
            <TableRow>
              <TableCell className="font-medium">
                Shift 1
              </TableCell>
              <TableCell className="hidden md:table-cell">
                <Badge variant="outline">
                  Active
                </Badge>
              </TableCell>
              <TableCell className="hidden md:table-cell">
                25 | 5 (last) - 250x
              </TableCell>
              <TableCell className="hidden md:table-cell">
                08:00 - 14:00
              </TableCell>
              <TableCell className="hidden md:table-cell">
                2023-07-12 10:42 AM
              </TableCell>
              <TableCell className="hidden md:table-cell">
                -
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
                    <DropdownMenuItem>Detail</DropdownMenuItem>
                    <DropdownMenuItem>Edit</DropdownMenuItem>
                  </DropdownMenuContent>
                </DropdownMenu>
              </TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </CardContent>
    </Card>
  )
}
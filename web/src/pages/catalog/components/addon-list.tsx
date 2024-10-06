import {Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle} from "@/components/ui/card.tsx";
import {Button} from "@/components/ui/button.tsx";
import {MoreHorizontal, PlusCircleIcon} from "lucide-react";
import {Table, TableBody, TableCell, TableHead, TableHeader, TableRow} from "@/components/ui/table.tsx";
import {Badge} from "@/components/ui/badge.tsx";
import {
  DropdownMenu,
  DropdownMenuContent, DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuTrigger
} from "@/components/ui/dropdown-menu.tsx";

export const AddonList = () => {
  return (
    <Card className="w-full xl:w-2/6">
      <CardHeader className="flex flex-row items-center justify-between">
        <div>
          <CardTitle>Addons</CardTitle>
          <CardDescription>
            Manage your product addons.
          </CardDescription>
        </div>
        <Button variant="outline">
          <PlusCircleIcon className="w-4 h-4"/>
        </Button>
      </CardHeader>
      <CardContent>
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Name</TableHead>
              <TableHead className="hidden md:table-cell">
                Price
              </TableHead>
              <TableHead className="hidden md:table-cell">
                Transactions
              </TableHead>
              <TableHead className="hidden md:table-cell">
                Status
              </TableHead>
              <TableHead>
                <span className="sr-only">Actions</span>
              </TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow>
              <TableCell className="font-medium">
                Extra Sweet
              </TableCell>
              <TableCell className="hidden md:table-cell">
                $10
              </TableCell>
              <TableCell className="hidden md:table-cell">
                250x
              </TableCell>
              <TableCell className="hidden md:table-cell">
                <Badge variant="outline">
                  Active
                </Badge>
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
      <CardFooter>
        <div className="text-xs text-muted-foreground">
          Showing <strong>1</strong> of <strong>32</strong>
          {" "} Addons
        </div>
      </CardFooter>
    </Card>
  )
}
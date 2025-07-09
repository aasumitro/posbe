import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
} from "@/components/ui/sheet"
import {useActionState} from "@/states/action-state";
import {useEffect, useState} from "react";
import {Button} from "@/components/ui/button";
import {Label} from "@/components/ui/label";
import {Input} from "@/components/ui/input";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select";

export const TeamDetailActionSheetState = "team_detail_action_sheet_state"

export function TeamDetailActionSheet() {
  const [open, setOpen] = useState(false);
  const { bool, setBoolState } = useActionState();

  useEffect(() => {
    if (bool[TeamDetailActionSheetState]){
      setOpen(bool[TeamDetailActionSheetState])
    }
  }, [bool]);

  function onOpenChange(newOpen: boolean) {
    setOpen(newOpen);
    if (!newOpen) {
      setBoolState(TeamDetailActionSheetState, false);
    }
  }

  return (
    <Sheet open={open} onOpenChange={onOpenChange}>
      <SheetContent>
        <SheetHeader>
          <SheetTitle>Team Member Info</SheetTitle>
          <SheetDescription>
            Preview team member details or remove them from the team.
          </SheetDescription>
        </SheetHeader>

        <div className="px-4 w-full space-y-2">
          <Label>Username</Label>
          <Input id="username" value="@zeros" disabled />
        </div>

        <div className="px-4 w-full space-y-2">
          <Label>Name</Label>
          <Input id="name" value="Zeros Mardigu" disabled />
        </div>

        <div className="px-4 w-full space-y-2">
          <Label>Email</Label>
          <Input id="email" value="zeros@member.pos" disabled />
        </div>

        <div className="px-4 w-full space-y-2">
          <Label>Role</Label>
          <Select>
            <SelectTrigger className="w-full h-10">
              <SelectValue placeholder="Select a role"/>
            </SelectTrigger>
            <SelectContent>
              {[
                {id: 1, name: "admin"},
                {id: 2, name: "cashier"},
                {id: 3, name: "waiter"},
              ].map((role) => (
                <SelectItem
                  key={role.id}
                  value={`${role.id}`}
                >
                  {role.name}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
        </div>

        <div className="flex justify-between items-center px-4 mt-2">
          <Button
            variant="link"
            className="text-red-500"
          >
            DELETE
          </Button>

          <div className="space-x-2">
            <Button variant="outline">Cancel</Button>
            <Button>Save</Button>
          </div>
        </div>
      </SheetContent>
    </Sheet>
  )
}
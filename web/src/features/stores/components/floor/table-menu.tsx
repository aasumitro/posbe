import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem, DropdownMenuPortal, DropdownMenuSeparator, DropdownMenuSub,
  DropdownMenuSubContent, DropdownMenuSubTrigger,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { cn } from "@/lib/utils"
import {IconAdjustmentsCog, IconCheck} from "@tabler/icons-react";
import {useActionState} from "@/states/action-state";
import {useSeatingState} from "@/states/seating-state";
import {TableActionDeleteModalState} from "@/features/stores/components/floor/table-action-delete";
import {TableActionUpdateModalState} from "@/features/stores/components/floor/table-action-update";
import {useUpdateTable} from "@/hooks/use-seating";
import {useQueryClient} from "@tanstack/react-query";
import {isHTTPResponse} from "@/lib/api";
import {toast} from "sonner";

type tableErrorResponse = {
  floor_id?: string[]
}

export function TableMenu({
  id,
  fid,
  shape,
  chairs,
}: {
  id: number
  fid: number
  shape: string
  chairs: number
}) {
  const { setBoolState } = useActionState();
  const {setSelectedTableId, floors} =  useSeatingState();
  const { mutate: updateTable} = useUpdateTable();
  const queryClient = useQueryClient();

  const openUpdateModal = (id: number) => {
    setSelectedTableId(id);
    setBoolState(TableActionUpdateModalState, true);
  }

  const onTransferFloor = (id: number, prevId: number, nextId: number) => {
    updateTable({
      id: id, body: JSON.stringify({ floor_id: nextId}),
    }, {
      onSuccess: async () => {
        await queryClient.invalidateQueries({ queryKey: ["floors"] });
        await queryClient.invalidateQueries({ queryKey: ["floor.tables", prevId] });
        await queryClient.invalidateQueries({ queryKey: ["floor.tables", nextId] });
      },
      onError: (error) => {
        if (error && isHTTPResponse<null>(error)) {
          if (typeof error.data === "string") {
            toast.error(error.data);
            return;
          }

          if (typeof error.data === "object" && error.data !== null) {
            const data = error.data as tableErrorResponse;

            if (data.floor_id && data.floor_id.length > 0) {
              toast.error(data.floor_id[0]);
            }
          }
        }

        if (error instanceof Error) {
          const clientError = error as Error
          toast.error(clientError.message);
        }
      }
    });
  }

  const confirmDelete = (id: number) => {
    setSelectedTableId(id);
    setBoolState(TableActionDeleteModalState, true);
  }

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <div
          className={cn(
            "w-full h-full",
            shape === "circle" && "p-[2px]"
          )}
        >
          <div
            className={cn(
              "w-full h-full flex justify-center items-center text-xs",
              "cursor-pointer transition-colors hover:brightness-90",
              "bg-gray-50/10 hover:bg-gray-100/50",
              shape === "circle" ? "rounded-full": "rounded-sm",
              shape === "rectangle" && chairs <= 2 && "-rotate-90"
            )}
          >
            <IconAdjustmentsCog className="w-4 h-4 text-white" />
          </div>
        </div>
      </DropdownMenuTrigger>
      <DropdownMenuContent className="w-56" align="start">
        <DropdownMenuGroup>
          <DropdownMenuItem
            className="cursor-pointer"
            onClick={(e) => {
              e.preventDefault();
              openUpdateModal(id);
            }}
          >
            Edit
          </DropdownMenuItem>

          {floors && floors?.length > 1 && (
            <>
              <DropdownMenuSeparator />

              <DropdownMenuSub>
                <DropdownMenuSubTrigger>
                  Move to Floor
                </DropdownMenuSubTrigger>
                <DropdownMenuPortal>
                  <DropdownMenuSubContent>
                    {floors.map((f) => (
                      <DropdownMenuItem
                        key={f?.id}
                        className="flex justify-between items-center cursor-pointer"
                        onClick={(e) => {
                          e.preventDefault();
                          if (!f) return;
                          onTransferFloor(id, fid, f?.id)
                        }}
                        disabled={f?.id == fid}
                      >
                        {f?.name}
                        {f?.id === fid && <IconCheck />}
                      </DropdownMenuItem>
                    ))}
                  </DropdownMenuSubContent>
                </DropdownMenuPortal>
              </DropdownMenuSub>

              <DropdownMenuSeparator />
            </>
          )}

          <DropdownMenuItem
            className="cursor-pointer"
            variant="destructive"
            onClick={(e) => {
              e.preventDefault();
              confirmDelete(id);
            }}
          >
            Delete
          </DropdownMenuItem>
        </DropdownMenuGroup>
      </DropdownMenuContent>
    </DropdownMenu>
  )
}

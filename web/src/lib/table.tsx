import type {TableStatus} from "@/components/table";

export const TABLE_STATUSES: TableStatus[] = [
  "available",
  "occupied",
  "reserved",
  "needs-cleaning"
]

export function getLabelBgColorByStatus (status: TableStatus): string {
  switch (status) {
    case "available":
      return "#7a411c";
    case "occupied":
      return "#FFF176";
    case "reserved":
      return "#4FC3F7";
    case "needs-cleaning":
      return "#EF9A9A";
    default:
      return "#7a411c"
  }
}
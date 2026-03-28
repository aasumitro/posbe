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

export function getTableScale(shape: string, chairs: number): string {
  const circleScales: Record<number, string> = {
    1: "scale(0.8)",
    2: "scale(1.1)",
    3: "scale(0.85)",
  };
  const rectangleScales: Record<number, string> = {
    1: "scale(0.5)",
    2: "scale(0.5)",
    3: "scale(0.8)",
    4: "scale(1)",
    5: "scale(0.9)",
    6: "scale(1.2)",
    7: "scale(1.1)",
    8: "scale(1.4)",
    9: "scale(1.3)",
    12: "scale(1.8)",
  };

  if (shape === "circle") {
    return circleScales[chairs] ?? "scale(1)";
  } else {
    return rectangleScales[chairs] ?? "scale(1.5)";
  }
}

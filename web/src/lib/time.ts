import type {Int64R} from "@/types/http-response";
import type {TimeValue} from "react-aria";
import { Time as ITime } from "@internationalized/date";  // add this import

export function formatTimestamp(unixTimestamp: number) {
  const date = new Date(unixTimestamp * 1000);
  const options: Intl.DateTimeFormatOptions = {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: 'numeric',
    minute: '2-digit',
    hour12: true,
  };
  const formattedDate = date.toLocaleString('en-US', options);
  const [month, day, year] = formattedDate.split(',')[0].split('/');
  const time = formattedDate.split(', ')[1];
  return `${year}-${month}-${day} ${time}`;
}

export const intToTime = (timeInt: number): string => {
  if (!Number.isInteger(timeInt)) {
    return "";
  }
  const hour = Math.floor(timeInt / 100);
  const minute = timeInt % 100;
  const period = hour < 12 ? "AM" : "PM";
  const formattedHour = hour % 12 || 12;
  const formattedHourString = formattedHour.toString().padStart(2, "0");
  const formattedMinuteString = minute.toString().padStart(2, "0");
  return `${formattedHourString}:${formattedMinuteString}${period}`;
};

export const strTimeToInt = (timeStr: string): number => {
  const cleanTimeString = timeStr.replace(/(am|pm)/i, "").trim();
  const [hours, minutes] = cleanTimeString.split(":").map(Number);
  let convertedHours = hours;
  if (timeStr.toLowerCase().includes("pm") && hours !== 12) {
    convertedHours = hours + 12;
  }
  if (timeStr.toLowerCase().includes("am") && hours === 12) {
    convertedHours = 0; // Midnight case
  }
  return convertedHours * 100 + minutes;
};

export function sqlTimeToFormattedTime(time: Int64R) {
  if (!time.Valid) {
    return "-"
  }
  return formatTimestamp(time.Int64)
}

export function intToTimeValue(intVal?: number | null): TimeValue | null {
  if (intVal === undefined || intVal === null) return null;
  const hh = Math.floor(intVal / 100);
  const mm = intVal % 100;
  return new ITime(hh, mm); 
}
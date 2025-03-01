import {Time} from "@/lib/types/common.ts";

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

export function sqlTimeToFormattedTime(time: Time) {
  if (!time.Valid) {
    return "-"
  }
  return formatTimestamp(time.Int64)
}
import type {Time} from "@/types/http-response";

export const timeReference: string[] = Array.from({ length: 24 * 60 / 5 }, (_, i) => {
  const totalMinutes = i * 5;
  const hours = Math.floor(totalMinutes / 60);
  const minutes = totalMinutes % 60;
  const hour12 = hours % 12 === 0 ? 12 : hours % 12;
  const ampm = hours < 12 ? "AM" : "PM";
  const hourStr = hour12.toString().padStart(2, "0");
  const minuteStr = minutes.toString().padStart(2, "0");
  return `${hourStr}:${minuteStr}${ampm}`;
});

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

export function sqlTimeToFormattedTime(time: Time) {
  if (!time.Valid) {
    return "-"
  }
  return formatTimestamp(time.Int64)
}
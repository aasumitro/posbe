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

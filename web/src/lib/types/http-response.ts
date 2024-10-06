export interface HttpResponse<T> {
  data: T | null;
  status: string;
  code: number;
}
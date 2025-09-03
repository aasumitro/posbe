export interface HTTPResponse<T> {
  data: T | null;
  status: string;
  code: number;
}

export interface ResponseStatus {
  request_id: string;
  error: boolean;
  message: string;
  details: string | string[];
}

export interface ResponsePagination {
  limit: number;
  offset: number;
  current_page: number;
  total_pages: number;
  total_items: number;
}

export interface Int64R {
  Int64: number
  Valid: boolean
}
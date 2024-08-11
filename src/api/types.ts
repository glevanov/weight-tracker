export type Weight = {
  weight: number;
  timestamp: string;
};

export type SuccessResponse<Data> = {
  isSuccess: true;
  data: Data;
};

export type ErrorResponse = {
  isSuccess: false;
  error: string;
};

export type Response<Data> = SuccessResponse<Data> | ErrorResponse;

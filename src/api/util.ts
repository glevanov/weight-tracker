import type { Range } from "../visualization/types";
import type { ErrorResponse } from "./types";

const DAY = 24 * 60 * 60 * 1000;

export const mapRangeToDates = (range: Range) => {
  switch (range) {
    case "14-days":
      return {
        start: new Date(Date.now() - 14 * DAY),
        end: new Date(),
      };
    case "30-days":
      return {
        start: new Date(Date.now() - 30 * DAY),
        end: new Date(),
      };
    case "90-days":
      return {
        start: new Date(Date.now() - 90 * DAY),
        end: new Date(),
      };
    case "365-days":
      return {
        start: new Date(Date.now() - 365 * DAY),
        end: new Date(),
      };
    case "all-time":
      return {
        start: new Date(0),
        end: new Date(),
      };
  }
};

export const extractError = async (
  response: globalThis.Response,
): Promise<ErrorResponse> => {
  const error = (await response.text()) ?? "Произошла неизвестная ошибка";
  return { isSuccess: false, error };
};

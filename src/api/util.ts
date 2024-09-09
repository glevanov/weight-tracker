import { get } from "svelte/store";

import type { Range } from "../screens/chart/types";
import type { Response } from "./types";
import { switchScreen } from "../store/screen";
import { language } from "../store/language";

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

const validateResponse = <Data>(
  response: unknown,
): response is Response<Data> => {
  if (typeof response !== "object" || response === null) {
    return false;
  }

  if (!("isSuccess" in response) || typeof response.isSuccess !== "boolean") {
    return false;
  }

  if (response.isSuccess) {
    return "data" in response;
  } else {
    return "error" in response;
  }
};

export const extractResult = async <Data>(
  response: globalThis.Response,
): Promise<Response<Data>> => {
  try {
    const result = await response.json();
    const isValid = validateResponse<Data>(result);
    if (!isValid) {
      throw new Error("Invalid response");
    }

    return result;
  } catch {
    return { isSuccess: false, error: "Произошла неизвестная ошибка" };
  }
};

export const handleAuthError = (response: globalThis.Response) => {
  if (response.status === 401) {
    switchScreen("login");
  }
};

export const getAcceptLanguage = () => get(language);

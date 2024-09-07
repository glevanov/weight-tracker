import type { Range } from "../screens/chart/types";
import type { Response, Weight } from "./types";
import { extractResult, handleAuthError, mapRangeToDates } from "./util";

const apiUrl = import.meta.env.VITE_API_URL;

export const getWeights = async (range: Range): Promise<Response<Weight[]>> => {
  const { start, end } = mapRangeToDates(range);
  const response = await fetch(
    `${apiUrl}/weights?start=${start.toISOString()}&end=${end.toISOString()}`,
    {
      method: "GET",
      credentials: "include",
    },
  );

  handleAuthError(response);

  return await extractResult(response);
};

export const addWeight = async (weight: string): Promise<Response<string>> => {
  const response = await fetch(`${apiUrl}/weights`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ weight }),
    credentials: "include",
  });

  handleAuthError(response);

  return await extractResult(response);
};

export const login = async (
  username: string,
  password: string,
): Promise<Response<string>> => {
  const response = await fetch(`${apiUrl}/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ username, password }),
    credentials: "include",
  });

  return await extractResult(response);
};

export const checkSession = async () => {
  const response = await fetch(`${apiUrl}/session-check`, {
    method: "GET",
    credentials: "include",
  });

  return response.status === 200;
};

const RETRIES = 10;
const TIMEOUT = 50000;

export const checkHealth = async (): Promise<Response<string>> => {
  for (let attempt = 1; attempt <= RETRIES; attempt++) {
    const controller = new AbortController();
    const signal = controller.signal;
    const timeoutId = setTimeout(() => controller.abort(), TIMEOUT);

    try {
      const response = await fetch(`${apiUrl}/health-check`, {
        method: "GET",
        signal,
      });

      clearTimeout(timeoutId);

      if (response.status === 200) {
        return {
          isSuccess: true,
          data: "Сервис доступен",
        };
      }
    } catch {
      clearTimeout(timeoutId);
    }
  }
  return {
    isSuccess: false,
    error: "Сервис недоступен",
  };
};

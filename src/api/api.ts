import type { Range } from "../screens/chart/types";
import type { Response, Weight } from "./types";
import { extractError, handleAuthError, mapRangeToDates } from "./util";

// const apiUrl = "http://localhost:3000";
const apiUrl = "https://weight-tracker-service.onrender.com";

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

  if (response.status !== 200) {
    return await extractError(response);
  }

  const weights = await response.json();

  return {
    isSuccess: true,
    data: weights,
  };
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

  if (response.status !== 201) {
    return await extractError(response);
  }

  return {
    isSuccess: true,
    data: (await response.text()) ?? "Вес успешно добавлен",
  };
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

  if (response.status !== 200) {
    return await extractError(response);
  }

  return {
    isSuccess: true,
    data: (await response.text()) ?? "Вход выполнен",
  };
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

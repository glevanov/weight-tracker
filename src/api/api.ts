import type { Range } from "../screens/chart/types";
import type { Response, Weight } from "./types";
import { extractError, handleAuthError, mapRangeToDates } from "./util";

const apiUrl = "http://localhost:3000";

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

  if (!response.ok) {
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

  if (!response.ok) {
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

  if (!response.ok) {
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

export const checkHealth = async () => {
  await fetch(`${apiUrl}/health-check`, {
    method: "GET",
  });
};

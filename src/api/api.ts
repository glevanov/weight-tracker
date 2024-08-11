import type { Range } from "../visualization/types";
import type { Response, Weight } from "./types";
import { extractError, mapRangeToDates } from "./util";

const apiUrl = "http://localhost:3000";

export const getWeights = async (range: Range): Promise<Response<Weight[]>> => {
  const { start, end } = mapRangeToDates(range);
  const response = await fetch(
    `${apiUrl}/weights?start=${start.toISOString()}&end=${end.toISOString()}`,
  );

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
  });

  if (!response.ok) {
    return await extractError(response);
  }

  return {
    isSuccess: true,
    data: (await response.text()) ?? "Вес успешно добавлен",
  };
};

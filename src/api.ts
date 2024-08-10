import type { Range } from "./visualization/types";

const apiUrl = "http://localhost:3000";

const DAY = 24 * 60 * 60 * 1000;

const mapRangeToDates = (range: Range) => {
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

export const getWeights = async (range: Range) => {
  const { start, end } = mapRangeToDates(range);
  const result = await fetch(
    `${apiUrl}/weights?start=${start.toISOString()}&end=${end.toISOString()}`,
  );
  return result.json();
};

export const addWeight = async (weight: number) => {
  const result = await fetch(`${apiUrl}/weights`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ weight }),
  });
  return result.json();
};

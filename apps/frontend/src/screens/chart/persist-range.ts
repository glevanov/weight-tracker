import type { Range } from "./types";

const KEY = "weight-tracker:range";

export const readRange = () => {
  return localStorage.getItem(KEY) as Range | null;
};

export const storeRange = (range: Range) => {
  localStorage.setItem(KEY, range);
};

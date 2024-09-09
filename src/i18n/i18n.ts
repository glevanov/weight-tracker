import { ru } from "./locales/ru";
import { en } from "./locales/en";
import { sv } from "./locales/sv";
import { zhTW } from "./locales/zh-tw";

export type Lang = "ru" | "en" | "sv" | "zh-tw";

/**
 * Generic type to get a path of a nested object
 */
type Path<T> = T extends object
  ? {
      [K in keyof T]: K extends string
        ? T[K] extends object
          ? `${K}` | `${K}.${Path<T[K]>}`
          : `${K}`
        : never;
    }[keyof T]
  : never;

const traverse = (obj: Locale, path: Path<Locale>): string => {
  const parts = path.split(".");

  let value: Locale | Locale[keyof Locale] | string = obj;

  for (const part of parts) {
    if (typeof value !== "string") {
      // @ts-expect-error difficult to type
      value = value[part];
    }
  }

  return value as string;
};

export type Locale = typeof ru;

const locales: Record<Lang, Locale> = {
  ru,
  en,
  sv,
  "zh-tw": zhTW,
};

export const i18n = (lang: Lang, path: Path<Locale>): string => {
  return traverse(locales[lang], path);
};

import { ru } from "./locales/ru";
import { en } from "./locales/en";
import { sv } from "./locales/sv";

export type Lang = "ru" | "en" | "sv";

/**
 * Generic type to get a path of a nested object
 */
type Path<Type> = Type extends object
  ? {
      [Key in keyof Type]: Key extends string
        ? Type[Key] extends object
          ? `${Key}` | `${Key}.${Path<Type[Key]>}`
          : `${Key}`
        : never;
    }[keyof Type]
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
};

export const i18n = (lang: Lang, path: Path<Locale>): string => {
  return traverse(locales[lang], path);
};

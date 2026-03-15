import type { Lang } from "./i18n";

export const readLanguageFromNavigator = (): Lang | null => {
  const navigatorLang = navigator.language.toLowerCase();

  if (navigatorLang.startsWith("en")) {
    return "en";
  }

  if (navigatorLang.startsWith("sv")) {
    return "sv";
  }

  if (navigatorLang === "ru") {
    return "ru";
  }

  return null;
};

const supportedLocales: Set<Lang> = new Set(["ru", "en", "sv"]);

export const isSupported = (lang: unknown): lang is Lang =>
  supportedLocales.has(lang as Lang);

export const langToLocaleString: { [key in Lang]: Intl.LocalesArgument } = {
  en: "en-US",
  sv: "sv-SE",
  ru: "ru-RU",
};

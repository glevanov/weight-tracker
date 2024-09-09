import type { Lang } from "./i18n";

export const readLanguageFromNavigator = (): Lang | null => {
  const navigatorLang = navigator.language.toLowerCase();

  if (navigatorLang.startsWith("en")) {
    return "en";
  }

  if (navigatorLang.startsWith("sv")) {
    return "sv";
  }

  if (navigatorLang === "zh-tw") {
    return "zh-tw";
  }

  if (navigatorLang === "ru") {
    return "ru";
  }

  return null;
};
const supportedLocales: Set<Lang> = new Set(["ru", "en", "sv", "zh-tw"]);

export const isSupported = (lang: unknown): lang is Lang =>
  supportedLocales.has(lang as Lang);

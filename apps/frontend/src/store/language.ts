import { derived, writable } from "svelte/store";
import type { Lang } from "../i18n/i18n";
import { i18n as i18nUnwrapped } from "../i18n/i18n";
import { isSupported, readLanguageFromNavigator } from "../i18n/util";

const LANGUAGE_KEY = "weight-tracker:language";

const loadLanguage = (): Lang => {
  const storedLanguage = localStorage.getItem(LANGUAGE_KEY);
  if (isSupported(storedLanguage)) {
    return storedLanguage;
  }

  const browserLanguage = readLanguageFromNavigator();
  if (browserLanguage !== null) {
    persistLanguage(browserLanguage);
    return browserLanguage;
  }

  persistLanguage("en");
  return "en";
};

const persistLanguage = (lang: Lang) =>
  localStorage.setItem(LANGUAGE_KEY, lang);

export const language = writable<Lang>(loadLanguage());

export const switchLanguage = (lang: Lang) => {
  language.set(lang);
  persistLanguage(lang);
};

export const i18n = derived(
  language,
  ($language) => (path: Parameters<typeof i18nUnwrapped>[1]) =>
    i18nUnwrapped($language, path),
);

type Lang = "ru";

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

const ru = {
  login: {
    login: "Логин",
    password: "Пароль",
    submit: "Войти",
    emptyCredentials: "Имя пользователя и пароль не могут быть пустыми",
  },
  addWeight: {
    header: "Введите ваш вес",
    success: "Вес успешно добавлен",
    submit: "Отправить",
    showGraph: "Показать график",
  },
  chart: {
    twoWeeks: "Две недели",
    month: "Месяц",
    quarter: "Квартал",
    year: "Год",
    allData: "Все данные",
    addWeight: "Добавить вес",
    loading: "Загружаем!",
    errorOccurred: "Произошла ошибка",
  },
  initialLoading: {
    loading: "Прогреваем сервер! Это может занять несколько минут.",
    failed:
      "Не удалось прогреть сервер. Попробуй обновить страницу или зайти позже.",
  },
};

type Locale = typeof ru;

const locales: Record<Lang, Locale> = {
  ru,
};

export const i18n = (path: Path<Locale>): string => {
  const selectedLocale: Lang = "ru";

  return traverse(locales[selectedLocale], path);
};

import type { Locale } from "../i18n";

export const sv: Locale = {
  name: "Svenska",
  login: {
    login: "Logga in",
    password: "Lösenord",
    submit: "Skicka",
    emptyCredentials: "Användarnamn och lösenord får inte vara tomma",
  },
  addWeight: {
    header: "Ange din vikt",
    success: "Vikt tillagd framgångsrikt",
    submit: "Skicka",
    showGraph: "Visa graf",
    empty: "Vikt kan inte vara tom",
  },
  chart: {
    twoWeeks: "Två veckor",
    month: "Månad",
    quarter: "Kvartal",
    year: "År",
    allData: "All data",
    addWeight: "Lägg till vikt",
    loading: "Laddar!",
    errorOccurred: "Ett fel uppstod",
  },
  initialLoading: {
    loading: "Startar servern! Detta kan ta några minuter.",
    failed:
      "Misslyckades med att starta servern. Försök att uppdatera sidan eller kom tillbaka senare.",
  },
};

import type { Locale } from "../i18n";

export const en: Locale = {
  name: "English",
  login: {
    login: "Login",
    password: "Password",
    submit: "Submit",
    emptyCredentials: "Username and password cannot be empty",
    logout: "Logout",
  },
  addWeight: {
    header: "Enter your weight",
    success: "Weight added successfully",
    submit: "Submit",
    showGraph: "Show graph",
    empty: "Weight cannot be empty",
  },
  chart: {
    twoWeeks: "Two weeks",
    month: "Month",
    quarter: "Quarter",
    year: "Year",
    allData: "All data",
    addWeight: "Add weight",
    loading: "Loading!",
    errorOccurred: "An error occurred",
  },
  initialLoading: {
    loading: "Warming up the server! This may take a few minutes.",
    wakeLock: "Acquire wake lock",
    failed:
      "Failed to warm up the server. Try refreshing the page or come back later.",
  },
};

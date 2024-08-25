import { writable } from "svelte/store";

export type Screen = "addWeight" | "chart" | "login" | "initialLoad";

export const currentScreen = writable<Screen>("initialLoad");

export const switchScreen = (screen: Screen) => currentScreen.set(screen);

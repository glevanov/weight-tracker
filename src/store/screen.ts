import { writable } from "svelte/store";

export type Screen = "addWeight" | "chart" | "login";

export const currentScreen = writable<Screen>("login");

export const switchScreen = (screen: Screen) => currentScreen.set(screen);

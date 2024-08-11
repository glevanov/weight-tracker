import { writable } from "svelte/store";
import { nanoid } from "nanoid";

const TIMEOUT = 3000;

export type Toast = {
  message: string;
  id: string;
};

export const toasts = writable<Toast[]>([]);

export const addToast = (message: string) => {
  const id = nanoid();
  toasts.update((all) => [{ message, id }, ...all]);

  setTimeout(() => dismissToast(id), TIMEOUT);
};

export const dismissToast = (id: string) => {
  toasts.update((all) => all.filter((toast) => toast.id !== id));
};

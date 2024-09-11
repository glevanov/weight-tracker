import { writable } from "svelte/store";
import { nanoid } from "nanoid";

const TIMEOUT = 9000;

export type ToastType = "success" | "error";

export type Toast = {
  message: string;
  id: string;
  type: ToastType;
};

export const toasts = writable<Toast[]>([]);

export const addToast = (message: string, type: ToastType) => {
  const id = nanoid();
  toasts.update((all) => [{ message, id, type }, ...all]);

  setTimeout(() => dismissToast(id), TIMEOUT);
};

export const dismissToast = (id: string) => {
  toasts.update((all) => all.filter((toast) => toast.id !== id));
};

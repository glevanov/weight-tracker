import { jwtDecode } from "jwt-decode";
import { writable } from "svelte/store";
import { readToken } from "../api/token";
import { isToken } from "../api/util";

const readUserFromToken = (token: string) => {
  try {
    const decoded = jwtDecode(token);

    if (!isToken(decoded)) {
      return null;
    }

    return decoded.username;
  } catch {
    return null;
  }
};

const getTokenOnLoad = () => {
  const data = readToken();
  if (data === null) {
    return null;
  }
  return readUserFromToken(data);
};

export const user = writable<string | null>(getTokenOnLoad());

export const updateUserFromToken = (token: string) => {
  user.set(readUserFromToken(token));
};

export const clearUser = () => {
  user.set(null);
};

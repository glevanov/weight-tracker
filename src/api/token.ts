const TOKEN_KEY = "weight-tracker:token";

export const saveToken = (token: string) => {
  localStorage.setItem(TOKEN_KEY, token);
};

export const readToken = () => {
  const token = localStorage.getItem(TOKEN_KEY);
  return typeof token === "string" ? token : null;
};

export const getAuthHeader = () => {
  const token = readToken();
  return `Bearer ${token}`;
};

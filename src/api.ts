const apiUrl = "http://localhost:3000";

export const getWeights = async () => {
  const result = await fetch(`${apiUrl}/weights`);
  return result.json();
};

export const addWeight = async (weight: number) => {
  const result = await fetch(`${apiUrl}/weights`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ weight }),
  });
  return result.json();
};

/// <reference types="svelte" />
import App from "./app.svelte";

const app = new App({
  target: document.getElementById("app") as Element,
});

export default app;

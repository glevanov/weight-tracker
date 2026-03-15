import App from "./app.svelte";
import { mount } from "svelte";

const app = mount(App, {
  target: document.getElementById("app") as Element,
});

export default app;

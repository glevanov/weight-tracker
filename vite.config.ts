import { defineConfig } from "vite";
import { svelte, vitePreprocess } from "@sveltejs/vite-plugin-svelte";

/** @type {import('vite').UserConfig} */
export default defineConfig({
  plugins: [
    svelte({
      preprocess: [vitePreprocess()],
    }),
  ],
  test: {
    include: ["**/*.test.ts"],
    environment: "happy-dom",
  },
});

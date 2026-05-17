import { defineConfig } from "vitest/config";
import { svelte, vitePreprocess } from "@sveltejs/vite-plugin-svelte";

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

<script lang="ts">
  import type { Lang } from "../i18n/i18n";
  import type { ChangeEventHandler } from "svelte/elements";
  import { language, switchLanguage } from "../store/language";
  import { i18n } from "../i18n/i18n";

  export let column = false;

  const selectOptions: { value: Lang; name: string }[] = [
    { value: "en", name: i18n("en", "name") },
    { value: "ru", name: i18n("ru", "name") },
    { value: "sv", name: i18n("sv", "name") },
    { value: "zh-tw", name: i18n("zh-tw", "name") },
  ];

  const handleSelect: ChangeEventHandler<HTMLSelectElement> = (event) => {
    const value = (event.target as HTMLSelectElement).value as Lang;
    switchLanguage(value);
  };

  let lang: Lang;
  language.subscribe((value) => (lang = value));
</script>

<div class={"page"}>
  <div class="header">
    <select bind:value={lang} on:change={handleSelect}>
      {#each selectOptions as option}
        <option value={option.value}>{option.name}</option>
      {/each}
    </select>
  </div>

  <div class={`content ${column ? "column" : ""}`}>
    <slot />
  </div>
</div>

<style>
  .page {
    height: 100%;

    box-sizing: border-box;
  }

  .header {
    display: flex;
    justify-content: flex-end;
    padding: 0.4em 1em;
  }

  .content {
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 1em;

    box-sizing: border-box;
  }

  .column {
    flex-direction: column;
  }
</style>

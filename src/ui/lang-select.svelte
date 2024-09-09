<script lang="ts">
  import { i18n, type Lang } from "../i18n/i18n";
  import type { ChangeEventHandler } from "svelte/elements";
  import { language, switchLanguage } from "../store/language";

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

<select bind:value={lang} on:change={handleSelect}>
  {#each selectOptions as option}
    <option value={option.value}>{option.name}</option>
  {/each}
</select>

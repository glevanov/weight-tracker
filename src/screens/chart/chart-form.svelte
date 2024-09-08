<script lang="ts">
  import type { Range } from "./types";
  import type { ChangeEventHandler } from "svelte/elements";
  import { i18n } from "../../i18n";

  export let handleShowAddWeight: () => void;
  export let selectedRange: Range;
  export let onSelectRange: (value: Range) => void;

  const selectOptions: { value: Range; name: string }[] = [
    { value: "14-days", name: i18n("chart.twoWeeks") },
    { value: "30-days", name: i18n("chart.month") },
    { value: "90-days", name: i18n("chart.quarter") },
    { value: "365-days", name: i18n("chart.year") },
    { value: "all-time", name: i18n("chart.allData") },
  ];

  const handleSelect: ChangeEventHandler<HTMLSelectElement> = (event) => {
    const value = (event.target as HTMLSelectElement).value as Range;
    onSelectRange(value);
  };
</script>

<form class="form">
  <select bind:value={selectedRange} on:change={handleSelect}>
    {#each selectOptions as option}
      <option value={option.value}>{option.name}</option>
    {/each}
  </select>

  <button type="button" on:click={handleShowAddWeight}
    >{i18n("chart.addWeight")}</button
  >
</form>

<style>
  .form {
    display: flex;
    justify-content: space-between;
    gap: 0.4em;
    width: 100%;
  }
</style>

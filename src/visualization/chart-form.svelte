<script lang="ts">
  import type { Range } from "./types";
  import type { ChangeEventHandler } from "svelte/elements";

  export let handleShowAddWeight: () => void;
  export let selectedRange: Range;
  export let onSelectRange: (value: Range) => void;

  const selectOptions: { value: Range; name: string }[] = [
    { value: "14-days", name: "Две недели" },
    { value: "30-days", name: "Месяц" },
    { value: "90-days", name: "Квартал" },
    { value: "365-days", name: "Год" },
    { value: "all-time", name: "Все данные" },
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

  <button type="button" on:click={handleShowAddWeight}>Ввести вес</button>
</form>

<style>
  .form {
    display: flex;
    justify-content: space-between;
    gap: 0.4em;
    width: 100%;
  }
</style>

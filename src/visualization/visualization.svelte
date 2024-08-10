<script lang="ts">
  import { onMount } from "svelte";
  import { getWeights } from "../api";
  import Chart from "./chart.svelte";
  import type { Range } from "./types";
  import { readRange, storeRange } from "./persist-range";

  export let handleShowAddWeight: () => void;

  let promise: Promise<{ weight: number; timestamp: number }[]>;

  let selectedRange: Range = readRange() ?? "14-days";

  const onSelectRange = (value: Range) => {
    selectedRange = value;
    storeRange(value);
    promise = getWeights(selectedRange);
  };

  onMount(() => {
    promise = getWeights(selectedRange);
  });
</script>

{#await promise}
  Загружаем!
{:then weights}
  <Chart {handleShowAddWeight} {weights} {selectedRange} {onSelectRange} />
{:catch error}
  Произошла ошибка: {error.message}
{/await}

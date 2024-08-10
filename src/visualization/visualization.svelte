<script lang="ts">
  import { onMount } from "svelte";
  import { getWeights } from "../api/api";
  import Chart from "./chart.svelte";
  import type { Range } from "./types";
  import { readRange, storeRange } from "./persist-range";
  import type { Weight } from "../api/types";

  export let handleShowAddWeight: () => void;

  let promise: Promise<Weight[]>;

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

<div class="page">
  {#await promise}
    Загружаем!
  {:then weights}
    <Chart {handleShowAddWeight} {weights} {selectedRange} {onSelectRange} />
  {:catch error}
    Произошла ошибка: {error.message}
  {/await}
</div>

<style>
  .page {
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    padding: 1em;
    box-sizing: border-box;
  }
</style>

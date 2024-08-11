<script lang="ts">
  import { getWeights } from "../api/api";
  import Chart from "./chart.svelte";
  import ChartForm from "./chart-form.svelte";
  import type { Range } from "./types";
  import { readRange, storeRange } from "./persist-range";

  export let handleShowAddWeight: () => void;

  let selectedRange: Range = readRange() ?? "14-days";

  const onSelectRange = (value: Range) => {
    selectedRange = value;
    storeRange(value);
    weightsRequest = getWeights(selectedRange);
  };

  let weightsRequest = getWeights(selectedRange);
</script>

<div class="page">
  {#await weightsRequest}
    Загружаем!
  {:then result}
    <Chart weights={result.isSuccess ? result.data : []} />
    <ChartForm {handleShowAddWeight} {selectedRange} {onSelectRange} />
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

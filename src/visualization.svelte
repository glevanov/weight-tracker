<script lang="ts">
  import { getWeights } from "./api";
  import Chart from "./chart.svelte";

  export let handleShowAddWeight: () => void;

  let promise = getWeights();
</script>

{#await promise}
  Загружаем!
{:then weights}
  <Chart {handleShowAddWeight} {weights} />
{:catch error}
  Произошла ошибка: {error.message}
{/await}

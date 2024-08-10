<script lang="ts">
  import { Line } from "svelte-chartjs";
  import type { Range } from "./types";

  import {
    Chart as ChartJS,
    LineElement,
    LinearScale,
    PointElement,
    CategoryScale,
    Tooltip,
  } from "chart.js";
  import type { ChangeEventHandler } from "svelte/elements";

  export let handleShowAddWeight: () => void;
  export let weights: { weight: number; timestamp: number }[] = [];
  export let selectedRange: Range;
  export let onSelectRange: (value: Range) => void;

  ChartJS.register(
    LineElement,
    LinearScale,
    PointElement,
    CategoryScale,
    Tooltip,
  );

  const weightData: number[] = [];
  const labels: string[] = [];

  for (const entry of weights) {
    weightData.push(entry.weight);
    labels.push(
      new Date(entry.timestamp).toLocaleDateString("ru-RU", {
        month: "short",
        day: "numeric",
      }),
    );
  }

  const chartData = {
    labels: labels,
    datasets: [
      {
        label: "Вес",
        data: weightData,
        borderColor: "#303F9F",
        backgroundColor: "#303F9F",
      },
    ],
  };

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

<div class="chart">
  <Line data={chartData} />
</div>

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

  .chart {
    --button-height: 30px;

    width: 100%;
    max-height: calc(100% - var(--button-height));
    box-sizing: border-box;
    display: flex;
    justify-content: center;
    align-items: center;
  }
</style>

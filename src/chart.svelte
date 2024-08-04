<script lang="ts">
  export let handleShowAddWeight: () => void;

  import { Line } from "svelte-chartjs";
  import { weights } from "./static-data";

  import {
    Chart as ChartJS,
    LineElement,
    LinearScale,
    PointElement,
    CategoryScale,
  } from "chart.js";

  ChartJS.register(LineElement, LinearScale, PointElement, CategoryScale);

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
        label: "График веса",
        data: weightData,
        borderColor: "#303F9F",
        backgroundColor: "#303F9F",
      },
    ],
  };
</script>

<div class="page">
  <div class="chart">
    <Line data={chartData} />
  </div>

  <button on:click={handleShowAddWeight}>Ввести вес</button>
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

<script lang="ts">
  import { Line } from "svelte-chartjs";
  import {
    Chart as ChartJS,
    LineElement,
    LinearScale,
    PointElement,
    CategoryScale,
    Tooltip,
  } from "chart.js";

  import type { Weight } from "../../api/types";
  import type { Lang } from "../../i18n/i18n";
  import { language } from "../../store/language";
  import { langToLocaleString } from "../../i18n/util";

  export let weights: Weight[] = [];

  ChartJS.register(
    LineElement,
    LinearScale,
    PointElement,
    CategoryScale,
    Tooltip,
  );

  let lang: Lang;
  language.subscribe((value) => (lang = value));

  let data: number[];
  $: data = weights.map((entry) => entry.weight);

  let labels: string[];
  $: labels = weights.map((entry) =>
    new Date(entry.timestamp).toLocaleDateString(langToLocaleString[lang], {
      month: "short",
      day: "numeric",
    }),
  );

  $: chartData = {
    labels,
    datasets: [
      {
        label: "Вес",
        data,
        borderColor: "#303F9F",
        backgroundColor: "#303F9F",
      },
    ],
  };
</script>

<div class="chart">
  <Line data={chartData} />
</div>

<style>
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

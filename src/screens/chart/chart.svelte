<script lang="ts">
  import { Line } from "svelte-chartjs";
  import {
    Chart as ChartJS,
    LineElement,
    LinearScale,
    PointElement,
    CategoryScale,
    Tooltip,
    type ChartOptions,
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
        borderColor: "#2898BD",
        backgroundColor: "#2898BD",
      },
    ],
  };

  const options: ChartOptions<"line"> = {
    scales: {
      x: {
        ticks: {
          font: {
            size: 16,
          },
        },
      },
      y: {
        ticks: {
          font: {
            size: 16,
          },
        },
      },
    },
    plugins: {
      tooltip: {
        bodyFont: {
          size: 16,
        },
        titleFont: {
          size: 16,
        },
      },
    },
  };
</script>

<div class="chart">
  <Line data={chartData} {options} />
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

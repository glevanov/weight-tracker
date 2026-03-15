<script lang="ts">
  import {
    Chart as ChartJS,
    LineController,
    LineElement,
    LinearScale,
    PointElement,
    CategoryScale,
    Tooltip,
    type ChartOptions,
  } from "chart.js";
  import { get } from "svelte/store";

  import type { Weight } from "../../api/types";
  import type { Lang } from "../../i18n/i18n";
  import { language } from "../../store/language";
  import { langToLocaleString } from "../../i18n/util";

  interface Props {
    weights?: Weight[];
  }

  let { weights = [] }: Props = $props();

  ChartJS.register(
    LineElement,
    LineController,
    LinearScale,
    PointElement,
    CategoryScale,
    Tooltip,
  );

  let lang: Lang = $state(get(language));
  language.subscribe((value) => (lang = value));

  let data: number[] = $derived(weights.map((entry) => entry.weight));

  let labels: string[] = $derived(
    weights.map((entry) =>
      new Date(entry.timestamp).toLocaleDateString(langToLocaleString[lang], {
        month: "short",
        day: "numeric",
      }),
    ),
  );

  let chartData = $derived({
    labels,
    datasets: [
      {
        label: "Вес",
        data,
        borderColor: "#2898BD",
        backgroundColor: "#2898BD",
      },
    ],
  });

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

  let ref: HTMLCanvasElement | undefined = $state();

  $effect(() => {
    if (ref) {
      new ChartJS(ref, {
        type: "line",
        data: chartData,
        options,
      });
    }
  });
</script>

<div class="chart">
  <canvas id="weights" bind:this={ref}></canvas>
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

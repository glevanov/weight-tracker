<script lang="ts">
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

<Line data={chartData} />

<script lang="ts">
  import Chart from "./chart.svelte";
  import AddWeight from "./add-weight.svelte";

  type Screen = "addWeight" | "chart";
  let currentScreen: Screen = "addWeight";
  const switchScreen = (screen: Screen) => {
    currentScreen = screen;
  };
</script>

<div class="app">
  {#if currentScreen === "addWeight"}
    <AddWeight handleShowGraph={() => switchScreen("chart")} />
  {:else if currentScreen === "chart"}
    <Chart handleShowAddWeight={() => switchScreen("addWeight")} />
  {/if}
</div>

<style>
  .app {
    height: 100svh;
  }
</style>

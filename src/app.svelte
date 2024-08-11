<script lang="ts">
  import Visualization from "./visualization/visualization.svelte";
  import AddWeight from "./add-weight/add-weight.svelte";
  import ToastContainer from "./ui/toast-container.svelte";

  type Screen = "addWeight" | "chart";
  let currentScreen: Screen = "addWeight";
  const showChart = () => (currentScreen = "chart");
  const showAddWeight = () => (currentScreen = "addWeight");
</script>

<div class="app">
  {#if currentScreen === "addWeight"}
    <AddWeight handleShowGraph={showChart} />
  {:else if currentScreen === "chart"}
    <Visualization handleShowAddWeight={showAddWeight} />
  {/if}

  <ToastContainer />
</div>

<style>
  .app {
    height: 100svh;
  }
</style>

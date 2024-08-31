<script lang="ts">
  import InitialLoad from "./screens/initial-load/initial-load.svelte";
  import Login from "./screens/login/login.svelte";
  import Visualization from "./screens/chart/visualization.svelte";
  import AddWeight from "./screens/add-weight/add-weight.svelte";
  import ToastContainer from "./ui/toast-container.svelte";
  import Error from "./screens/error/error.svelte";
  import { currentScreen } from "./store/screen";
  import type { Screen } from "./store/screen";

  let screen: Screen;
  currentScreen.subscribe((value) => (screen = value));
</script>

<div class="app">
  {#if screen === "initialLoad"}
    <InitialLoad />
  {:else if screen === "login"}
    <Login />
  {:else if screen === "addWeight"}
    <AddWeight />
  {:else if screen === "chart"}
    <Visualization />
  {:else if screen === "error"}
    <Error />
  {/if}

  <ToastContainer />
</div>

<style>
  .app {
    height: 100svh;
  }
</style>

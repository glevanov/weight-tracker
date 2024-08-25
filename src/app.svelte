<script lang="ts">
  import Login from "./screens/login/login.svelte";
  import Visualization from "./screens/chart/visualization.svelte";
  import AddWeight from "./screens/add-weight/add-weight.svelte";
  import ToastContainer from "./ui/toast-container.svelte";
  import { currentScreen } from "./store/screen";
  import type { Screen } from "./store/screen";

  let screen: Screen;
  currentScreen.subscribe((value) => (screen = value));
</script>

<div class="app">
  {#if screen === "login"}
    <Login />
  {:else if screen === "addWeight"}
    <AddWeight />
  {:else if screen === "chart"}
    <Visualization />
  {/if}

  <ToastContainer />
</div>

<style>
  .app {
    height: 100svh;
  }
</style>

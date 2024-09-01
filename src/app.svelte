<script lang="ts">
  import Fade from "./ui/fade.svelte";
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
    <Fade>
      <InitialLoad />
    </Fade>
  {:else if screen === "login"}
    <Fade>
      <Login />
    </Fade>
  {:else if screen === "addWeight"}
    <Fade>
      <AddWeight />
    </Fade>
  {:else if screen === "chart"}
    <Fade>
      <Visualization />
    </Fade>
  {:else if screen === "error"}
    <Fade>
      <Error />
    </Fade>
  {/if}

  <ToastContainer />
</div>

<style>
  .app {
    height: 100svh;
  }
</style>

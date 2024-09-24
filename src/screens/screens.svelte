<script lang="ts">
  import { currentScreen, type Screen } from "../store/screen";
  import InitialLoad from "./initial-load/initial-load.svelte";
  import AddWeight from "./add-weight/add-weight.svelte";
  import Login from "./login/login.svelte";
  import Visualization from "./chart/visualization.svelte";
  import Error from "./error/error.svelte";

  let screen: Screen;
  currentScreen.subscribe((value) => (screen = value));
</script>

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

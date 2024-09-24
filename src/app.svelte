<script lang="ts">
  import Header from "./ui/header.svelte";
  import Screens from "./screens/screens.svelte";
  import ToastContainer from "./ui/toast-container.svelte";
</script>

<div class="app">
  <div class="page">
    <Header />
    <Screens />
  </div>

  <ToastContainer />
</div>

<style>
  .app {
    height: 100svh;
  }

  .page {
    height: 100%;
    display: grid;
    grid-template-rows: auto 1fr;
  }
</style>

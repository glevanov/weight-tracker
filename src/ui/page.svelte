<script lang="ts">
  import Header from "./header.svelte";

  export let column = false;
</script>

<div class={"page"}>
  <Header />

  <div class={`content ${column ? "column" : ""}`}>
    <slot />
  </div>
</div>

<style>
  .page {
    height: 100%;
    display: grid;
    grid-template-rows: auto 1fr;
  }

  .content {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 1em;

    box-sizing: border-box;
  }

  .column {
    flex-direction: column;
  }
</style>

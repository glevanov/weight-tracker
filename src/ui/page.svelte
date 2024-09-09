<script lang="ts">
  import LangSelect from "./lang-select.svelte";

  export let column = false;
</script>

<div class={"page"}>
  <div class="header">
    <LangSelect />
  </div>

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

  .header {
    display: flex;
    justify-content: flex-end;
    padding: 0.4em 1em;
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

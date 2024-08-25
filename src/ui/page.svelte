<script lang="ts">
  export let column = false;
</script>

<div class={`page ${column ? "column" : ""}`}>
  <slot />
</div>

<style>
  .page {
    height: 100%;
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

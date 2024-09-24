<script lang="ts">
  export let column = false;
</script>

<div class={`content ${column ? "column" : ""}`}>
  <slot />
</div>

<style>
  .content {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 1em;
    height: 100%;

    box-sizing: border-box;
  }

  .column {
    flex-direction: column;
  }
</style>

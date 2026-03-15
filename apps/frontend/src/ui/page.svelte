<script lang="ts">
  interface Props {
    column?: boolean;
    children?: import("svelte").Snippet;
  }

  let { column = false, children }: Props = $props();
</script>

<div class={`content ${column ? "column" : ""}`}>
  {@render children?.()}
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

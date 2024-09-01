<script lang="ts">
  import { fade } from "svelte/transition";
</script>

<div class="wrapper" transition:fade|global>
  <slot />
</div>

<style>
  .wrapper {
    width: 100%;
    height: 100%;
  }
</style>

<script lang="ts">
  import { fade } from "svelte/transition";
</script>

<div class="toast" transition:fade>
  <slot></slot>
</div>

<style>
  .toast {
    display: flex;
    align-items: center;
    width: 100%;
    max-width: 20rem;
    margin: 0 auto 0.5rem auto;
    padding: 0.75rem 1.5rem;

    background-color: IndianRed;
    border-radius: 0.2rem;
  }
</style>

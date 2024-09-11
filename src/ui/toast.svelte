<script lang="ts">
  import { fade } from "svelte/transition";

  export let type: "error" | "success" = "error";
</script>

<div class="toast {type}" transition:fade>
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

    border-radius: 0.2rem;

    color: var(--inverse-text-color);
  }

  .error {
    background-color: #c9372c;

    @media (prefers-color-scheme: dark) {
      background-color: #f87168;
    }
  }

  .success {
    background-color: #1f845a;

    @media (prefers-color-scheme: dark) {
      background-color: #4bce97;
    }
  }
</style>

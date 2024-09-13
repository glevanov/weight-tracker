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
    width: calc(100vw - 40px);
    max-width: 400px;
    margin: 0 auto 16px auto;
    padding: 16px;
    box-sizing: border-box;

    border-radius: 4px;
    box-shadow:
      0 8px 12px #091e4226,
      0 0 1px #091e424f;

    color: var(--inverse-text-color);

    pointer-events: auto;

    @media (prefers-color-scheme: dark) {
      box-shadow:
        0 0 0 1px #39424a,
        0 8px 12px #0304045c,
        0 0 1px 1px #03040480;
    }
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

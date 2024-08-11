<script lang="ts">
  import Toast from "./toast.svelte";

  import { toasts } from "../store/toast";
</script>

{#if $toasts}
  <section>
    {#each $toasts as toast (toast.id)}
      <Toast>{toast.message}</Toast>
    {/each}
  </section>
{/if}

<style>
  section {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    width: 100%;
    display: flex;
    margin-top: 1rem;
    justify-content: center;
    flex-direction: column;
    z-index: 1;
  }
</style>

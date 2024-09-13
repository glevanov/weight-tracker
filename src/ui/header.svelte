<script lang="ts">
  import LangSelect from "./lang-select.svelte";
  import { user as userStore } from "../store/user";

  let user: string | null;
  userStore.subscribe((value) => (user = value));
</script>

<div class="header">
  <div>
    {#if user !== null}
      {user}
    {/if}
  </div>

  <LangSelect />
</div>

<style>
  .header {
    --border: #091e4224;

    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 8px 20px;
    border-block-end: 1px solid var(--border);

    @media (prefers-color-scheme: dark) {
      --border: #a6c5e229;
    }
  }
</style>

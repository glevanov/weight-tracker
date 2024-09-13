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
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0.4em 1em;
  }
</style>

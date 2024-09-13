<script lang="ts">
  import LangSelect from "./lang-select.svelte";
  import { clearUser, user as userStore } from "../store/user";
  import { i18n } from "../store/language";
  import { clearToken } from "../api/token";
  import { switchScreen } from "../store/screen";

  let user: string | null;
  userStore.subscribe((value) => (user = value));

  const logout = () => {
    clearUser();
    clearToken();
    switchScreen("login");
  };
</script>

<div class="header">
  <div class="user">
    {#if user !== null}
      {user}
      <button
        class="g-button g-button--default g-button--icon logout"
        type="button"
        aria-label={$i18n("login.logout")}
        on:click={logout}
      >
        <svg width="24" height="24" viewBox="0 0 24 24" role="presentation"
          ><g fill="currentcolor" fill-rule="evenodd"
            ><path
              d="M4.977 11A.99.99 0 0 0 4 12c0 .551.437 1 .977 1h11.046A.99.99 0 0 0 17 12a.99.99 0 0 0-.977-1z"
            ></path><path
              d="M6.23 8.31 3.3 11.27a1.05 1.05 0 0 0 0 1.48l2.93 2.96a1.03 1.03 0 0 0 1.47 0 1.05 1.05 0 0 0 0-1.48L5.5 12l2.2-2.22a1.05 1.05 0 0 0 0-1.48 1.03 1.03 0 0 0-1.47 0M15.5 3H12v2h7v14h-7v2h7.01c1.1 0 1.99-.89 1.99-1.99V5a1.99 1.99 0 0 0-1.99-2z"
            ></path></g
          ></svg
        >
      </button>
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

  .user {
    display: flex;
    align-items: center;
    gap: 8px;
  }
</style>

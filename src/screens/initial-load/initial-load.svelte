<script lang="ts">
  import { onMount } from "svelte";
  import Page from "../../ui/page.svelte";
  import { checkHealth, checkSession } from "../../api/api";
  import { switchScreen } from "../../store/screen";
  import cats from "./cats.webp";
  import { i18n } from "../../store/language";
  import { acquireWakeLock, releaseWakeLock } from "./wake-lock";

  onMount(async () => {
    const healthCheckResult = await checkHealth();
    await releaseWakeLock();
    if (!healthCheckResult.isSuccess) {
      switchScreen("error");
      return;
    }

    const hasSession = await checkSession();

    if (hasSession) {
      switchScreen("addWeight");
    } else {
      switchScreen("login");
    }
  });
</script>

<Page column={true}>
  <button
    class="invisible-button"
    onclick={acquireWakeLock}
    aria-label={$i18n("initialLoading.wakeLock")}
  >
    <img class="cats" src={cats} alt="" />
  </button>
  <span class="text">{$i18n("initialLoading.loading")}</span>
</Page>

<style>
  .cats {
    width: 100%;
    max-width: 500px;

    animation: spin 15s infinite ease-in-out;
  }

  .text {
    margin-top: 1em;

    text-align: center;
  }

  @keyframes spin {
    0% {
      transform: rotate(0deg);
    }
    100% {
      transform: rotate(360deg);
    }
  }

  .invisible-button {
    border: none;
    background: none;
    outline: none;
    padding: 0;
  }
</style>

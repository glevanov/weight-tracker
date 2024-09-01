<script lang="ts">
  import { onMount } from "svelte";
  import Page from "../../ui/page.svelte";
  import { checkHealth, checkSession } from "../../api/api";
  import { switchScreen } from "../../store/screen";
  import cats from "./cats.webp";

  onMount(async () => {
    const healthCheckResult = await checkHealth();
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
  <img class="cats" src={cats} alt="" />
  <span class="text">Прогреваем сервер! Это может занять несколько минут.</span>
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
</style>

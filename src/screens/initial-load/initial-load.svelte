<script lang="ts">
  import { onMount } from "svelte";
  import Page from "../../ui/page.svelte";
  import { checkHealth, checkSession } from "../../api/api";
  import { switchScreen } from "../../store/screen";

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

<Page
  >Прогреваем сервер! Так как он бесплатный, это может занять несколько минут.</Page
>

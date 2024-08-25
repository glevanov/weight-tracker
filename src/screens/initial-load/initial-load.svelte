<script lang="ts">
  import { onMount } from "svelte";
  import Page from "../../ui/page.svelte";
  import { checkHealth, checkSession } from "../../api/api";
  import { switchScreen } from "../../store/screen";

  onMount(async () => {
    await checkHealth();

    const hasSession = await checkSession();

    if (hasSession) {
      switchScreen("addWeight");
    } else {
      switchScreen("login");
    }
  });
</script>

<Page>Загружаемся! Это может занять около минуты.</Page>

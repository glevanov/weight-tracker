<script lang="ts">
  import Page from "../../ui/page.svelte";
  import { i18n } from "../../store/language";
</script>

<Page>
  <span class="text">
    {$i18n("initialLoading.failed")}
  </span>
</Page>

<style>
  .text {
    text-align: center;
  }
</style>

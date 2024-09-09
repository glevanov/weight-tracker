<script lang="ts">
  import { onMount } from "svelte";
  import { addToast } from "../../store/toast";
  import { addWeight } from "../../api/api";
  import type { FormEventHandler } from "svelte/elements";
  import { switchScreen } from "../../store/screen";
  import Page from "../../ui/page.svelte";
  import { i18n } from "../../store/language";

  let value = "";

  const onSubmit: FormEventHandler<HTMLFormElement> = async () => {
    const result = await addWeight(value);
    if (!result.isSuccess) {
      addToast(result.error);
    } else {
      addToast($i18n("addWeight.success"));
    }
  };

  const handleShowGraph = () => switchScreen("chart");

  let ref: HTMLInputElement;

  onMount(() => {
    ref?.focus();
  });
</script>

<Page>
  <form class="form" on:submit|preventDefault={onSubmit} autocomplete="off">
    <label class="label" for="weight-input">{$i18n("addWeight.header")}</label>

    <input
      class="input"
      id="weight-input"
      bind:value
      inputmode="decimal"
      bind:this={ref}
    />

    <button class="button submit" type="submit"
      >{$i18n("addWeight.submit")}</button
    >

    <button class="button showGraph" type="button" on:click={handleShowGraph}>
      {$i18n("addWeight.showGraph")}
    </button>
  </form>
</Page>

<style>
  .form {
    display: grid;
    gap: 0.4em;
    grid-template-areas:
      "label label"
      "input button"
      "show show";
    grid-template-columns: auto auto;
    width: 100%;
    max-width: 400px;
  }

  .label {
    grid-area: label;
  }

  .submit {
    grid-area: button;
  }

  .showGraph {
    grid-area: show;
  }

  .input {
    grid-area: input;
    width: 100%;
    box-sizing: border-box;
  }
</style>

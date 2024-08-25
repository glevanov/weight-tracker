<script lang="ts">
  import { onMount } from "svelte";
  import { addToast } from "../../store/toast";
  import { addWeight } from "../../api/api";
  import type { FormEventHandler } from "svelte/elements";
  import { switchScreen } from "../../store/screen";
  import Page from "../../ui/page.svelte";

  let value = "";

  const onSubmit: FormEventHandler<HTMLFormElement> = async () => {
    const result = await addWeight(value);
    if (!result.isSuccess) {
      addToast(result.error);
    } else {
      addToast("Вес успешно добавлен");
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
    <label class="label" for="weight-input"> Введите ваш вес: </label>

    <input
      class="input"
      id="weight-input"
      bind:value
      inputmode="decimal"
      bind:this={ref}
    />

    <button class="button submit" type="submit">Отправить</button>

    <button class="button showGraph" type="button" on:click={handleShowGraph}>
      Показать график
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

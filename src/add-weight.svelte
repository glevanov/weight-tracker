<script lang="ts">
  export let handleShowGraph: () => void;
  let value = "";

  const onSubmit = (evt: Event) => {
    evt.preventDefault();
    alert("Ваш вес " + value);
  };
</script>

<div class="page">
  <form class="form" on:submit={onSubmit} autocomplete="off">
    <label class="label" for="weight-input"> Введите ваш вес: </label>

    <input class="input" id="weight-input" bind:value inputmode="decimal" />

    <button class="button submit" type="submit">Отправить</button>

    <button class="button showGraph" type="button" on:click={handleShowGraph}>
      Показать график
    </button>
  </form>
</div>

<style>
  .page {
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 1em;
    box-sizing: border-box;
  }

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

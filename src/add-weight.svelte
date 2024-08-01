<script lang="ts">
  let value = "";
  let error = "";

  const onSubmit = (evt: Event) => {
    evt.preventDefault();
    error = "текст ошибки";
    alert("ошибюка");
  };
</script>

<div class="page">
  <form class="form" on:submit={onSubmit}>
    <label class="label" for="weight-input"> Введите ваш вес: </label>

    <input class="input" id="weight-input" bind:value inputmode="decimal" />

    <button class="submit" type="submit">Отправить</button>
    {#if error.length > 0}
      <div class="error">
        {error}
      </div>
    {/if}
  </form>
</div>

<style>
  .page {
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .form {
    max-width: 300px;
    display: grid;
    gap: 0.4em;
    grid-template-areas:
      "label label"
      "input button"
      "error error";
  }

  .label {
    grid-area: label;
  }

  .submit {
    grid-area: button;
  }

  .error {
    grid-area: error;

    color: red;
  }

  .input {
    grid-area: input;
  }
</style>

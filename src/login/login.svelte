<script lang="ts">
  import { login } from "../api/api";
  import { addToast } from "../store/toast";

  let username = "";
  let password = "";

  const onSubmit = async () => {
    const isEmpty = username.trim() === "" || password.trim() === "";
    if (isEmpty) {
      addToast("Имя пользователя и пароль не могут быть пустыми");
      return;
    }

    const result = await login(username, password);
    if (!result.isSuccess) {
      addToast(result.error);
      return;
    }
    console.log(result);
  };
</script>

<div class="page">
  <form class="form" on:submit|preventDefault={onSubmit}>
    <label for="login">Логин</label>
    <input type="text" name="login" id="login" bind:value={username} />

    <label class="withGap" for="password">Пароль</label>
    <input
      type="password"
      name="password"
      id="password"
      bind:value={password}
    />

    <button class="withGap" type="submit">Войти</button>
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
    width: 100%;
    max-width: 400px;
    box-sizing: border-box;
  }

  .form > * {
    display: block;
    width: 100%;
    box-sizing: border-box;
  }

  .withGap {
    margin-top: 0.4em;
  }
</style>

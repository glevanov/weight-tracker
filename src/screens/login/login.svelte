<script lang="ts">
  import Page from "../../ui/page.svelte";
  import { login } from "../../api/api";
  import { addToast } from "../../store/toast";
  import { switchScreen } from "../../store/screen";
  import { i18n } from "../../store/language";

  let username = "";
  let password = "";
  let isLoading = false;

  const onSubmit = async () => {
    if (isLoading) {
      return;
    }

    const isEmpty = username.trim() === "" || password.trim() === "";
    if (isEmpty) {
      addToast($i18n("login.emptyCredentials"));
      return;
    }

    isLoading = true;
    const result = await login(username, password);
    isLoading = false;

    if (!result.isSuccess) {
      addToast(result.error);
    } else {
      switchScreen("addWeight");
    }
  };
</script>

<Page>
  <form class="form" on:submit|preventDefault={onSubmit}>
    <label for="login">{$i18n("login.login")}</label>
    <input type="text" name="login" id="login" bind:value={username} />

    <label class="withGap" for="password">{$i18n("login.password")}</label>
    <input
      type="password"
      name="password"
      id="password"
      bind:value={password}
    />

    <button class="g-button withGap" type="submit"
      >{$i18n("login.submit")}</button
    >
  </form>
</Page>

<style>
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

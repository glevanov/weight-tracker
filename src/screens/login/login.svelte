<script lang="ts">
  import type { FormEventHandler } from "svelte/elements";
  import Page from "../../ui/page.svelte";
  import { login } from "../../api/api";
  import { addToast } from "../../store/toast";
  import { switchScreen } from "../../store/screen";
  import { i18n } from "../../store/language";
  import { saveToken } from "../../api/token";
  import { updateUserFromToken } from "../../store/user";

  let username = $state("");
  let password = $state("");
  let isLoading = false;

  const onSubmit: FormEventHandler<HTMLFormElement> = async (evt) => {
    evt.preventDefault();
    if (isLoading) {
      return;
    }

    const isEmpty = username.trim() === "" || password.trim() === "";
    if (isEmpty) {
      addToast($i18n("login.emptyCredentials"), "error");
      return;
    }

    isLoading = true;
    const result = await login(username, password);
    isLoading = false;

    if (!result.isSuccess) {
      addToast(result.error, "error");
    } else {
      saveToken(result.data);
      updateUserFromToken(result.data);
      switchScreen("addWeight");
    }
  };
</script>

<Page>
  <form class="form" onsubmit={onSubmit}>
    <label for="login">{$i18n("login.login")}</label>
    <input
      class="g-input"
      type="text"
      name="login"
      id="login"
      bind:value={username}
    />

    <label class="withGap" for="password">{$i18n("login.password")}</label>
    <input
      class="g-input"
      type="password"
      name="password"
      id="password"
      bind:value={password}
    />

    <button class="g-button g-button--primary withGap" type="submit"
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

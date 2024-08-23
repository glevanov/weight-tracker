<script lang="ts">
  const onSubmit = async () => {
    console.log("submit");
  };
</script>

<div class="page">
  <form class="form" on:submit|preventDefault={onSubmit}>
    <label for="login">Логин</label>
    <input type="text" name="login" id="login" />

    <label for="password">Пароль</label>
    <input type="password" name="password" id="password" />
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
</style>

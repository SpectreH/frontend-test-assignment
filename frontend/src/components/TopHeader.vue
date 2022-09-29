<template>
  <header>
    <div class="header-content">
      <img class="logo" alt="Vue logo" src="../assets/Logo.svg">
      <div class="header-buttons">
        <button>Users</button>
        <button>Sign up</button>        
      </div>
    </div>
  </header>
</template>

<script>
export default {

}
</script>

<style lang="scss" scoped>
  header {
    justify-content: center;
    width: 100%;
    height: 60px;
    display: flex;

    .logo {
      height: 26px;
      width: 106px;
    }

    .header-content {
      max-width: 1170px;
      justify-content: space-between;     
      align-items: center;
      display: flex;
      width: inherit;
    }
    
    .header-buttons {
      display: flex;
      gap: 10px;
    }
  }
</style>
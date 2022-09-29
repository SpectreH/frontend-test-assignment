<template>
  <div class="container">
    <nav>
      <div class="nav-content">
        <div class="nav-content-text">
          <h1>Test assignment for front-end developer</h1>
          <p>What defines a good front-end developer is one that has skilled knowledge of HTML, CSS, JS with a vast understanding of User design thinking as they'll be building web interfaces with accessibility in mind. They should also be excited to learn, as the world of Front-End Development keeps evolving.</p>
        </div>
        <button>Sign up</button>
      </div>
    </nav>
  </div>
</template>

<script>

export default {
  name: 'HomeView',
  components: {
  }
}
</script>

<style lang="scss">
  nav {
    display: grid;
    background: linear-gradient(0deg, rgba(0, 0, 0, 0.5), rgba(0, 0, 0, 0.5)), url("../assets/pexels-alexandr-podvalny-1227513.jpeg");
    background-position: 60% 81%;
    background-size: 2500px;

    .nav-content {
      margin: auto;
      justify-items: center;
      text-align: center;
      max-width: 380px;
      display: grid;
      gap: 32px;

      .nav-content-text {
        color: #fff;
        display: grid;
        gap: 21px;
      }

      h1, p {
        margin: 0;
      }
    }
  }
</style>

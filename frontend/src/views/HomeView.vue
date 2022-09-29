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
    <main>
      <section>
        <h1 class="main-header">Working with GET request</h1>

        <div class="main-people-list">
          <PeopleCard/>
        </div>

        <div class="main-people-more">
          <button>Show More</button>
        </div>
      </section>

      <section style="margin-bottom: 100px">
        <h1 class="form-header">Working with POST request</h1>
        <form action="POST" style="width: 380px; display: grid; margin: auto;">
          <div style="display: grid; gap: 50px">
            <NormalInput label="Test" />
            <NormalInput label="Test" />
            <NormalInput label="Test" helperMessage="Help" />
          </div>

          <div class="radio-input-group" style="display: grid; gap: 7px; margin-top: 25px; margin-bottom: 47px">
            <label>Select your position</label>
            <RadioButton :value="1" label="Test" selected/> 
            <RadioButton :value="1" label="Test" />                       
            <RadioButton :value="1" label="Test" />                       
            <RadioButton :value="1" label="Test" />                                                                     
          </div>

          <div class="input-group upload" error="false">
            <button type="button" upload>Upload</button>
            <input placeholder="Upload your photo"  value="Item" disabled>
          </div>

          <div style="display: flex; justify-content: center; margin-top: 50px">
            <button @click.prevent>Sign up</button>
          </div>
        </form>
      </section>
    </main>
  </div>
</template>

<script>
import PeopleCard from "@/components/PeopleCard.vue"
import RadioButton from "@/components/RadioButton.vue"
import NormalInput from "@/components/NormalInput.vue"
export default {
  name: 'HomeView',
  components: {
    PeopleCard,
    RadioButton,
    NormalInput
  }
}
</script>

<style lang="scss">
  nav {
    display: grid;
    background: linear-gradient(0deg, rgba(0, 0, 0, 0.5), rgba(0, 0, 0, 0.5)), url("../assets/pexels-alexandr-podvalny-1227513.jpeg");
    background-position: 60% 81%;
    background-size: 2500px;
    margin-bottom: 140px;

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

  main {
    .main-header, .form-header {
      text-align: center;
      margin: 0;
      margin-bottom: 50px;
      margin-left: 16px;
      margin-right: 16px;
    }

    .main-people-list {
      justify-content: center;
      display: flex;
      flex-wrap: wrap;
      gap: 29px;
    }

    .main-people-more {
      margin-top: 50px;
      display: flex;
      justify-content: center;
      margin-bottom: 140px;      
    }
  }

  .card {
    width: 370px;
    height: 254px;
    background-color: #fff;
    border-radius: 10px;

    .card-content {
      padding: 20px;
      display: grid;
      gap: 20px;
      justify-items: center;
    }
  }
</style>

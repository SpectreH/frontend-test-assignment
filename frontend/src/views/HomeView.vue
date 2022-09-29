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
      <section ref="section1">
        <h1 class="main-header">Working with GET request</h1>

        <div class="main-people-list">
          <PeopleCard/>
        </div>

        <div class="main-people-more">
          <button>Show More</button>
        </div>
      </section>

      <section style="margin-bottom: 100px" ref="section2">
        <h1 class="form-header">Working with POST request</h1>
        <form action="POST" style="width: 380px; display: grid; margin: auto;">
          <div style="display: grid; gap: 50px">
            <NormalInput label="Your name" v-model="name" />
            <NormalInput label="Email" v-model="email" />
            <NormalInput label="Phone" v-model="phone" helperMessage="+38 (XXX) XXX - XX - XX" />
          </div>

          <div class="radio-input-group" style="display: grid; gap: 7px; margin-top: 25px; margin-bottom: 47px">
            <label>Select your position</label>
            <RadioButton v-model="position" value="Frontend developer" label="Frontend developer" selected/> 
            <RadioButton v-model="position" value="Backend developer" label="Backend developer" />                       
            <RadioButton v-model="position" value="Designer" label="Designer" />                       
            <RadioButton v-model="position" value="QA" label="QA" />                                                                     
          </div>

          <UploadInput label="Upload you foto"/>

          <div style="display: flex; justify-content: center; margin-top: 50px">
            <button @click.prevent="submit">Sign up</button>
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
import UploadInput from "@/components/UploadInput.vue"
export default {
  name: 'HomeView',
  components: {
    PeopleCard,
    RadioButton,
    NormalInput,
    UploadInput
  },
  data() {
    return {
      name: "",
      email: "",
      phone: "",
      position: "Frontend developer",
    }
  },
  methods: {
    submit() {
      console.log(this.name, this.email, this.phone, this.position)
    },
  }
}
</script>
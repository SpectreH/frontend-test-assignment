<template>
  <div>
    <label class="radio-container">
      <span>{{ label }}</span>
      <input name="radio" type="radio" :value="value" :checked="selected"/>
      <div class="b-input"></div>
    </label>
  </div> 
</template>

<script>
export default {
  props: {
    label: String,
    value: String,
    selected: Boolean
  }
}
</script>

<style>

</style>
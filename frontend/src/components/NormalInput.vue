<template>
  <div>
    <div class="input-group" error="false">
      <label class="floating-label">{{ label }}</label>         
      <input placeholder="Label">
    </div>
    <label v-if="helperMessage && !errorMessage" class="helper-message">{{ helperMessage }}</label>    
    <label v-if="errorMessage" class="helper-message error">{{ errorMessage }}</label>        
  </div>
</template>

<script>
export default {
  props: {
    label: String,
    helperMessage: String,
    errorMessage: String,
  }
}
</script>

<style>

</style>
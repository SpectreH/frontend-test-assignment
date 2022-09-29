<template>
  <div>
    <div class="input-group" :error="errorMessage">
      <label v-if="modelValue" class="floating-label">{{ label }}</label>         
      <input :placeholder="label" :value="modelValue" @input="$emit('update:modelValue', $event.target.value)">
    </div>
    <label v-if="helperMessage && !errorMessage" class="helper-message">{{ helperMessage }}</label>    
    <label v-if="errorMessage" class="helper-message error">{{ errorMessage }}</label>        
  </div>
</template>

<script>
export default {
  props: {
    modelValue: String,
    label: String,
    helperMessage: String,
    errorMessage: String,
  } 
}
</script>

<style>

</style>
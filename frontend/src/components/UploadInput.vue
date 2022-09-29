<template>
  <div>
    <div class="input-group upload" error="false">
      <button type="button" upload>Upload</button>
      <input :placeholder="label" disabled>
    </div>
    <label v-if="helperMessage && !errorMessage" class="helper-message">{{ helperMessage }}</label>    
    <label v-if="errorMessage" class="helper-message error">{{ errorMessage }}</label>   
  </div>   
</template>

<script>
export default {
  props: {
    value: File,
    label: String,
    helperMessage: String,
    errorMessage: String,    
  }
}
</script>

<style>

</style>
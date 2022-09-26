<template>
  <div class="hello" style="display: grid; gap: 20px">
    <h1>The five boxing wizards jump quickly.</h1>
    <p>Lorem ipsum Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis a lectus vitae urna dignissim laoreet a eget arcu. Sed eget gravida nisl. Integer molestie purus velit, et efficitur quam porttitor ut. Vestibulum tincidunt odio at lacus ornare luctus.</p>
    <button>Normal</button>
    <button disabled>Disabled</button>
    
    <div style="width: 380px; display: grid; gap: 20px">
      <div>
        <div class="input-group" error="true">
          <button type="button" upload>Upload</button>
          <input placeholder="Upload your photo" disabled>
        </div>
        <label class="error-message">Error text</label>
      </div>   

      <div class="input-group" error="false">
        <button type="button" upload>Upload</button>
        <input placeholder="Upload your photo"  value="Item" disabled>
      </div>   

      <input placeholder="Upload your photo">      
    </div>
        
    <p>Normal</p>
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  props: {
    msg: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
</style>

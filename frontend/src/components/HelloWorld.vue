<template>
  <div class="hello" style="display: grid; gap: 20px">
    <h1>The five boxing wizards jump quickly.</h1>
    <p>Lorem ipsum Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis a lectus vitae urna dignissim laoreet a eget arcu. Sed eget gravida nisl. Integer molestie purus velit, et efficitur quam porttitor ut. Vestibulum tincidunt odio at lacus ornare luctus.</p>
    <button>Normal</button>
    <button disabled>Disabled</button>
    
    <div style="width: 380px; display: grid; gap: 20px">
      <div>
        <div class="input-group upload" error="true">
          <button type="button" upload>Upload</button>
          <input placeholder="Upload your photo" disabled>
        </div>
        <label class="helper-message error">Error text</label>
      </div>   

      <div class="input-group upload" error="false">
        <button type="button" upload>Upload</button>
        <input placeholder="Upload your photo"  value="Item" disabled>
      </div> 

      <div>
        <div class="input-group" error="true">
          <label class="floating-label">Label</label>         
          <input placeholder="Label">
        </div>
        <label class="helper-message error">Error text</label>              
      </div>
        

      <div>
        <input placeholder="Upload your photo">     
        <label class="helper-message">Helper text</label>
      </div>   
    </div>

    <div class="profile-image">
      <img src="../assets/photo-cover.svg" height="70" width="70">
    </div>

    <div class="tooltip">MaximilianMaximilian@gmail.com
      <span class="tooltiptext">MaximilianMaximilian@gmail.com</span>
    </div>

    <div class="spinner"></div>


    <div class="radio-input-group" style="display: grid; gap: 7px">
      <div>
        <label class="radio-container">
          <span>Checked</span>
          <input name="radio" type="radio" />
          <div class="b-input"></div>
        </label>
      </div>
      <div>
        <label class="radio-container">
          <span>Unchecked</span>
          <input name="radio" type="radio" />
          <div class="b-input"></div>
        </label>
      </div>
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

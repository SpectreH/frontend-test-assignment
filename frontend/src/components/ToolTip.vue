<template>
  <div class="tooltip">
    <div class="tooltip-main-text">
      {{ text }}
    </div>
    <span class="tooltiptext">{{ text }}</span>
  </div> 
</template>

<script>
export default {
  props: {
    text: String
  }
}
</script>

<style>

</style>
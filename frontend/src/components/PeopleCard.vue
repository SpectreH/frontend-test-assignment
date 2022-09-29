<template>
  <div class="card">
    <div class="card-content">
      <div class="profile-image">
        <img src="../assets/photo-cover.svg" height="70" width="70">
      </div>

      <ToolTip text="0"/>

      <div style="display: grid; justify-items: center">
        <ToolTip text="1"/>
        <ToolTip text="2"/>
        <ToolTip text="3"/>                    
      </div>
    </div>
  </div>
</template>

<script>
import ToolTip from "@/components/ToolTip.vue"
export default {
  components: {
    ToolTip
  }
}
</script>

<style>

</style>
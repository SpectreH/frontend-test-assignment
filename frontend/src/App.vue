<template>
  <TopHeader/>
  <router-view/>
</template>

<script>
import TopHeader from "@/components/TopHeader.vue"
export default {
  components: {
    TopHeader
  }
}
</script>

<style lang="scss">
</style>

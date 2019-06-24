<template lang="html">
  <search></search>
  {{headLabel}}
  <menu></menu>
</template>

<script>
import search from '@component/search.vue'
import menu from '@component/menu.vue'
export default {
  props:{
    headLabel:'hoge'
  }
}
</script>

<style lang="css" scoped>
</style>

<template lang="html">
<table>
  <monoItem v-for:"item in list" v-bind:mono="item" v-bind:key="item.id"></monoItem>
</table>
</template>

<script>
import monoItem from './mono-item.vue'

export default {
  props:{
    list:null
  }
  components:{
    monoItem
  }
}
</script>

<style lang="css" scoped>
</style>

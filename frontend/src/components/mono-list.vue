<template lang="html">
  <table>
    <monoItem v-for="item in list" mono="item" :key="item.id"></monoItem>
  </table>
</template>

<script>
import monoItem from './mono-item.vue'

export default {
  computed:{
    list(){
      return this.$store.state.monoList
    }
  },
  mounted:function(){
    this.$store.dispatch('getMonoList', null, null)
  },
  components:{
    monoItem
  }
}
</script>

<style lang="css" scoped>
</style>

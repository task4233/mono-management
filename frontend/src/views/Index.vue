<template lang="html">
  <div class="index">
    <search></search>
    <tagList></tagList>
    <menuBar></menuBar>
    <monoList></monoList>
  </div>
</template>

<script>
import monoList from '../components/mono-list.vue'
import tagList from '../components/tag-list.vue'
import search from '../components/search.vue'
import menuBar from '../components/menu.vue'

export default {
  name:'index',
  components:{
    monoList,
    tagList,
    search,
    menuBar
  }
}
</script>

<style lang="css" scoped>
</style>

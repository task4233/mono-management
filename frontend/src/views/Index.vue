<template lang="html">
  <div class="index">
    <search></search>
    <tagList></tagList>
    <menu></menu>
    <monoList></monoList>
  </div>
</template>

<script>
import monoList from '@components/mono-list.vue'
import tagList from '@components/tag-list.vue'
import search from '@component/search.vue'
import menu from '@component/menu.vue'

export default {
  name:'index',
  components:{
    monoList,
    tagList,
    search,
    menu
  }
}
</script>

<style lang="css" scoped>
</style>

<template lang="html">
  <div class="index">
    <tagList></tagList>
    <monoList></monoList>
  </div>
</template>

<script>
import monoList from '../components/mono-list.vue'
import tagList from '../components/tag-list.vue'

export default {
  name:'index',
  components:{
    monoList,
    tagList,
  }
}
</script>

<style lang="css" scoped>
</style>

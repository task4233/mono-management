<template lang="html">
  <b-container class="index">
    <b-form-group>
      <tagList></tagList>
    </b-form-group>
    <b-form-group>
      <monoList></monoList>
    </b-form-group>
  </b-container>
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

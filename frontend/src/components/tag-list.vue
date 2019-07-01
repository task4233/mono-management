<template lang="html">
  <select v-model="selected">
    <option v-for "option in headLabels" v-bind:value="option.id">
      {{option.name}}
    </option>
  </select>
</template>

<script>
export default {
  data:{
    selected:null
  },
  computed:{
    headLabel(){
      this.$store.tag_list
    }
  },
  mounted:function(){
    this.$store.dispatch('getTagList')
  },
  watch:{
    selected(){
      this.$store.dispach('getMonoList', selected, null)
    }
  }

}
</script>

<style lang="css" scoped>
</style>

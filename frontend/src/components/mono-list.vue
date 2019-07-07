<template lang="html">
  <div>
    <b-table :items="list" :fields="fields" striped>
    </b-table>
  </div>
</template>

<script>
export default {
  data (){
    return {
      fields: ['name']
    }
  },
  computed: {
    list() {
      console.log(this.$store.state.mono_list)
      return this.$store.state.mono_list
    }
  },
  mounted:function(){	
    this.$store.dispatch("getMonoList", {name:null, tagId:0})
  }
}
</script>

<style lang="css" scoped>
</style>

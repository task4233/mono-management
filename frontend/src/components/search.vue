<template lang="html">
  <div v-on:click="search">search</div>
  <div v-if:"searchMode">
    <form>
      <input type="text" name="" value="">
    </form>
    <div v-on:click="exitSearch">x</div>
  </div>
</template>

<script>
export default {
  data:{
    searchMode:false
    words:''
  },
  methods:{
    search:function(){
      if(this.searchMode){
        this.$store.dispatch('getMonoList', null, words)
      }else{
        this.searchMode = true
      }
    },
    exitSearch:function(){
      this.searchMode = false
      this.words=''
      
    }

  }
}
</script>

<style lang="css" scoped>
</style>

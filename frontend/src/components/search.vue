<template lang="html">
  <div>
    <div v-on:click="search">search</div>
    <div v-if = "searchMode">
      <form>
        <input type="text" maxlength="64">
      </form>
      <div v-on:click="exitSearch">x</div>
    </div>
  </div>
</template>

<script>
export default {
  data:function(){
    return{
      searchMode:false,
      words:''
    }
  },
  methods:{
    search:function(){
      if(this.searchMode){
        this.$store.dispatch('getMonoList', null, this.words)
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

<template lang="html">
  <div>
    <div v-on:click="search">search</div>
    <div v-if = "searchMode">
      <form v-on:keyup.enter="search" v-on:keyup.esc="exitSearch">
        <input type="text" v-model="keyword" maxlength="64">
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
      keyword:''
    }
  },
  methods:{
    search:function(){
      if(this.searchMode){
        const searchdata = {
          name:this.keyword,
          tagId:this.$store.state.select_tag
        }
        this.$store.dispatch('getMonoList', searchdata)
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

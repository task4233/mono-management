<template lang="html">
  <div>
    <div v-on:click="search">search</div>
    <div v-if = "searchMode">
      <form>
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
        console.log(this.keyword)
        const searchdata = {
          name:this.keyword,
          tagId:1
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

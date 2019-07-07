<template lang="html">
  <div>
    <div @click="search">search</div>
    <div v-if = "searchMode">
      <form @submit.prevent="dummy">
        <input
        type="text"
        v-model="keyword"
        maxlength="64"
        @keyup.enter="search"
        @keyup.esc="exitSearch"
        >
      </form>
      <div @click="exitSearch">x</div>
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
    dummy() {
      // dummy
      // inputでenterを押した時のリロードを抑止する役割
    },
    search(){
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
    exitSearch(){
      this.searchMode = false
      this.words=''

    }

  }
}
</script>

<style lang="css" scoped>
</style>

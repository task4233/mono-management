<template lang="html">
  <search v-on:searchWord="searchMonoList"></search>
  <select v-model="selected">
    <option v-for "option in headLabels" v-bind:value="option.id">
      {{option.name}}
    </option>
  </select>
  <menu></menu>
</template>

<script>
import search from '@component/search.vue'
import menu from '@component/menu.vue'
export default {
  props:{
    headLabel:[
      {id:0, name:'hoge', parentid:1}
    ]
  },
  data:{
    selected:0
  },
  methods:{
    searchMonoList:funciton(words){
      
    }
  },
  watch:{
    selected:function(){
      this.$emit('changeTag', this.selected)
    }
  }
}
</script>

<style lang="css" scoped>
</style>

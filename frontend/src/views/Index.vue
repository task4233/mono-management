<template lang="html">
  <div class="index">
    <header></header>
    <monoList></monoList>
  </div>
</template>

<script>
import monoList from '@components/mono-list.vue'
import header from '@components/header.vue'
export default {
  name:'index',
  components:{
    monoList,
    header
  },
  data:{
    isModal:false
  }
  mounted:funciton(){
    api_base = '/api/v1/'
    axios.get(api_base + 'tag/')
      .then(funtion(response){
        response
        axios.get(api_base + 'mono/')
          .then(response){

          }
      })
      .catch(function(error){
        if(error.response.status == 401){
          //loginページへジャンプ
        }
      })
  }
}
</script>

<style lang="css" scoped>
</style>

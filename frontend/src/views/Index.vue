<template lang="html">
  <div class="index">
    <monoList></monoList>
  </div>
</template>

<script>
import monoList from '@components/mono-list.vue'
import
export default {
  name:'index',
  components:{
    monoList
  },
  data:{
    isModal:false
  }
  mounted:funciton(){
    api_base = '/api/v1/'
    axios.get(api_base + 'tag/')
      .then(funtion(response){
        axios.get(api_base + 'mono/')
      })
      .catch(function(error){
        if()
      })
  }
}
</script>

<style lang="css" scoped>
</style>

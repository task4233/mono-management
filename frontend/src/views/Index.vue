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
    BASE_URL = '/api/v1/'
    axios.get(BASE_URL + 'tag/')
      .then(funtion(response){
        axios.get(BASE_URL + 'mono/')
      })
      .catch(function(error){
        if()
      })
  }
}
</script>

<style lang="css" scoped>
</style>

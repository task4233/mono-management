<template lang="html">
  <tr v-on:click=""><td>{{mono.name}}</td><td>{{mono.info[{{mono.primary}}]}}</td></tr>
</template>

<script>
export default {
  name:'monoItem',
  data:function(){
    mono = null
  }
}
</script>

<style lang="css" scoped>
</style>

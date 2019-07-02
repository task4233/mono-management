<template lang="html">
  <table>
    <!-- <monoItem v-for="item in list" item="item" :key="item.Id"></monoItem> -->
    <tr v-for="item in list" :key="item.Id">
      <td>{{ item.name }}</td>
      </tr>
  </table>
</template>

<script>
import monoItem from './mono-item.vue'

export default {
  computed:{
    list(){
      console.log(this.$store.state.mono_list)
      return this.$store.state.mono_list
    }
  },
  mounted:function(){
    this.$store.dispatch('getMonoList', null, null)
  },
  components:{
    monoItem
  }
}
</script>

<style lang="css" scoped>
</style>

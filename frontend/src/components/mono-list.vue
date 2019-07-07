<template lang="html">
  <div>
    <b-table :items="list" :fields="fields" striped>
    </b-table>
  </div>
  <!--
  <table>
    <monoItem v-for="item in list" item="item" :key="item.Id"></monoItem>
    <tr v-for="item in list" :key="item.Id">
      <td>{{ item.name }}</td>
      </tr>
  </table>
  -->
</template>

<script>
export default {
  data (){
    return {
      fields: ['name']
    }
  },
  computed: {
    list() {
      console.log(this.$store.state.mono_list)
      return this.$store.state.mono_list
    }
  }
}
</script>

<style lang="css" scoped>
</style>

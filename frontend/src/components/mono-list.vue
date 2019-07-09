<template lang="html">
  <div>
    <table striped>
      <tr v-for="item in list" :key="item.Id" @click="editMono(item.Id)">
        <td>
          {{item.name}}
        </td>
      </tr>
    </table>
    <b-button @click="createMono">+</b-button>
  </div>
</template>

<script>
export default {
  data (){
    return {
      fields: ['name'],
      monoId:0
    }
  },
  computed: {
    list() {
      console.log(this.$store.state.mono_list)
      return this.$store.state.mono_list
    }
  },
  mounted:function(){
    this.$store.dispatch("getMonoList", {name:null, tagId:0})
  },
  methods:{
    createMono(){
      this.$router.push({path:'/mono/new'})
    },
    editMono(Id){
      const path = '/mono/' + Id
      this.$router.push({path:path})
    }
  }
}
</script>

<style lang="css" scoped>
</style>

<template lang="html">
  <div>
    <b-table :items="list" :fields="fields" striped>
      <template slot="row-details" slot-scope="row">
        <span @click="editMono({{row.Id}})"></span>
      </template>
    </b-table>
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
      const path = '/mono/' + String(Id)
      this.$router.push({path:path})
    }
  }
}
</script>

<style lang="css" scoped>
</style>

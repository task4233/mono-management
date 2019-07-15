<template lang="html">
  <div>
    <table striped>
      <tr v-for="li in list" :key="li.Id" v-b-modal.showMonoData @click="showModal(li)">
        <td>
          {{ li.name }}
        </td>
      </tr>
    </table>

    <b-modal
    id="showMonoData"
    scrollable
    v-bind:title="this.modalName"
    ref="modal"
    @hidden="resetModal"
    @ok="editMono(this.monoId)">
      <table striped>
        <tr v-for="data in datas" :key="data.id">
          <td>{{ data.name }}</td>
          <td>{{ data.type == "num" ? data.num : (data.type == "str" ? data.str : data.timestamp) }}</td>
        </tr>
      </table>
    </b-modalid="showMonoData">
  </div>


</template>

<script>
export default {
  data (){
    return {
      monoName: '',
      monoId:0
    }
  },
  computed: {
    list() {
      console.log(this.$store.state.mono_list)
      return this.$store.state.mono_list
    },
    items() {
      return this.$store.state.item_list
    },
    datas() {
      return this.$store.state.data_list
    }
  },
  mounted:function(){
    this.$store.dispatch("getMonoList", {name:null, tagId:0})
  },
  methods:{
    async showModal(item) {
      this.monoId = item.Id
      this.monoName = item.name + "編集"
      console.log(this.monoId)
      await this.$store.dispatch("getMonoData", this.monoId)
    },
    createMono(){
      this.$router.push({path:'/mono/new'})
    },
    editMono(Id){
      const path = '/mono/' + Id
      this.$router.push({path:path})
    },
    resetModal(){
      this.monoId = 0,
      this.monoName = ''
    },
    handleOK(bvModalEvt){
      bvModalEvt.preventDefault()
      editMono(this.monoId)
    }
  }
}
</script>

<style lang="scss" scoped>
table {
  margin: auto;
}
</style>
<template lang="html">
<<<<<<< HEAD
  <b-form-select v-model="selected">
    <option value="0">タグを選択</option>
    <option v-for="option in headLabels" v-bind:value="option.id" :key="option.id">
      {{ option.name}}
=======
  <select v-model="selected">
    <option v-bind:value="0">全てのタグ</option>
    <option v-for="option in headLabels"
      v-bind:value="option.Id"
      v-bind:key="option.Id">
      {{ option.name }}
>>>>>>> 048d4850fdd2e58a6ac82aad31d4abaa66f2e994
    </option>
  </b-form-select>
</template>

<script>
export default {
  data() {
    return {
      selected: 0
    }
  },
  computed: {
    headLabels(){
      return this.$store.state.tag_list
    }
  },
  mounted: function() {
    this.$store.dispatch("getTagList")
    this.$store.dispatch("tagSelect", this.selected)
    this.selected = this.$store.state.select_tag
  },
  watch: {
    selected() {
      this.$store.dispatch("tagSelect", this.selected)
      this.$store.dispatch("getMonoList", {name:null, tagId:this.selected})
    }
  }
}

</script>

<style lang="css" scoped>
</style>

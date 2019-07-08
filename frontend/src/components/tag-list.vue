<template lang="html">
  <select v-model="selected">
    <option v-bind:value="0">全てのタグ</option>
    <option v-for="option in headLabels"
      v-bind:value="option.Id"
      v-bind:key="option.Id">
      {{ option.name }}
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

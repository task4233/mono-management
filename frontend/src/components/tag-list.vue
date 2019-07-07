<template lang="html">
  <select v-model="selected">
    <option value="0">全てのタグ</option>
    <option v-for="option in headLabels"
      v-bind:value="option.Id"
      v-bind:key="option.Id">
      {{ option.name }}
    </option>
  </select>
</template>

<script>
export default {
  data() {
    return {
      selected: ""
    }
  },
  computed: {
    headLabels(){
      return this.$store.state.tag_list;
    }
  },
  mounted: function() {
    this.$store.dispatch("getTagList")
    this.$store.dispatch("tagSelect", this.selected)
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

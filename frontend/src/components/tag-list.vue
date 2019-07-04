<template lang="html">
  <select v-model="selected">
    <option value="0">全てのタグ</option>
    <option v-for="option in headLabels" v-bind:value="option.id" :key="option.id">
      {{ option.name}}
    </option>
  </select>
</template>

<script>
export default {
  data: function() {
    return {
      selected: null
    }
  },
  computed: {
    headLabels() {
      return this.$store.tag_list;
    }
  },
  mounted: function() {
    this.$store.dispatch("getTagList");
  },
  watch: {
    selected() {
      console.log(this.selected)
      this.$store.dispatch("getMonoList", {name:null, tagId:this.selected})
    }
  }
}

</script>

<style lang="css" scoped>
</style>

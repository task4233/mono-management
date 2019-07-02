<template lang="html">
  <select v-model="selected">
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
    };
  },
  computed: {
    headLabel() {
      return this.$store.tag_list;
    }
  },
  mounted: function() {
    this.$store.dispatch("getTagList");
  },
  watch: {
    selected() {
      this.$store.dispach("getMonoList", this.selected, null);
    }
  }
};
</script>

<style lang="css" scoped>
</style>

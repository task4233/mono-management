<template lang="html">
  <b-form-select v-model="selected">
    <option v-bind:value="0">タグを選択</option>
    <option v-for="option in tags"
      v-bind:value="option.Id"
      v-bind:key="option.Id">
      {{ option.name }}
    </option>
  </b-form-select>
</template>

<script>
export default {
  data: function() {
    return {
      selected: 0
    }
  },
  computed: {
    tags(){
      return this.$store.state.tag_list
    }
  },
  mounted: function() {
    this.$store.dispatch('getTagList')
    this.selected = this.$store.state.select_tag
    if(this.selected == null){
      this.selected = 0
      this.$store.dispatch('tagSelect', 0)
    }
  },
  watch: {
    selected() {
      this.$store.dispatch('tagSelect', this.selected)
      this.$store.dispatch('getMonoList', {name:null, tagId:this.selected})
    }
  }
}

</script>

<style lang="css" scoped>
</style>

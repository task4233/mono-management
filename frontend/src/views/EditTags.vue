<template lang="html">
  <div>
    <ol>
      <li v-for="tag in tags" :key="tag.id" v-on:click=""></li>
    </ol>
  </div>
</template>

<script>
export default {
  computed:{
    tags(){
      this.$store.state.tag_list
    }
  },
  method:{

  }
}
</script>

<style lang="css" scoped>
</style>

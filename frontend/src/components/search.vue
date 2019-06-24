<template lang="html">
  <div v-if:"searchMode">
    <form>検索ワードを記入</form>
  </div>
  <div v-else>

  </div>

</template>

<script>
export default {
  data:{
    searchMode:false
  }
}
</script>

<style lang="css" scoped>
</style>

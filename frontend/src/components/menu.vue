<template lang="html">
  <div v-if:"menuMode">
    ×<br>
    <span v-on:click="logout">ログアウト</span><br>
    <span>ヘルプ</span><br>
    <span>このソフトについて</span><br>
  </div>
  <div v-else>
    ≡
  </div>
</template>

<script>
export default {
  data:{
    menuMode:false
  }
  methods:{
    logout:functon(event){
      axios.delete('api/v1/user/logoiut')
      this.$router.push({path:'/login'})
    }
  }
}
</script>

<style lang="css" scoped>
</style>

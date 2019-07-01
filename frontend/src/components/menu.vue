<template lang="html">
  <div v-if:"menuMode">
    ×<br>
    <span v-on:click="logout">ログアウト</span><br>
    <a href="https://hackmd.io/giR3aLTXSH6VUQigiqEK0A"><span>このソフトについて</span><br></a>
  </div>
  <div v-on:click="changeView" v-else>
    ≡
  </div>
</template>

<script>
export default {
  data:{
    menuMode:false
  }
  methods:{
    logout:function(){
      axios.delete('api/v1/user/logout')
      this.$store.commit('resetUserData')
      this.$router.push({path:'/login'})
    },
    changeView:function(){
      this.menuMode = !this.menuMode
    }
  }
}
</script>

<style lang="css" scoped>
</style>

<template lang="html">
  <div>
    <div v-if="menuMode">
      <div v-on:click="changeView">×</div><br>
      <span v-on:click="jumpEditTags">タグ編集</span><br>
      <span v-on:click="logout">ログアウト</span><br>
      <a href="https://hackmd.io/giR3aLTXSH6VUQigiqEK0A"><span>このソフトについて</span><br></a>
    </div>
    <div v-else>
      <div v-on:click="changeView">≡</div>
    </div>
  </div>
</template>

<script>
export default {
  data:function(){
    return{
      menuMode:false
    }
  },
  methods:{
    logout:function(){
      this.$axios.delete('api/v1/user/logout')
      this.$store.commit('resetUserData')
      this.$router.push({path:'/login'})
    },
    changeView:function(){
      this.menuMode = !this.menuMode
    },
    jumpEditTags(){
      this.$router.push({path:'/EditTags'})
    }
  }
}
</script>

<style lang="css" scoped>
</style>

<template>
  <div id="app">
    <navbar v-if="isLogin"></navbar>
    <router-view/>
  </div>
</template>

<script>
import navbar from './components/navbar.vue'
export default{
  components: {
    navbar
  },
  computed:{
    isLogin(){
      const name = this.$route.name
      if(name !== 'login' && name !== 'signup' && name !== 'NotFound') return true
      else return false
    }
  },
  created(){
    this.$axios.dispatch('getUserName')
  }
}
</script>

<style>
#app {
  font-family: 'Makinas-4-Square';
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>

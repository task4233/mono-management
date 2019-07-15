<template>
  <div id="app">
    <navbar v-if="isLogin"></navbar>
    <router-view/>
    <errorModal></errorModal>
    <!--<button @click="showError">showErrortest</button>-->
  </div>
</template>

<script>
import errorModal from './components/errorModal.vue'
import navbar from './components/navbar.vue'
export default{
  components: {
    navbar,
    errorModal
  },
  computed:{
    isLogin(){
      const name = this.$route.name
      if(name !== 'login' && name !== 'signup' && name !== 'NotFound') return true
      else return false
    }
  },
  methods:{
    /*
    showError(){
      this.$store.commit('showError', 'showErrorTest')
    }
    */
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

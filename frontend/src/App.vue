<template>
  <div id="app">
    <navbar></navbar>
    <router-view/>
  </div>
</template>

<script>
import navbar from './components/navbar.vue'
export default{
  components: {
    navbar
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

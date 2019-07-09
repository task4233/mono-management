<template>
  <div id="app">
    <div id="nav">
      <router-link to="/">Home</router-link> |
      <router-link to="/about">About</router-link> |
      <router-link to="/mono/new">Create</router-link> |
      <router-link to="/EditTags">EditTags</router-link>
    </div>
    <router-view/>
  </div>
</template>

<script>
export default{

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

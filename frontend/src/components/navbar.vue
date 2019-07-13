<template lang="html">
  <div>
    <b-navbar toggleable="lg" type="dark" variant="success" class="mb-3">
      <b-navbar-brand to="/about">mono-management</b-navbar-brand>
      <b-navbar-toggle target="nav-collapse" class="ml-auto"></b-navbar-toggle>

      <b-collapse id="nav-collapse" is-nav>
        <b-navbar-nav>
          <b-nav-item to="/">Home</b-nav-item>
          <b-nav-item to="/about">About</b-nav-item>
          <b-nav-item to="/mono/new">Create</b-nav-item>
          <b-nav-item to="/EditTags">Tags</b-nav-item>
        </b-navbar-nav>
        <b-navbar-nav class="ml-auto">
          <b-nav-form @submit.prevent="dummy">
            <b-form-input size="sm" class="mr-sm-2" type="text"
              placeholder="Search" maxlength="64" v-model="keyword" 
              @keyup="search" @keyup.esc="exitSearch">
            </b-form-input>
            <b-button size="sm" class="my-2 my-sm-0" type="submit"
              @click="search">
              Search
            </b-button>
          </b-nav-form>
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>
  </div>
</template>
<script>
import router from '../router'

export default {
  data:function(){
    return{
      searchMode:false,
      keyword:''
    }
  },
  methods:{
    dummy() {
      // dummy
      // inputでenterを押した時のリロードを抑止する役割
    },
    search(){
      router.push("/")
      const searchdata = {
        name:this.keyword,
        tagId:0
      }
      this.$store.dispatch('getMonoList', searchdata)
    },
    exitSearch(){
      this.keyword=''

    }

  }
}
</script>

<template lang="html">
  <div class="index">
    <header v-on:changeTag="getMonoList"></header>
    <monoList list="mono_list"></monoList>
  </div>
</template>

<script>
import monoList from '@components/mono-list.vue'
import header from '@components/header.vue'
export default {
  name:'index',
  components:{
    monoList,
    header
  },
  data:{
    isModal:false,
    modal_message:'spam',
    mono_list:null,
    tag_list:null
  },
  methods:{
    getMonoList:function(tagId){
      axios.get('/api/v1/mono/'+tagId+'/')
        .then(funtion(response){
        })
    }
    getTagList:function(){
      axios.get('/api/v1/tag/')
        .then(function(response){
          if(response.status){
            this.tag_list = response.status.data.tags
          }
        })
    }
  },
  mounted:funciton(){
    api_base = '/api/v1/'
    axios.get(api_base + 'tag/')
      .then(funtion(response){
        tagList = response
        axios.get(api_base + 'mono/')
          .then(response){
            if(response.status == true){
              monoList = response.monoList
            }else{
              modalMessage = response.msg
            }
          }
      })
      .catch(function(error){
        if(error.response.status == 401){
          //loginページへジャンプ
          this.$router.push({path:'/login'})
        }
      })
  }
}
</script>

<style lang="css" scoped>
</style>

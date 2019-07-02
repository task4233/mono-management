import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const api_base = '/api/v1/'

export default new Vuex.Store({
  state: {
    tag_list:null,
    mono_list:null,
    mono_data:null,
    modal_message:''
  },
  mutations: {
    setTagList:function(state, tags){
      state.tag_list = tags
    },
    setMonoList:function(state, monos){
      state.mono_list = monos
    },
    setMonoData:function(state, data){
      state.mono_data = data
    },
    setModalMessage:function(state, msg){
      state.modal_message = msg
    },
    resetMonoData:function(state){
      state.mono_data = null
    },
    resetUserData:function(state){
      state.tag_list = null
      state.mono_list = null
    }
  },
  actions: {
    getMonoList({commit},tagId, name){
      this.$axios.post(api_base + 'search/',{
        name:name,
        tagId:tagId
      })
      .then(function(response){
        if(response.status){
          commit('setMonoList',response.monoList)
        }else{
          commit('setModalMessage',response.msg)
        }
      })
      .catch(function(error){
        if(error.response.status == 401){
          commit('resetUserData')
          //loginページへジャンプ
          this.$router.push({path:'/login'})
        }
      })
    },
    getTagList({commit}){
      this.$axios.get(api_base + 'tag/')
      .then(function(response){
        if(response.status){
          commit('setTagList',response.tags)
        }
      })
      .catch(function(error){
        if(error.response.status == 401){
          commit('resetUserData')
          //loginページへジャンプ
          this.$router.push({path:'/login'})
        }
      })
    }

  }
})

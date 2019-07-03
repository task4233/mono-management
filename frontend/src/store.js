import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

import Axios from 'axios'
import router from './router'

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
    getMonoList({commit},mtagId, mname){
      console.log(mname)
      const data = {
        name: mname,
        tagId: mtagId
      }
      console.log(data)
      Axios.post(api_base + 'search/', data)
      .then(function(response){
        if(response.data.Status){
          console.log(response.data.Data)
          commit('setMonoList',response.data.Data)
        }else{
          commit('setModalMessage',response.data.message)
        }
      })
      .catch(function(error){
        if(error.response.status == 401){
          commit('resetUserData')
          //loginページへジャンプ
          router.push({path:'/login'})
        }
      })
    },
    getTagList({commit}){
      Axios.get(api_base + 'tag/')
      .then(function(response){
        if(response.status){
          commit('setTagList',response.tags)
        }
      })
      .catch(function(error){
        if(error.response.status == 401){
          commit('resetUserData')
          //loginページへジャンプ
          router.push({path:'/login'})
        }
      })
    }

  }
})

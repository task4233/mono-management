import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

import Axios from 'axios'
import Router from './router'

const api_base = '/api/v1/'

export default new Vuex.Store({
  state: {
    tag_list:null,
    mono_list:null,
    mono_data:null,
    select_tag:null,
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
    setSelectTag:function(state, tagId){
      console.log("mutation:setSelectTag," + tagId)
      state.select_tag = Number(tagId)
    },
    resetMonoData:function(state){
      state.mono_data = null
    },
    resetUserData:function(state){
      state.tag_list = null
      state.mono_list = null
      state.select_tag = null
    }
  },
  actions: {
    getMonoList({commit},searchdata){
      console.log(searchdata)
      // searchdataがnullの時はデータを整形する
      if (!searchdata) {
        searchdata = {
          name: ''
        }
      }
      Axios.post(api_base + 'search/',searchdata)
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
          Router.push({path:'/login'})
        }
      })
    },
    getTagList({commit}){
      Axios.get(api_base + 'tag/')
      .then(function(response){
        if(response.status){
          commit('setTagList',response.data.tags)
        }
      })
      .catch(function(error){
        if(error.response.status == 401){
          commit('resetUserData')
          //loginページへジャンプ
          Router.push({path:'/login'})
        }
      })
    },
    tagSelect({commit}, tagId){
      console.log("tagSelect:" + tagId)
      commit("setSelectTag", tagId)
    },
    createTag({commit, dispatch}, tagData){
      Axios.post(api_base + 'tag/new', tagData)
      .then(function(response){
        if(response.data.Status){
          dispatch('getTagList')
        }else{
          //エラーメッセージを表示
        }
      })
      .catch(function(error){
        if(error.response.status == 401){
          commit('resetUserData')
          //loginページへジャンプ
          Router.push({path:'/login'})
        }
      })
    },
    changeTagData({commit, dispatch}, tagData){
      Axios.put(api_base + 'tag/:'+ tagData.Id , tagData)
      .then(function(response){
        if(response.data.Status){
          dispatch('getTagList')
        }else{
          //エラーメッセージを表示

        }
      })
      .catch(function(error){
        if(error.response.status == 401){
          commit('resetUserData')
          //loginページへジャンプ
          Router.push({path:'/login'})
        }
      })
    }
  }
})

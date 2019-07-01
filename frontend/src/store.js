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
    updateTagList:function(state, tags){
      state.tag_list = tags
    },
    updateMonoList:function(state, monos){
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
    getMonoList({commit, state},tagId, name){
      axios.post(api_base + 'search/',{
        name:name,
        tagId:tagId
      })
      .then(function(response){
        if(response.status == true){
          monoList = response.monoList
        }else{
          commit.setModalMessage(response.msg)
        }
      })
      .catch(function(error){
        if(error.response.status == 401){
          store.commit('resetUserData')
          //loginページへジャンプ
          this.$router.push({path:'/login'})
        }
      })
    },
    getTagList({commit, state}){
      axios.get(api_base + 'tag/')
      .then(function(response){
        if(response.status){
          this.tag_list = response.status.data.tags
        }
      })
      error(function(error){
        if(error.response.status == 401){
          store.commit('resetUserData')
          //loginページへジャンプ
          this.$router.push({path:'/login'})
        }
      })
    }

  }
})

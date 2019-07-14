<template lang="html">
  <div>
    <div class="header">
      <div class="sm-col-1" @click="back">
        ←
      </div>
      <div class ="sm-col-8">タグ編集</div>
      <div class = "sm-col-1"></div>
    </div>
    <ol>
      <li v-for="tag in tags" :key="tag.id" v-b-modal.editTag @click="editMode(tag)">
        {{tag.name}}
      </li>
    </ol>
    <b-button v-b-modal.editTag @click="modalName='タグを追加'">+</b-button>
    <!--以下，モーダルウィンドウ-->
    <b-modal
    id="editTag"
    v-bind:title="this.modalName"
    ref="modal"
    @hidden="resetModal"
    @ok="handleOK">
      <!--https://bootstrap-vue.js.org/docs/components/modal/#prevent-closing-->
      <b-form-group>
        タグ名<b-form-input type="text" v-model="tagName" :state="this.isMatchedName"/>

        親タグ<tagList></tagList>
        <span class="errmsg" v-if="isTagLoop">このタグを親タグにすることはできません！</span>
      </b-form-group>
    </b-modal>
  </div>
</template>

<script>
import tagList from '../components/tag-list.vue'
export default {
  data:function(){
    return{
      modalName:'hoge',
      tagId:0,
      tagName:'',

    }
  },
  components:{
    tagList
  },
  computed:{
    tags(){
      return this.$store.state.tag_list
    },
    isTagLoop(){
      // statusがtrueなら問題ないので, !演算子でひっくり返しています
      return !this.$store.state.modal_status
    },
    isMatchedName(){
      return (0 < this.tagName.length && this.tagName.length < 64) ? true : false
    }
  },
  methods:{
    updateTag: async function (){
      console.log("updateTag called!")
      // 一応ここでもcommitしとく
      this.$store.commit('setModalStatus', true)
      if(this.tagId == 0){
        await this.$store.dispatch('createTag', {name:this.tagName, Id:0, parentId:this.$store.state.select_tag})
      }
      else{
        await this.$store.dispatch('changeTagData', {name:this.tagName, Id:this.tagId, parentId:this.$store.state.select_tag})
      }
      
      // changeTagDataが実行されたのちにstateが更新されるので
      // ループ判定は後にするべき
      if(!this.isMatchedName || this.isTagLoop){
        console.log(this.tagName)
        console.log("tagLoop Occured!")
        return
      }

      this.$nextTick(()=>{
        this.$refs.modal.hide()
      })

      console.log(this.tags)
      
      console.log("updateTag end!")
    },
    editMode(tag){
      console.log("editMode called!")

      this.modalName='タグを編集',
      this.tagId = tag.Id
      this.tagName = tag.name
      this.$store.dispatch('tagSelect', tag.parentId)
      console.log("editMode end!")
    },
    resetModal(){
      // 処理的にここで呼ぶべきだったっぽい
      // ここでModalの状態を直接trueに初期化
      // actionsを飛ばしているので, ダメだったらactionsに何か挟んで
      this.$store.commit('setModalStatus', true)
      this.tagId = 0,
      this.tagName = ''
      this.$store.dispatch('tagSelect', 0)
    },
    handleOK(bvModalEvt){
      bvModalEvt.preventDefault()
      this.updateTag()
    },
    back(){
      this.$router.push({path:'/'})
    }
  },
  mounted:function(){    
    this.$store.dispatch("getTagList")
  }
}
</script>

<style lang="css" scoped>
</style>

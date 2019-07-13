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
      return (this.tagId != 0 && this.$store.state.tag_list == this.tagId) ? true : false
    },
    isMatchedName(){
      return (0 < this.tagName.length && this.tagName.length < 64) ? true : false
    }
  },
  methods:{
    updateTag(){
      if(!this.isMatchedName || this.isTagLoop){
        return
      }
      if(this.tagId == 0){
        this.$store.dispatch('createTag', {name:this.tagName, Id:0, parentId:this.$store.state.select_tag})
      }
      else{
        this.$store.dispatch('changeTagData', {name:this.tagName, Id:this.tagId, parentId:this.$store.state.select_tag})
      }
      this.$nextTick(()=>{
        this.$refs.modal.hide()
      })
    },
    editMode(tag){
      this.modalName='タグを編集',
      this.tagId = tag.Id
      this.tagName = tag.name
      this.$store.dispatch('tagSelect', tag.parentId)
    },
    resetModal(){
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

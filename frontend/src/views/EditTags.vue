<template lang="html">
  <div>
    <ol>
      <li v-for="tag in tags" :key="tag.id" v-b-modal.editTag @click="editMode(tag)">
        {{tag.name}}
      </li>
    </ol>
    <b-button v-b-modal.editTag @click="modalName='タグを追加'">+</b-button>
    <b-modal id="editTag" v-bind:title="this.modalName" ref="modal" @hidden="resetModal" @ok="updateTag">
      <!--https://bootstrap-vue.js.org/docs/components/modal/#prevent-closing-->
      <form class="">
        タグ名<b-form-input type="text" label="タグ名" v-model="tagName"/>
        親タグ<tagList></tagList>
      </form>

    </b-modal>
  </div>
</template>

<script>
import tagList from '../components/tag-list.vue'
export default {
  data:function(){
    return{
      modalName:'hoge',
      tagId:null,
      tagName:null,
      //parent:null
    }
  },
  components:{
    tagList
  },
  computed:{
    tags(){
      return this.$store.state.tag_list
    }
  },
  method:{
    updateTag(){
      if(this.tagId == null){
        this.$store.dispatch('createTag', {name:this.tagName, tagId:this.tagId, parent:this.$store.state.select_tag})
      }
      else{
        this.$store.dispatch('chengeTagData', {name:this.tagName, tagId:this.tagId, parent:this.$store.state.select_tag})
      }
    },
    editMode(tag){
      this.modalName='タグを編集',
      this.tagId = tag.tagId
      this.tagName = tag.tagName
    },
    resetModal(){
      this.tagId = null,
      this.tagName = null
    }
  },
  mounted:function(){
    this.$store.dispatch("getTagList")
  }
}
</script>

<style lang="css" scoped>
</style>

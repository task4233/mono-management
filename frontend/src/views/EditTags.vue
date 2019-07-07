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
        <b-form-input type="text" label="タグ名" v-model="tagName"/>
        <!--親タグ
        <b-form-select v-model="parent">
          <option value ="0"></option>
          <option v-for="tag in tags" :key="tag.id">{{tag.name}}</option>
        </b-form-select>
      -->
      </form>

    </b-modal>
  </div>
</template>

<script>
//import TagList from '../components/tag_list.vue'
export default {
  data:function(){
    return{
      modalName:'hoge',
      tagId:null,
      tagName:null,
      //parent:null
    }
  },
  computed:{
    tags(){
      return this.$store.state.tag_list
    }
  },
  method:{
    updateTag(){
      if(this.tagId == null){
        this.$store.dispatch('createTag', {name:this.tagName, tagId:this.tagId})
      }
      else{
        this.$store.dispatch('chengeTagData', {name:this.tagName, tagId:this.tagId})
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

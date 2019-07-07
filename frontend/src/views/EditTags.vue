<template lang="html">
  <div>
    <ol>
      <li v-for="tag in tags" :key="tag.id" v-b-modal.editTag v-on:click="modalName='タグを編集'">
        {{tag.name}}
      </li>
    </ol>
    <b-button v-b-modal.editTag v-on:click="modalName='タグを追加'">+</b-button>
    <b-modal id="editTag" v-bind:title="this.modalName">
      <form class="">
        タグ名
        <b-form-input type="text" placeholder="タグ名" v-model="tagName"/>
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
      parent:null,
    }
  },
  computed:{
    tags(){
      this.$store.state.tag_list
    }
  },
  method:{
    updateTag(){
      if()
    }
  },
  mounted:function(){
    this.$store.dispatch("getTagList")
  }
}
</script>

<style lang="css" scoped>
</style>

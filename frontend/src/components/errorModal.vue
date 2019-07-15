<template>
  <div>
    <b-modal v-model="isError" title="Error!" hide-footer>
      <div class="d-block text-center">
        <p>{{message}}</p>
      </div>
      <b-button class="mt-3" variant="outline-danger" block @click="checked">OK</b-button>
    </b-modal>
  </div>
</template>

<script>
  export default {
    computed:{
      isError(){
        return this.$store.state.error_show
      },
      message(){
        return this.$store.state.error_message
      }
    },
    methods: {
      checked() {
        this.$store.commit('setErrorModalStatus', false)
        this.$store.commit('setErrorMessage', '')
      }
    },
    mounted:function(){
      this.$store.commit('setErrorModalStatus', false)
    }
  }
</script>

<template>
  <div class="errMOD">
    <b-modal ref="err-modal" hide-footer title="Error Message">
      <div class="d-block text-center">
        <h3>Error</h3>
        <p>{{errmes}}</p>
      </div>
      <b-button class="mt-3" variant="outline-danger" block @click="hideModal">OK</b-button>
    </b-modal>
  </div>
</template>

<script>
  export default {
    name : 'errMOD',
    data : function(){
      return{
        errmes : ''
      }
    },
    methods: {
      showModal() {
        this.$refs['err-modal'].show()
      },
      hideModal() {
        this.$refs['err-modal'].hide()
      }
    }
  }
</script>

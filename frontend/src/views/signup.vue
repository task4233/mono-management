<template>
  <b-container class="signup">
    <h1>mono management</h1>
    <b-form-group>
      <b-button variant="outline-success" class="mx-auto" href="/login">login</b-button>
    </b-form-group>
    <accountForm v-on:formInput="formInput" password-required="1" request-retype="1"></accountForm>
    <b-form-group>
      <b-button v-on:click="signup" variant="primary" class="mx-auto">サインアップ</b-button>
    </b-form-group>
    <b-row class="errorForm">
      <p v-if="error" v-html="error" class="mx-auto"></p>
    </b-row>
  </b-container>
</template>

<script>
import axios from 'axios'
import router from '../router'
import accountForm from '../components/accountForm.vue'
export default {
  name: 'signup',
  components: {
    accountForm
  },
  data: function() {
    return {
      signupId: '',
      signupPass: '',
      signupPassRetype: '',
      formError: '何も入力されていません',
      error : ''
    }
  },
  methods: {
    formInput: function(v) {
      this.signupId = v.userId
      this.signupPass = v.userPass
      this.signupPassRetype = v.userPassRetype
      this.formError = v.error
    },
    signup: function() {
      var self = this;
      this.error = this.formError;
      if (this.error === '') {
        var data = {name : this.signupId, password : this.signupPass };
        axios.post('/api/v1/user/new', data)
          .then(response => {
            console.log(response)
            if (response.data.status) {
              console.log('body:', response.data);
              router.push('/')
            } else {
              self.error = response.message
            }
          }).catch(function(error) {
            console.log(error);
            self.error = String(error.response.status); // Vueの中にaxiosが入れ子になっているため、参照できない => thisを変数selfにする
            if (self.error === '404') self.error = 'サーバとの接続が確立できませんでした。'
            console.log(error.response.data.message)
          });
      }
    }
  },
  created:function(){
    this.$axios.get('/api/v1/user/info')
    .then(function(response){
      if(response.data.status){
        router.push('/')
      }
    })
  }
}
</script>

<style>
.signup {
  font-family: 'Makinas-4-Square';
}

</style>

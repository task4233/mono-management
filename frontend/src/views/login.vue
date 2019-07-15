<template>
  <b-container class="login">
    <h1>mono management</h1>
    <b-form-group>
      <b-button variant="outline-success" class="mx-auto" href="/signup">signup</b-button>
    </b-form-group>
    <accountForm v-on:formInput="formInput" password-required="1" request-retype=""></accountForm>
    <b-form-group>
      <b-button v-on:click="login" variant="primary" class="mx-auto">ログイン</b-button>
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
  name: 'login',
  components: {
    accountForm
  },
  data: function() {
    return {
      loginId: '',
      loginPass: '',
      error : '',
      formError: '何も入力されていません'
    }
  },
  methods: {
    formInput: function(v) {
      this.loginId = v.userId
      this.loginPass = v.userPass
      this.formError = v.error
    },
    login: function() {
      var self = this;
      this.error = this.formError;
      if (this.error === '') {
        var data = {name : this.loginId, password : this.loginPass };
        axios.post('/api/v1/user/login', data)
          .then(response => {
            console.log(response)
            if (response.data.status) {
              console.log('body:', response.data);
              router.push('/')
            } else {
              self.error = response.message
            }
          }).catch(function(error) {
            console.log(error); // 通信エラーをコンソールに表示
            self.error = String(error.response.status); // Vueの中にaxiosが入れ子になっているため、参照できない => thisを変数selfにする
            if (self.error === '404') self.error = 'サーバとの接続が確立できませんでした。'
            console.log(error.response.data.message)
          });
      }
    },
  },
  created:function(){
    this.$axios.get('/api/v1/user/info')
    .then(function(response){
      if(response.data.status){
        this.$router.push('/')
      }
    })
  }
}
</script>
<style >
.login {
  font-family: 'Makinas-4-Square';
}

</style>

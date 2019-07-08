<template>
  <b-container class="signup">
    <h1>mono management</h1>
    <b-form-group>
      <b-button variant="outline-success" class="mx-auto" href="/">login</b-button>
    </b-form-group>
    <b-form-group>
      <div class="signupId">
        <b-form-input type="text" placeholder="ユーザ名" v-model="signupId"></b-form-input>
      </div>
    </b-form-group>
    <b-form-group>
      <div class="signupPass">
        <b-form-input type="password" placeholder="パスワード" v-model="signupPass"></b-form-input>
      </div>
    </b-form-group>
    <b-form-group>
      <div class="signupPassRetype">
        <b-form-input type="password" placeholder="パスワード再入力" v-model="signupPassRetype"></b-form-input>
      </div>
    </b-form-group>
    <b-form-group>
      <b-button v-on:click="signup" variant="primary" class="mx-auto">サインアップ</b-button>
    </b-form-group>
    <b-row class="errorForm">
      <p v-if="error" class="mx-auto">{{ error }}</p>
    </b-row>
  </b-container>
</template>

<script>
import axios from 'axios'
export default {
  name: 'signup',
  data: function() {
    return {
      signupId: '',
      signupPass: '',
      signupPassRetype: '',
      error : '',
      flag : 0, //  flag変数をloginされるたびに変えて、watchを呼び出す。
      server : 0
    }
  },
  watch : {
    flag: function () {
      this.error = ''; // errorの初期化
      if (this.signupId === '') {
        this.error = this.error + 'ユーザ名が入力されていません。'
      }
      if (this.signupId.length > 64) {
        this.error = this.error + 'ユーザ名が長すぎます。'
      }
      if (this.signupPass === '') {
        this.error = this.error + 'パスワードが入力されていません。'
      }
      if (this.signupPass !== this.signupPassRetype) {
        this.error = this.error + 'パスワード再入力が一致していません。'
      }
      if (this.signupPass.length > 64) {
        this.error = this.error + 'パスワードが長すぎます。'
      }
      if (this.server) {
        this.error = this.error + this.server
      }
    }
  },
  methods: {
    signup: function() {
      var self = this;
      this.flag++;
      if (this.signupId !== '' && this.signupPass !== '' && this.signupId.length <= 64 && this.signupPass.length <= 64) {
        var data = {name : this.signupId, password : this.signupPass };
        axios.post('/api/v1/user/new', data)
          .then(response => {
            if (response.data.Status) {
              console.log('body:', response.data);
              this.$router.push('/')
            } else {
              self.server = response.message
              self.flag = 0
            }
          }).catch(function(error) {
            console.log(error);
            self.server = error.response.data.message
            self.flag = 0;
          });
      }
    }
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

<style>
.signup {
  font-family: 'Makinas-4-Square';
}

.errorform {
  /*margin: 0 auto;
  width: 250px;
  height: 200px;
  background-color: #f0f0f0;*/
}
</style>

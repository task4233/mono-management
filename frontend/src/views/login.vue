<template>
  <b-container class="login">
    <h1>mono management</h1>
    <b-form-group>
      <b-button variant="outline-success" class="mx-auto" href="/signup">signup</b-button>
    </b-form-group>
    <b-form-group>
      <div class="loginId">
        <b-form-input type="text" placeholder="ユーザ名" v-model="loginId"></b-form-input>
      </div>
    </b-form-group>
    <b-form-group>
      <div class="loginPass">
        <b-form-input type="password" placeholder="パスワード" v-model="loginPass"></b-form-input>
      </div>
    </b-form-group>
    <b-form-group>
      <b-button v-on:click="login" variant="primary" class="mx-auto">ログイン</b-button>
    </b-form-group>
    <b-row class="errorForm">
      <p v-if="error" class="mx-auto">{{ error }}</p>
    </b-row>
  </b-container>
</template>

<script>
import axios from 'axios'
import router from '../router'

export default {
  name: 'login',
  data: function() {
    return {
      loginId: '',
      loginPass: '',
      error : '',
      flag : 0, //  flag変数をloginされるたびに変えて、watchを呼び出す。
      server : 0
    }
  },
  watch : {
    flag: function () {// 文字制限64
      this.error = ''; // errorの初期化
      if (this.loginId === '') {
        this.error = this.error + 'ユーザ名が入力されていません。'
      }
      if (this.loginId.length > 64) {
        this.error = this.error + 'ユーザ名が長すぎます。'
      }
      if (this.loginPass === '') {
        this.error = this.error + 'パスワードが入力されていません。'
      }
      if (this.loginPass.length > 64) {
        this.error = this.error + 'パスワードが長すぎます。'
      }
      if (this.flag === -1) {
        this.error = this.error + this.server
        this.server = '';
      }
    }
  },
  methods: {
    login: function() {
      var self = this;
      this.flag++;
      this.server = "";
      if (this.loginId !== '' && this.loginPass !== '' && this.loginId.length <= 64 && this.loginPass.length <= 64) {
        var data = {name : this.loginId, password : this.loginPass };
        axios.post('/api/v1/user/login', data)
          .then(function(response){
            console.log('body:', response.data); // サーバに送信したデータをコンソールに表示
            router.push('/')
          }).catch(function(error) {
            console.log(error); // 通信エラーをコンソールに表示
            self.server = error.response.status; // Vueの中にaxiosが入れ子になっているため、参照できない => thisを変数selfにする
            if (self.server === 404) self.server = 'サーバとの接続が確立できませんでした。'
            else self.server = error;
            self.flag  = -1;
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

.errorForm {
  /*margin: 0 auto;
  width: 250px;
  height: 200px;
  background-color: #f0f0f0;*/
}
</style>

<template>
  <div class="login">
    <font size="9" face="ＭＳ 明朝" color="#ffffff">
      mono<br>management
    </font>
    <br>
    <router-link to="/signup">signup</router-link>
    <br>
    <font size="2">
    ユーザー名/メールアドレス
    <div class="loginId">
      <input type="text" placeholder="ログインID" v-model="loginId">
    </div>
    パスワード
    </font>
    <div class="loginPass">
      <input type="text" placeholder="ログインPASS" v-model="loginPass">
    </div>
    <button v-on:click="login">ログイン</button>
    <div class = "errorForm">
      <p v-if="error">{{ error }}</p>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'login',
  data: function() {
    return {
      loginId: '',
      loginPass: '',
      users: [],
      error : '',
      flag : 0, //  flag変数をloginされるたびに変えて、watchを呼び出す。
    }
  },
  watch : {
    flag: function (value) {
      if (this.loginId === '') {
        this.error = this.error + '\n' + 'ログインIDが入力されていません。'
      }
      if (this.loginPass === '') {
        this.error = this.error + '\n' + 'ログインPASSが入力されていません。'
      }
      if (value === 404) {
        this.error = this.error + '\n' + 'サーバに接続できませんでした。'
      }
    }
  },
  methods: {
    login: function() {
      var self = this;
      this.error = ''; // errorの初期化
      this.flag++;
      if (this.loginId !== null && this.loginPass !== null) {
        var data = {name : this.loginId, password : this.loginPass };
        axios.post('/api/v1/user/login', data)
          .then(response => {
            console.log('body:', response.data); // サーバに送信したデータをコンソールに表示
          }).catch(function(error) {
            console.log(error); // 通信エラーをコンソールに表示
            self.flag = 404; // Vueの中にaxiosが入れ子になっているため、参照できない => thisを変数selfにする
          });
      }
    },
  }
}
</script>

<style scoped>
.login {
  margin: 0 auto;
  width: 300px;
  height: 500px;
  background-color: #6fdf6f;
}

.errorForm {
  margin: 0 auto;
  width: 250px;
  height: 200px;
  background-color: #f0f0f0;
}
</style>

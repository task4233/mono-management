<template>
  <div class="signup">
    <font size="9" face="ＭＳ 明朝" color="#ffffff">
      mono<br>management
    </font>
    <br>
    <router-link to="/">login</router-link>
    <br>
    <font size="2">
    ユーザー名/メールアドレス
    <div class="signupId">
      <input type="text" placeholder="サインアップID" v-model="signupId">
    </div>
    パスワード
    </font>
    <div class="signupPass">
      <input type="text" placeholder="サインアップPASS" v-model="signupPass">
    </div>
    <button @click="signup">サインアップ</button>
    <div class = "errorform">
      <p v-if="error">{{ error }}</p>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'signup',
  data: function() {
    return {
      signupId: '',
      signupPass: '',
      error : '',
      flag : 0, //  flag変数をloginされるたびに変えて、watchを呼び出す。
    }
  },
  watch : {
    flag: function (value) {
      this.error = ''; // errorの初期化
      if (this.signupId === '') {
        this.error = this.error + 'サインアップIDが入力されていません。'
      }
      if (this.signupId.length > 255) {
        this.error = this.error + 'サインアップIDの文字数が長すぎます。'
      }
      if (this.signupPass === '') {
        this.error = this.error + 'サインアップPASSが入力されていません。'
      }
      if (this.signupPass.length > 255) {
        this.error = this.error + 'サインアップPASSの文字数が長すぎます。'
      }
      if (value === 404) {
        this.error = this.error + 'サーバに接続できませんでした。'
      }
    }
  },
  methods: {
    signup: function() {
      var self = this;
      this.flag++;
      if (this.signupId !== null && this.signupPass !== null) {
        var data = {name : this.signupId, password : this.signupPass };
        axios.post('/api/v1/user/new', data)
          .then(response => {
            console.log('body:', response.data);
          }).catch(function(error) {
            console.log(error);
            self.flag = 404;
          });
      }
    }
  }
}
</script>

<style>
.signup {
  margin: 0 auto;
  width: 300px;
  height: 500px;
  background-color: #6fdf6f;
}

.errorform {
  margin: 0 auto;
  width: 250px;
  height: 200px;
  background-color: #f0f0f0;
}
</style>

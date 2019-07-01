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
    <div class="loginId">
      <input type="text" placeholder="ログインID" v-model="loginId">
    </div>
    パスワード
    </font>
    <div class="loginPass">
      <input type="text" placeholder="ログインPASS" v-model="loginPass">
    </div>
    <button @click="signup">サインアップ</button>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'signup',
  data: function() {
    return {
      loginId: '',
      loginPass: '',
    }
  },
  methods: {
    signup: function() {
      /*特別な $event 変数を使うことでメソッドに DOM イベントを渡すことができます*/
      // checkloginイベント(htmlファイル内で定義している)の内容を記述
        // document.getElementById('loginResult').innerHTML = 'Login Success !'
      if (this.userId !== null && this.userPass !== null) {
        var data = {name : this.loginId, password : this.loginPass };
        axios.post('/api/v1/user/new', data)
          .then(response => {
            console.log('body:', response.data);
          }).catch(function(error) {
            console.log(error);
        });
      } else {
        // エラーメッセージ作成
      }
    }
  }
}
</script>

<style scoped>
.signup {
  margin: 0 auto;
  width: 300px;
  height: 500px;
  background-color: #6fdf6f;
}
</style>

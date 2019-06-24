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
    </font>
    <div class="loginId">
      <input type="text" placeholder="ログインID" v-model="loginId">
    </div>
    パスワード
    <div class="loginPass">
      <input type="text" placeholder="ログインPASS" v-model="loginPass">
    </div>
    googleアカウントでログイン
    <br>
    <button v-on:click="login">ログイン</button>
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
    }
  },
  created() {
    axios.get('http://localhost:3000/users').then(response => {
      console.log('status:', response.status);
      console.log('body:', response.data);
      this.users = response.data
    });
  },
  methods: {
    login: function() {
      /*特別な $event 変数を使うことでメソッドに DOM イベントを渡すことができます*/
      var nextPage = this.$route.query.next
      if (nextPage === undefined) {
        nextPage = 'login'
      }
      console.log(nextPage)
      // checkloginイベント(htmlファイル内で定義している)の内容を記述
      for (var i = 0; this.users.length < 10; i++) {
        if (this.users[i] === null) break
        if (this.loginPass === this.users[i].userPass &&
          this.loginId === this.users[i].userId) {
          this.$router.push('/header')
        } else {
          // errorメッセージを作成
        }
      }
    }
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

</style>

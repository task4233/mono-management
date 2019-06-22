<template>
  <div class="login">
    <font size="9" face="ＭＳ 明朝" color="#ffffff">
      mono<br>management
    </font>
    <br>
    <router-link to="sigup">signup</router-link>
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
    <button @click="login">ログイン</button>
  </div>
</template>

<script>
export default {
  name: 'login',
  data: function() {
    return {
      loginId: '',
      loginPass: '',
    }
  },
  methods: {
    login: function(event) {
      /*特別な $event 変数を使うことでメソッドに DOM イベントを渡すことができます*/
      var nextPage = this.$route.query.next
      if (nextPage === undefined) {
        nextPage = 'login'
      }
      console.log(nextPage)
      // checkloginイベント(htmlファイル内で定義している)の内容を記述
      var correctLoginId = 'LoginId'
      var correctPass = 'LoginPass'
      var myLoginId = document.getElementById('loginId').value
      var myLoginPass = document.getElementById('loginPass').value
      if ((myLoginId === correctLoginId) &&
        (myLoginPass === correctPass)) {
        // document.getElementById('loginResult').innerHTML = 'Login Success !'
        this.$router.push({name: nextPage, query: { auth: 'authenticated' }})
      } else {
        document.getElementById('loginResult').innerHTML = 'Login Failed !'
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

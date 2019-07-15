<template lang="html">
  <div>
    <b-form-group>
      <b-form-input type="text" placeholder="ユーザ名" v-model="userId" v-on:input="onInput"></b-form-input>
    </b-form-group>
    <b-form-group>
      <b-form-input type="password" placeholder="パスワード" v-model="userPass" v-on:input="onInput"></b-form-input>
    </b-form-group>
    <b-form-group v-if="requestRetype">
      <b-form-input type="password" placeholder="パスワード再入力" v-model="userPassRetype" v-on:input="onInput"></b-form-input>
    </b-form-group>
  </div>
</template>
<script>
export default {
  data: function() {
    return {
      userId: '',
      userPass: '',
      userPassRetype: '',
      error: ''
    }
  },
  props: [
    'passwordRequired',
    'requestRetype'
  ],
  methods: {
    onInput: function() {
      this.error = ''
      if (!this.userId) {
        this.error += 'ユーザ名が入力されていません。<br />'
      }
      if (64 < this.userId.length) {
        this.error += 'ユーザ名が長すぎます。<br />'
      }
      if (this.passwordRequired && !this.userPass) {
        this.error += 'パスワードが入力されていません。<br />'
      }
      if (this.requestRetype && this.userPass !== this.userPassRetype) {
        this.error += 'パスワードの再入力が一致していません。<br />'
      }
      this.$emit('formInput', {
        userId: this.userId,
        userPass: this.userPass,
        userPassRetype: this.userPassRetype,
        error: this.error
      })
    }
  }
}
</script>

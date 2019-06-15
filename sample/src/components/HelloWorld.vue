<template>
  <div class="hello">
  <button v-on:click=getData>Add 1</button>
  <p>{{ message }}</p>
</div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'HelloWorld',
  data: function () {
      return {
        message: 0, // 初期値
      }
    },
    methods: {
      getData: function () {
      // axiosはmain.jsやbootstrap.jsで定義しておく
      axios.get('/api/ping')
      .then(res => {
        // console.log(res)
        this.message = res.data.name; // apiで取得したuserIdをdataのuserIdにセット
      })
      .catch(error => {
        console.log(error.response)
      });
      }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>

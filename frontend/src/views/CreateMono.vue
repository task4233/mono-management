<template>
  <div class="signup">
    <h1>新規作成</h1>
    <input type="text" placeholder="名前" v-model="name" />
    <input type="text" placeholder="タグ" v-model="tagId" />
    <button @click="create">保存する</button>
  </div>
</template>

<script>
import Axios from "axios";

export default {
  name: 'CreateMono',
  data: function() {
    return {
      name: "",
      tagId: "",
      data: [""]
    };
  },
  methods: {
    create: function() {
      Axios.post("0.0.0.0:80/api/v1/mono/new", {
        name: this.name,
        tagId: this.tagId,
        data: [
          { name: "数量", value: "3", type: "num" },
          {
            name: "消費期限",
            value: "Fri, 28 Jun 2019 08:09:52 GMT",
            type: "timestamp"
          },
          { name: "メモ", value: "赤いよ", type: "str" }
        ]
      }).then(response => {
        console.log(response)
        const status = response.data.status
        alert(response.data.message)
        if (status) {
          this.$router.push('/')
        }
      }).catch(err => {
        console.log(err.response)
      })
    }
  }
};
</script>

<style scoped>
h1,
h2 {
  font-weight: normal;
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
.signup {
  margin-top: 20px;
  display: flex;
  flex-flow: column nowrap;
  justify-content: center;
  align-items: center;
}
input {
  margin: 10px 0;
  padding: 10px;
}
</style>
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
import qs from "qs";

export default {
  name: "CreateMono",
  methods: {
    create: function() {
      const data = {
        name: String(this.name),
        tagId: Number(this.tagId)
      }
      console.log(data)


      Axios.post("/api/v1/mono/new", data, {
        headers: {
          "Content-Type": "application/json",
          "Access-Control-Allow-Origin": "*",
          "Access-Control-Allow-Headers": "*",
          useCredentials: true
        }
      })
        .then(response => {
          console.log(response);
          const status = response.data.status;
          alert(response.data.message);
          if (status) {
            this.$router.push("/");
          }
        })
        .catch(err => {
          console.log(err.response);
        });
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
<template>
  <div class="home">
    <h1 v-if="itemdatas.length === 0">モノが1つも無いようです</h1>
    <p v-else v-for="dat in itemdatas" :key="dat.name">
      <router-link :to="'/mono/' + dat.id">{{ dat.name }}</router-link>
    </p>
    <!-- <b-button size="lg" variant="outline-primary" @click="rerun">リロード用ボタン</b-button> -->
  </div>
</template>

<script>
import Axios from "axios";

export default {
  name: "Home",
  data() {
    return {
      itemdatas: []
    };
  },
  created: function() {
    Axios.get("/api/v1/mono/")
      .then(response => {
        const stat = response.data.Status;
        if (stat) {
          const datas = response.data.Data;
          for (let i = 0; i < datas.length; ++i) {
            const mid = datas[i].Id;
            const mname = datas[i].name;
            const mtagId = datas[i].tagId;
            const muserId = datas[i].userId;
            this.itemdatas.push({
              id: mid,
              name: mname,
              tagId: mtagId,
              userId: muserId
            })
          }
        }
      })
      .catch(err => {
        console.log(err.response);
      });
  }
}
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
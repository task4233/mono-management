<template>
  <div class="signup">
    <h1>新規作成</h1>
    <p>
      <input type="text" placeholder="名前" v-model="name" />
      <input type="text" placeholder="タグ" v-model="tagId" />
    </p>

    <div class="dynamic">
      <p v-for="dat in data" :key="dat.name">
        {{ dat.name }}
        {{ dat.type }}
        {{ dat.value }}
      </p>
      <p>
        <input type="text" placeholder="新しい要素名を追加！" v-model="dataName" />
        <select v-model="dataType">
          <option disabled value>型を選んでね</option>
          <option>num</option>
          <option>str</option>
          <option>timestamp</option>
        </select>
        <datepicker v-if="dataType=='timestamp'" v-model="dataValue" :format="DatePickerFormat"></datepicker>
        <input v-else type="text" placeholder="新しい要素の値を追加！" v-model="dataValue" />
      </p>
    </div>
    <button @click="addData" class="btn btn-primary btn-sm">+</button>
    <button @click="create">保存する</button>
  </div>
</template>

<script>
import Axios from "axios";
import Datepicker from "vuejs-datepicker"

export default {
  name: "CreateMono",
  data() {
    return {
      dataName: "",
      dataValue: "",
      dataType: "",
      DatePickerFormat: 'yyyy-MM-dd',
      data: []
    };
  },
  components: {
    Datepicker
  },
  methods: {
    addData: function() {
      const newDataName = this.dataName.trim();
      if (!newDataName) {
        return;
      }
      const newDataType = this.dataType.trim();
      if (!newDataType) {
        return;
      }
      const newDataValue = newDataType==='timestamp' ? this.dataValue.toUTCString() : this.dataValue.trim();
      if (!newDataValue) {
        console.log(this.dataValue)
        return;
      }
      this.data.push({
        name: newDataName,
        type: newDataType,
        value: newDataValue
      });
      alert("追加するよ!");
      this.dataName = " ";
      this.dataType = " ";
      this.dataValue = " ";
    },
    create: function() {
      const data = {
        name: String(this.name),
        tagId: Number(this.tagId),
        data: this.data
      };
      console.log(data);

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
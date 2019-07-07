<template>
  <div class="signup">
    <h1>新規作成</h1>
    <p>
      <b-form-input type="text" placeholder="名前" v-model="name" />
      <b-form-select placeholder="タグ" v-model="tagId">
        <option :value="null">タグを選択</option>
        <option v-for="tag in tags" value="tag.tagId">{{tag.name}}</option>
      </b-form-select>
    </p>

    <b-form-group class="dynamic">
      <p v-for="dat in data" :key="dat.name">
        {{ dat.name }}
        {{ dat.type }}
        {{ dat.value }}
      </p>
      <p>
        <b-form-input type="text" placeholder="新しい要素名を追加！" v-model="dataName" />
        <b-form-select v-model="dataType">
          <option disabled value>型を選んでね</option>
          <option>num</option>
          <option>str</option>
          <option>timestamp</option>
        </b-form-select>
        <datepicker v-if="dataType=='timestamp'" v-model="dataValue" :format="DatePickerFormat"></datepicker>
        <b-form-input v-else type="text" placeholder="新しい要素の値を追加！" v-model="dataValue" />
      </p>
    </b-form-group>
    <b-form-group>
      <b-button @click="addData" variant="outline-primary" pill>+</b-button>
    </b-form-group>
    <b-form-group>
      <b-button @click="create" variant="primary">保存する</b-button>
    </b-form-group>
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
      data: [],
      tagId:null,
      name:""
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
  },
  computed:{
    tags(){
      return this.$store.state.tag_list
    }
  },
  created(){
    this.$store.dispatch("getTagList")
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

<template>
  <div class="signup">
    <h1>新規作成</h1>
    <p>
      <b-form-input type="text" placeholder="名前" v-model="name" />
      <tagList></tagList>
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
      <b-button @click="addData" variant="outline-primary" pill>要素を追加</b-button>
    </b-form-group>
    <b-form-group>
      <b-button @click="create" variant="primary">保存する</b-button>
    </b-form-group>
    <b-row class="errorForm">
      <p v-if="error" class="mx-auto">{{ error }}</p>
    </b-row>
  </div>
</template>

<script>
import Axios from "axios";
import Datepicker from "vuejs-datepicker";
import tagList from "../components/tag-list.vue";

export default {
  name: "CreateMono",
  data() {
    return {
      name:"",
      dataName: "",
      dataValue: "",
      dataType: "",
      DatePickerFormat: "yyyy-MM-dd",
      data: [],
      error: ""
    };
  },
  components: {
    Datepicker,
    tagList
  },
  methods: {
    addData: function() {
      const newDataName = this.dataName.trim();
      if (!newDataName) {
        this.error = "データ名が入力されていません。";
        return;
      }
      if (newDataName.length > 64) {
        this.error = this.error + 'データ名が長すぎます。'
        return
      }
      const newDataType = this.dataType.trim();
      if (!newDataType) {
        this.error = "データの型が選択されていません。";
        return
      }
      const newDataValue =
        newDataType === "timestamp"
          ? this.dataValue.toUTCString()
          : this.dataValue.trim();
      if (!newDataValue) {
        console.log(this.dataValue);
        this.error = "データの値が入力されていません。";
        return;
      }
      if (newDataValue.length > 64) {
        this.error = this.error + 'データの値が長すぎます。'
        return
      }
      if (newDataType=="num") {
        const pattern = /^[-]?(\d+)(\.\d+)?$/;
        if (!pattern.test(newDataValue)) {
          this.error = "数値を入力してください"
          return
        }
      }

      this.data.push({
        name: newDataName,
        type: newDataType,
        value: newDataValue
      });
      alert("追加するよ");
      this.dataName = "";
      this.dataType = "";
      this.dataValue = "";
    },
    create: function() {
      const data = {
        name: String(this.name),
        tagId: Number(this.$store.state.select_tag),
        data: this.data
      };
      console.log(data);

      if (!data.name.trim()) {
        this.error = "mono名が入力されていません。";
        return;
      }

      if (!data.tagId) {
        this.error = "tagが選択されていません。";
        return;
      }

      if (this.dataName.trim() || this.dataValue) {
        if (!(!this.dataName.trim() && this.dataType==="timestamp")) {
          this.error = "+ボタンを押して情報を追加してください。"
          return;
        }
      }

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
.errorForm {
  /*margin: 0 auto;
  width: 250px;
  height: 200px;
  background-color: #f0f0f0;*/
}
</style>

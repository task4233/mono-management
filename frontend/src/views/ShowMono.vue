<template>
  <div class="edit">
    <h1>{{ itemdatas[$route.params.id].name }}詳細画面</h1>
    <p >
        
        </p>
        </div>
</template>

<script>
import Axios from "axios";
import Datepicker from "vuejs-datepicker";

export default {
  name: "CreateMono",
  data() {
    return {
      dataName: "",
      dataValue: "",
      dataType: "",
      DatePickerFormat: "yyyy-MM-dd",
      data: [],
      itemdatas: [],
    };
  },
  components: {
    Datepicker
  },
  created: function() {
    Axios.get("/api/v1/mono/")
      .then(response => {
        console.log(response);
        const stat = response.data.Status;
        if (stat) {
          const datas = response.data.Data;
          console.log(datas);
          console.log(datas.length);
          for (let i = 0; i < datas.length; ++i) {
            const mid = datas[i].Id;
            const mname = datas[i].name;
            const mtagId = datas[i].tagId;
            const muserId = datas[i].userId;
            console.log({
              id: mid,
              name: mname,
              tagId: mtagId,
              userId: muserId
            });
            this.itemdatas.push({
              id: mid,
              name: mname,
              tagId: mtagId,
              userId: muserId
            });
          }
        }
      })
      .catch(err => {
        console.log(err.response);
      });
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
      const newDataValue =
        newDataType === "timestamp"
          ? this.dataValue.toUTCString()
          : this.dataValue.trim();
      if (!newDataValue) {
        console.log(this.dataValue);
        return;
      }
      this.data.push({
        name: newDataName,
        type: newDataType,
        value: newDataValue
      });
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

      Axios.put("/api/v1/mono/"+this.$route.params.id, data, {
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
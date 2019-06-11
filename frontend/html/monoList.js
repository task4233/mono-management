Vue.component('mono-item',{
  props:['mono'],
  template:'<tr><td>{{mono.name}}</td><td>{{mono.info[0]}}</td><td>{{mono.info[1]}}</td></tr>'
})

Vue.component('mono-head',{
  props:['tag'],
  template:'<tr><th>名称</th><th>{{tag.info[0]}}</th><th>{{tag.info[1]}}</th></tr>'
})

var app = new Vue({
  el:'#monoList',
  data:{
    //monoList,Tagの中身はAPIから動的に受け取る．
    monoList:[
      {id: 0, name:'トマト', info:[3,'2019/06/30']},
      {id: 1, name:'ほうれん草', info:[5, '2019/5/15']},
      {id: 5, name:'ピーマン', info:[4,'2019/6/15']},
    ],
    Tag:{id:11, name:'冷蔵庫', info:['個数','賞味期限']}
  },
  template:'<table class="table table-sprited"><mono-head v-bind:tag = "Tag"></mono-head><mono-item v-for="item in monoList" v-bind:mono = "item" v-bind:key = "item.id"></mono-item></table>'

})

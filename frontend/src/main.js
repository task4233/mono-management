import Vue from 'vue'
import Router from 'vue-router'
import App from './App.vue'
import Hello from '@/components/HelloWorld'

Vue.config.productionTip = false

Vue.use(Router)

export default new Router({
    routes: [
        {
            path: '/',
            name: 'Hello',
            component: Hello
        }
        ]
})

new Vue({
  render: h => h(App),
}).$mount('#app')

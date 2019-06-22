import Vue from 'vue'
import Router from 'vue-router'
// import HelloWorld from '@/components/HelloWorld'
import Login from '../components/login.vue'
import SignUp from '../components/signup.vue'

Vue.use(Router)

export default new Router({
  mode: 'history', //  https://router.vuejs.org/guide/essentials/history-mode.html
  routes: [
    {
      path: '/',
      /* name: 'top', */
      component: Login
    },
    {
      path: 'login',
      name: 'login',
      query: { next: '' },
      component: Login
    },
    {
      path: '/signup',
      name: 'signup',
      query: { auth: '' },
      component: SignUp
    }
  ]
})

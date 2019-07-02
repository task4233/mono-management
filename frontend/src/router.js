import Vue from 'vue'
import Router from 'vue-router'
import Login from './components/login.vue'
import SignUp from './components/signup.vue'
import Header from './components/HelloWorld.vue'
import Home from './views/Home.vue'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/signup',
      name: 'signup',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: SignUp
    },
    {
      path: '/header',
      name: 'header',
      component: Header
      component: () => import(/* webpackChunkName: "about" */ './views/About.vue')
    },
    {
      path: '/mono/new',
      name: 'createmono',
      component: () => import('@/views/CreateMono.vue')
    }
  ]
})

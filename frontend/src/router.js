import Vue from 'vue'
import Router from 'vue-router'
import Index from './views/Index.vue'
import EditMono from './views/EditMono.vue'
import CreateMono from './views/CreateMono.vue'
import Login from './views/login.vue'
import SignUp from './views/signup.vue'
import EditTags from './views/EditTags.vue'
import NotFound from './views/404Page.vue'
Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Index
    },
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
      path: '/mono/new',
      name: 'createmono',
      component: CreateMono
    },
    {
      path: '/mono/:id',
      name: 'editmono',
      component: EditMono
    },
    {
      path:'/EditTags',
      name:'EditTags',
      component: EditTags
    },
    {
      path:'*',
      name:'NotFound',
      component: NotFound
    }
  ]
})

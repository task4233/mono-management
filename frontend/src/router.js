import Vue from 'vue'
import Router from 'vue-router'
import Index from './views/Index.vue'
import EditMono from './views/EditMono.vue'
import CreateMono from './views/CreateMono.vue'

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
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import(/* webpackChunkName: "about" */ './views/About.vue')
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
  ]
})

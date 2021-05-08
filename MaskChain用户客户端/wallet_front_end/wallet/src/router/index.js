import Vue from 'vue'
import VueRouter from 'vue-router'
import Newwallet from '../views/Newwallet.vue'
import Main from '../views/Main.vue'
import Loadwallet from '../views/Loadwallet.vue'
import Mainaction from '../views/Mainaction.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Main',
    component: Main
  },
  {
    path: '/Newwallet',
    name: 'Newwallet',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: Newwallet
  },
  {
    path: '/Loadwallet',
    name: 'Loadwallet',
    component: Loadwallet
  },
  {
    path: '/Mainaction',
    name: 'Mainaction',
    component: Mainaction
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router


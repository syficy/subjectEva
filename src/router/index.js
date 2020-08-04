import Vue from 'vue'
import VueRouter from 'vue-router'



Vue.use(VueRouter)



const routes = [
  {
    path: '/',
    redirect: '/home',
    meta: {
      title: 'zstu课程评价'
    },
    component: () => import('../views/home/Home'),
  },
  {
    path: '/home',
    /*redirect: '/home',*/
    meta: {
      title: 'zstu课程评价'
    },
    component: () => import('../views/home/Home')
  },
  {
    path: '/Search',
    meta: {
      title: 'zstu课程评价'
    },
    /*redirect: '/Search',*/
    component: () => import('../views/Search/Search')
  },
  {
    path: '/Upload',
    meta: {
      title: 'zstu课程评价'
    },
    /*redirect: '/Upload',*/
    component: () => import('../views/Upload/Upload')
  },
  {
    path: '/profile',
    meta: {
      title: 'zstu课程评价'
    },
    /*redirect: '/profile',*/
    component: () => import('../views/profile/Profile')
  },
]

const router = new VueRouter({
  routes,
  mode: 'history'
})





export default router

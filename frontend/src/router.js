import Vue from 'vue'
import Router from 'vue-router'
import SignUp from './views/SignUp.vue'
import SignIn from './views/SignIn.vue'
import VoteKnowledge from './views/VoteKnowledge.vue'
import VoteResult from './views/VoteResult.vue'
import VoteResultById from './views/VoteResultById.vue'
import store from '@/store'

Vue.use(Router)

const router = new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/signin',
      name: 'signin',
      component: SignIn,
    },
    {
      path: '/signup',
      name: 'signup',
      component: SignUp,
    },
    {
      path: '/',
      name: 'voteknowledge',
      component: VoteKnowledge,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/kekka',
      name: 'voteresult',
      component: VoteResult,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/kekka/:id',
      name: 'voteresultbyid',
      component: VoteResultById,
      meta: {
        requiresAuth: true
      }
    },
  ]
})

export default router;

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    // すでに認証済みの場合は遷移
    // if (store.getters['auth/isSignIn'].isSignIn) {
    if (localStorage.getItem("token") != "") {
      store.commit("auth/setSignIn", true);
      next()
      return
    } else {
      next('/signin')
    }
  }
  next()
})
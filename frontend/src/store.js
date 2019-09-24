import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    auth: {
      namespaced: true,
      state: {
        isSignIn: false
      },
      mutations: {
        setSignIn(state, isSignIn) {
          state.isSignIn = isSignIn;
        }
      },
      getters: {
        isSignIn(state) {
          return state.isSignIn;
        }
      },
      actions: {
      }
    }
  }
})

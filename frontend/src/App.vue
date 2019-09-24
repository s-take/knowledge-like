<template>
  <v-app>
    <v-app-bar app color="blue darken-3" dark>
      <v-toolbar-title class="headline text-uppercase">
        <span>Voting for BEST Knowledge</span>
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn icon color="white lighten-2">
        <router-link v-bind:to="{ name : 'voteresult' }"><v-icon>note</v-icon></router-link>
      </v-btn>
      <v-btn icon color="white lighten-2" @click="signout" v-if="isSignIn">
        <v-icon>exit_to_app</v-icon>
      </v-btn>
    </v-app-bar>

    <v-content>
      <router-view />
    </v-content>
  </v-app>
</template>

<script>
import { mapGetters } from "vuex";
import store from "@/store";

export default {
  name: "App",
  data() {
    return {
      //
    };
  },
  computed: {
    ...mapGetters({
      isSignIn: "auth/isSignIn"
    })
  },
  methods: {
    signout: function() {
      store.commit("auth/setSignIn", false);
      localStorage.removeItem("token");
      this.$router.push("/signin");
    }
  }
};
</script>

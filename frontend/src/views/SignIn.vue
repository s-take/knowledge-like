<template>
  <v-form>
    <h2 class="headline font-weight-bold mb-3 text-center">SignIn</h2>
    <div>{{ message }}</div>
    <v-text-field prepend-icon="person" v-model="username" label="ユーザ名"></v-text-field>
    <v-text-field prepend-icon="lock" v-model="password" label="パスワード" type="password"></v-text-field>
    <v-card-actions>
      <v-btn primary large block @click="signIn">SignIn</v-btn>
    </v-card-actions>
    <router-link to="/signup">
      <h2 class="headline font-weight-bold mb-3 text-center">未登録の方はこちら</h2>
    </router-link>
  </v-form>
</template>

<script>
import axios from "axios";
import store from "@/store";

export default {
  name: "Signin",
  data() {
    return {
      message: "",
      username: "",
      password: ""
    };
  },
  methods: {
    signIn: function() {
      axios
        .post("http://localhost:8080/signin", {
          name: this.username,
          password: this.password
        })
        .then(
          response => {
            localStorage.setItem("token", response.data.token);
            store.commit("auth/setSignIn", true);
            this.$router.push("/");
          },
          err => {
            console.log(err);
            this.message =
              "ログインに失敗しました。ユーザ名またはパスワードが違います";
          }
        );
    }
  }
};
</script>

<style>
</style>
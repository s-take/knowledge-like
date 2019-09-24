<template>
  <v-form>
    <h2 class="headline font-weight-bold mb-3 text-center">SignUp</h2>
    <div>{{ message }}</div>
    <v-text-field prepend-icon="person" v-model="username" label="ユーザ名"></v-text-field>
    <v-text-field prepend-icon="lock" v-model="password" label="パスワード" type="password"></v-text-field>
    <v-card-actions>
      <v-btn primary large block @click="signUp">SignUp</v-btn>
    </v-card-actions>
  </v-form>
</template>

<script>
import axios from "axios";

export default {
  name: "Signup",
  data() {
    return {
      message: "",
      username: "",
      password: ""
    };
  },
  methods: {
    signUp: function() {
      axios
        .post("http://localhost:8080/signup", {
          name: this.username,
          password: this.password
        })
        .then(
          response => {
            this.message = "ユーザを作成しました";
            this.$router.push({ path: "/signin" });
            console.log(response);
          },
          err => {
            console.log(err);
            this.message =
              "サインアップに失敗しました。既に登録済みかユーザ名またはパスワードが空です";
          }
        );
    }
  }
};
</script>

<style>
</style>
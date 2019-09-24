<template>
  <v-row>
    <v-col cols="12">
      <v-data-table
        :headers="headers"
        :items="knowledges"
        item-key="id"
        :sort-by="['count']"
        :sort-desc="[true]"
        class="elevation-1"
        :loading="isLoading"
      >
        <template v-slot:item.title="props">
          <router-link
            v-bind:to="{ name : 'voteresultbyid', params : { id: props.item.id }}"
          >{{ props.item.title }}</router-link>
        </template>
      </v-data-table>
    </v-col>
  </v-row>
</template>

<script>
import { mapGetters } from "vuex";
import { RepositoryFactory } from "./../Repositories/RepositoryFactory";
const KnowledgeRepository = RepositoryFactory.get("knowledge");
const LikeRepository = RepositoryFactory.get("like");

export default {
  data() {
    return {
      // table param
      headers: [
        { text: "タイトル", value: "title" },
        { text: "作成者", value: "author" },
        { text: "獲得数", value: "count" }
      ],
      // other
      isLoading: false,
      likes: [],
      knowledges: []
    };
  },
  created() {
    this.fetch();
    this.likefetch();
  },
  methods: {
    async fetch() {
      this.isLoading = true;
      const { data } = await KnowledgeRepository.get();
      for (let i in data) {
        data[i].count = 0;
      }
      this.isLoading = false;
      this.knowledges = data;
    },
    async likefetch() {
      const { data } = await LikeRepository.get();
      this.likes = data;
      //console.log(this.likes);
      for (let i in this.knowledges) {
        this.likes.forEach(like => {
          if (this.knowledges[i].id === like.knowledgeId) {
            this.$set(
              this.knowledges[i],
              "count",
              this.knowledges[i].count + 1
            );
          }
        });
      }
    }
  }
};
</script>

<style>
</style>
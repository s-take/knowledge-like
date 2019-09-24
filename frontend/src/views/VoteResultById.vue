<template>
  <v-row>
    <v-col cols="12">
      <v-data-table
        :headers="headers"
        :items="likes"
        item-key="id"
        class="elevation-1"
        :loading="isLoading"
      >
        <template v-slot:top></template>
      </v-data-table>
    </v-col>
  </v-row>
</template>

<script>
import { mapGetters } from "vuex";
import { RepositoryFactory } from "./../Repositories/RepositoryFactory";
const LikeRepository = RepositoryFactory.get("like");

export default {
  data() {
    return {
      // table param
      headers: [
        { text: "理由", value: "reason" },
        { text: "コメント", value: "comment" }
      ],
      // other
      isLoading: false,
      likes: []
    };
  },
  created() {
    this.likefetch();
  },
  methods: {
    async likefetch() {
      const { data } = await LikeRepository.getByKnowledgeID(
        this.$route.params.id
      );
      this.likes = data;
    }
  }
};
</script>

<style>
</style>
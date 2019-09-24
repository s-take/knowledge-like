<template>
  <v-row>
    <v-col cols="12">
      <v-data-table
        v-model="selected"
        :headers="headers"
        :items="knowledges"
        :single-select="singleSelect"
        item-key="id"
        show-select
        class="elevation-1"
        dense
        :loading="isLoading"
        :hide-default-footer="true"
      >
        <template v-slot:item.reason="props">
          <v-edit-dialog
            :return-value.sync="props.item.reason"
            large
            persist
            @save="save"
            @cancel="cancel"
            @open="open"
            @close="close"
          >
            {{ props.item.reason }}
            <template v-slot:input>
              <v-radio-group v-model="props.item.reason">
                <v-radio label="技術的" value="技術的"></v-radio>
                <v-radio label="面白い" value="面白い"></v-radio>
                <v-radio label="役に立った" value="役に立った"></v-radio>
                <v-radio label="その他" value="その他"></v-radio>
              </v-radio-group>
            </template>
          </v-edit-dialog>
        </template>
        <template v-slot:item.comment="props">
          <v-edit-dialog
            :return-value.sync="props.item.comment"
            @save="save"
            @cancel="cancel"
            @open="open"
            @close="close"
          >
            {{ props.item.comment }}
            <template v-slot:input>
              <v-text-field v-model="props.item.comment" label="Edit" single-line autofocus></v-text-field>
            </template>
          </v-edit-dialog>
        </template>
      </v-data-table>
    </v-col>
    <v-col class="text-center" cols="12">
      <div>{{ message }}</div>
    </v-col>
    <v-col class="text-center" cols="12">
      <v-btn large color="primary" @click="postLike">投票</v-btn>
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
      singleSelect: false,
      snack: false,
      snackColor: "",
      snackText: "",
      selected: [],
      headers: [
        { text: "タイトル", value: "title" },
        { text: "作成者", value: "author" },
        { text: "理由", value: "reason" },
        { text: "コメント(あれば)", value: "comment" }
      ],
      // other
      isLoading: false,
      knowledges: [],
      likes: [],
      message: "",
      predelete: true
    };
  },
  computed: {
    ...mapGetters({
      isSignIn: "auth/isSignIn"
    })
  },
  created() {
    this.fetch();
    this.likefetch();
  },
  methods: {
    async fetch() {
      this.isLoading = true;
      const { data } = await KnowledgeRepository.get();
      this.isLoading = false;
      this.knowledges = data;
    },
    async likefetch() {
      const { data } = await LikeRepository.getByUserID();
      if (data.length === 0) {
        this.predelete = false;
        return;
      }
      this.likes = data;
      for (let i in this.knowledges) {
        for (let like of this.likes) {
          if (this.knowledges[i].id === like.knowledgeId) {
            this.knowledges[i].reason = like.reason;
            this.knowledges[i].comment = like.comment;
            like.id = like.knowledgeId;
            this.selected.push(like);
          }
        }
      }
      // console.log(this.selected);
    },
    async postLike() {
      //console.log(this.selected);
      const likes = this.selected;
      if (likes.length != 3) {
        this.message = "３つ選んでください";
        return;
      }

      if (this.predelete) {
        await LikeRepository.deleteLike();
      }

      for (let like of likes) {
        if (like.reason === undefined) {
          like.reason = "";
        }
        if (like.comment === undefined) {
          like.comment = "";
        }
        let body = {
          knowledgeId: like.id,
          reason: like.reason,
          comment: like.comment
        };
        await LikeRepository.createLike(body);
      }

      this.message =
        "投票ありがとうございました。期日までは何回でもやり直しできます。";
      // console.log(likes);
    },
    // edit table
    save() {
      this.snack = true;
      this.snackColor = "success";
      this.snackText = "Data saved";
    },
    cancel() {
      this.snack = true;
      this.snackColor = "error";
      this.snackText = "Canceled";
    },
    open() {
      this.snack = true;
      this.snackColor = "info";
      this.snackText = "Dialog opened";
    },
    close() {
      //console.log("Dialog closed");
    }
  }
};
</script>

<style>
</style>
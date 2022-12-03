<template>
  <div class="view-article-container">
    <div class="article">
      <ArticleViewCardVue :article="article"></ArticleViewCardVue>
    </div>
    <!-- <div class="input-comment">
      <CommentVue></CommentVue>
    </div> -->
    <div class="comment-list">
      <CommentListVue :comments="comments"></CommentListVue>
    </div>
  </div>
</template>

<script lang="ts" setup>
import ArticleViewCardVue from "@/components/ArticleViewCard.vue";
import CommentListVue from "@/components/CommentList.vue";
import { ArticleResponse, CommentResponse } from "@/model/response";
import { onMounted, ref } from "vue";
import { getArticle } from "@/api/api_article";
import { getComment } from "@/api/api_comments";
import { useRoute } from "vue-router";

// const props = defineProps<{ article: ArticleResponse }>();
const route = useRoute();
const article = ref<ArticleResponse>();
const comments = ref<Array<CommentResponse>>();

const obtainArticleDetail = async () => {
  const { data } = await getArticle(route.params.id as string);
  if (data.code == 0) {
    console.log("获取文章成功");
    console.log(data.data);
    article.value = data.data;
  }
};

const obtainArticleComments = async () => {
  const { data } = await getComment({
    entityType: "article",
    entityId: +route.params.id,
  });
  if (data.code == 0) {
    console.log("获取文章评论成功");
    console.log(data.data);
    comments.value = data.data;
  }
};

onMounted(async () => {
  await obtainArticleDetail();
  await obtainArticleComments();
});
</script>

<style lang="scss" scoped>
.view-article-container {
  position: absolute;
  top: 65px;
  left: 15%;
  .input-comment {
    margin-top: 0px;
  }
  .comment-list {
    margin-top: 0px;
  }
}
</style>

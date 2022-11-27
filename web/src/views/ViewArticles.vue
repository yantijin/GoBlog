<template>
  <div class="view-article-container">
    <div class="article">
      <ArticleViewCardVue></ArticleViewCardVue>
    </div>
    <!-- <div class="input-comment">
      <CommentVue></CommentVue>
    </div> -->
    <div class="comment-list">
      <CommentListVue></CommentListVue>
    </div>
  </div>
</template>

<script lang="ts" setup>
import ArticleViewCardVue from "@/components/ArticleViewCard.vue";
import CommentVue from "@/components/Comment.vue";
import CommentListVue from "@/components/CommentList.vue";
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

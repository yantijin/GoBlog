<template>
  <div class="article-preview">
    <!-- <el-button @click="clickBtn">点击按钮</el-button>
    <el-button @click="router.push('/article/create')">创建文章</el-button> -->
    <ArticlePreviewCardVue></ArticlePreviewCardVue>
  </div>
</template>

<script lang="ts" setup>
import ArticlePreviewCardVue from "@/components/ArticlePreviewCard.vue";
import { useRouter } from "vue-router";
const clickBtn = () => {
  console.log(import.meta.env);
};
const router = useRouter();
</script>

<style lang="scss" scoped>
.article-preview {
  position: absolute;
  top: 70px;
  left: 15%;
  width: 70%;
}
</style>

<template>
  <el-card class="article-view-card">
    <template #header>
      <h2 id="title">{{ article.Title }}</h2>
      <div class="article-view-meta">
        <span
          >由
          <router-link
            :to="'/user/' + article.UserInfo.ID"
            class="article-author"
            >{{ article.UserInfo.NickName }}</router-link
          >
        </span>
        发布于
        <time>{{ article.CreateTime }}</time>
      </div>
    </template>
    <div v-html="article.Content" class="content"></div>
  </el-card>
</template>

<script lang="ts" setup>
// const article = async() => {
//   await
// }

import { ArticleResponse } from "@/model/response";
import { RouterLink } from "vue-router";

const article: ArticleResponse = {
  ArticleId: 1,
  Title: "测试",
  Content: "<a href='https://whiteyan.top'>小白的个人博客</a>",
  ViewCount: 1,
  CreateTime: 100,
  UserInfo: { ID: 1, NickName: "小白", Avatar: "", CreateTime: 0 },
  LikeCount: 4,
  CommentCount: 10,
  CommnetTime: 30,
};
</script>

<style>
.article-view-card {
  width: 70vw;
  margin-bottom: 0px;
}

#title {
  text-align: left;
}

.article-view-meta {
  display: flex;
  margin-top: 0px;
  margin-bottom: 0px;
  padding: 0 6px 0 0;
  align-items: center;
  font-size: 13px;
}

.article-author {
  font-weight: bold;
  padding: 0 3px;
  color: black;
}

.content {
  text-align: left;
  font-size: 16px;
}
</style>

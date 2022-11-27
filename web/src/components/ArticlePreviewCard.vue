<!-- 首页预览文章的card -->
<template>
  <div
    v-for="article in props.articles"
    :key="article.ArticleId"
    class="article-item"
  >
    <el-card class="article-preview-container">
      <div class="article-item-main">
        <div class="article-info">
          <router-link
            :to="'/article/' + article.ArticleId"
            class="article-title"
            >{{ article.Title }}</router-link
          >
        </div>
      </div>

      <div class="article-meta">
        <div class="article-meta-left">
          <div class="article-meta-item">
            <router-link
              :to="'/user' + article.UserInfo.ID"
              class="article-author"
            >
              <span>
                {{ article.UserInfo.NickName }}
              </span>
            </router-link>
            <time>发布于 {{ article.CreateTime }}</time>
          </div>
        </div>

        <div class="article-meta-right">
          <span
            ><el-icon><View /></el-icon></span
          >&nbsp;{{ article.ViewCount }}
          <span
            ><el-icon><ChatDotRound /></el-icon></span
          >&nbsp;{{ article.CommentCount }}
          <span
            ><el-icon><CaretTop /></el-icon></span
          >&nbsp;{{ article.LikeCount }}
        </div>
      </div>
    </el-card>
  </div>
</template>

<script lang="ts" setup>
import { ArticleResponse } from "@/model/response";
import { computed } from "vue";
import { RouterLink } from "vue-router";
import { View, ChatDotRound, CaretTop } from "@element-plus/icons-vue";

interface articlesProps {
  articles: any[];
}
// const props = defineProps<{
//   articles: ArticleResponse[];
// }>();
const props = withDefaults(defineProps<articlesProps>(), {
  // articles: () => [] as any[],
  articles: () => {
    return [
      {
        ArticleId: 1,
        Title: "这是测试的第一个文章标题",
        Content: "测试1",
        ViewCount: 1,
        CreateTime: 100,
        UserInfo: { ID: 1, NickName: "小白", Avatar: "", CreateTime: 0 },
        LikeCount: 4,
        CommentCount: 10,
        CommnetTime: 30,
      },
      {
        ArticleId: 2,
        Title: "这是测试的第二个文章标题",
        Content: "测试2",
        ViewCount: 1,
        CreateTime: 100,
        UserInfo: { ID: 1, NickName: "小白", Avatar: "", CreateTime: 0 },
        LikeCount: 4,
        CommentCount: 10,
        CommnetTime: 30,
      },
    ];
  },
});

const format = computed(val => {
  return (val: any) => {
    return;
  };
});
</script>

<style>
.article-preview-container {
  /* width: 70vw; */
  padding: 0px;
  margin-top: 2px;
}

.article-item {
  padding: 0px 0px;
  border-radius: 3px;
  line-height: 24px;
  margin-bottom: 0px;
}

.article-item-main {
  width: 100%;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.article-title {
  font-size: 18px;
  line-height: 30px;
  font-weight: bold;
  color: black;
  overflow: hidden;
  text-overflow: ellipsis;
  float: left;
}

.article-meta {
  display: flex;
  align-items: center;
  font-size: 13px;
  padding-top: 10px;
}

.article-meta-item {
  padding: 0 6px 0 0;
}

.article-author {
  font-weight: bold;
  padding: 0 6px;
  color: black;
}

.article-meta-right {
  margin-left: 20px;
}

.el-card__body {
  padding: 10px 10px 5px 10px;
}
</style>

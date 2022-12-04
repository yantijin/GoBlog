<!-- 首页预览文章的card -->
<template>
  <div
    v-for="article in props.articles"
    :key="article.articleId"
    class="article-item"
  >
    <el-card class="article-preview-container">
      <div class="article-item-main">
        <div class="article-info">
          <router-link
            :to="'/article/' + article.articleId"
            class="article-title"
            >{{ article.title }}</router-link
          >
        </div>
      </div>

      <div class="article-meta">
        <div class="article-meta-left">
          <div class="article-meta-item">
            <router-link
              :to="'/user/' + article.user.id"
              class="article-author"
            >
              <span>
                {{ article.user.nickname }}
              </span>
            </router-link>
            <time>发布于 {{ unix2Date(article.createTime) }}</time>
          </div>
        </div>

        <div class="article-meta-right">
          <span
            ><el-icon><View /></el-icon></span
          >&nbsp;{{ article.viewCount }}
          <span
            ><el-icon><ChatDotRound /></el-icon></span
          >&nbsp;{{ article.commentCount }}
          <span
            ><el-icon><CaretTop /></el-icon></span
          >&nbsp;{{ article.likeCount }}
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
import unix2Date from "@/utils/date";

// interface articlesProps {
//   articles: any[];
// }
const props = defineProps<{
  articles: ArticleResponse[];
}>();
// const props = withDefaults(
//   defineProps<{ articles: Array<ArticleResponse> }>(),
//   {
//     // articles: () => [] as any[],
//     articles: () => {
//       return [
//         {
//           articleId: 1,
//           title: "这是测试的第一个文章标题",
//           content: "测试1",
//           viewCount: 1,
//           createTime: 100,
//           user: {
//             id: 1,
//             nickname: "小白",
//             avatar: "",
//             email: "yantijn@163.com",
//             username: "yantijin",
//             uuid: "",
//           },
//           likeCount: 4,
//           commentCount: 10,
//           commentTime: 30,
//         },
//         {
//           articleId: 2,
//           title: "这是测试的第二个文章标题",
//           content: "测试2",
//           viewCount: 1,
//           createTime: 100,
//           user: {
//             id: 1,
//             nickname: "小白",
//             avatar: "",
//             email: "yantijn@163.com",
//             username: "yantijin",
//             uuid: "",
//           },
//           likeCount: 4,
//           commentCount: 10,
//           commentTime: 30,
//         },
//       ];
//     },
//   }
// );

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

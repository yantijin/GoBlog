<template>
  <el-card class="article-view-card">
    <template #header>
      <h2 id="title">{{ article.title }}</h2>
      <div class="article-view-meta">
        <span
          >由
          <router-link
            :to="'/user/' + article.user.id"
            class="article-author"
            >{{ article.user.nickname }}</router-link
          >
        </span>
        发布于&nbsp;
        <time>{{ unix2Date(article.createTime) }}</time>
      </div>
    </template>
    <ViewContentVue
      :md="article.content"
      mode="light"
      st="github"
    ></ViewContentVue>
    <hr />
    <div class="end-actions">
      <el-button text @click="handleClick">
        <el-icon><Edit /></el-icon>
        <span>回复</span>
      </el-button>
      <el-button text>
        <el-icon> <CaretTop /></el-icon>
        <span> 点赞 </span>
      </el-button>
    </div>
  </el-card>
  <div v-if="replyMain">
    <Comment
      placeholder=""
      entity-type="article"
      :entity-id="article.articleId"
    ></Comment>
  </div>
</template>

<script lang="ts" setup>
import { ArticleResponse } from "@/model/response";
import { RouterLink } from "vue-router";
import { Edit, CaretTop } from "@element-plus/icons-vue";
import Comment from "@/components/Comment.vue";
import { ref } from "vue";
import unix2Date from "@/utils/date";
import ViewContentVue from "./ViewContent.vue";

const replyMain = ref(false);

const handleClick = () => {
  if (replyMain.value === false) {
    replyMain.value = true;
  } else {
    replyMain.value = false;
  }
};

const props = withDefaults(defineProps<{ article: ArticleResponse }>(), {
  article: () => {
    return {
      articleId: 1,
      title: "测试",
      content: "",
      viewCount: 1,
      createTime: 100,
      user: {
        id: 1,
        nickname: "小白",
        avatar: "",
        email: "yantijin@163.com",
        username: "yantijin",
        uuid: "",
      },
      likeCount: 4,
      commentCount: 10,
      commentTime: 30,
    };
  },
});
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

.end-actions {
  display: flex;
  margin-bottom: 0px;
}
</style>

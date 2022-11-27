<!-- 主要是对文章下方的comment的list进行展示 -->
<template>
  <div
    v-for="comment in props.comments"
    :key="comment.CommentId"
    class="comment"
  >
    <!-- <div class=""> -->
    <!-- 头像， NickName, 发布时间在右侧, 评论内容， 下方： 点赞数目，评论数目 -->
    <el-card class="comment-container">
      <div class="comment-title">
        <div class="comment-item-left">
          <el-avatar
            :src="comment.UserInfo.Avatar"
            size="40"
            class="avatar"
          ></el-avatar>
        </div>
        <div class="comment-item-avatar">
          <router-link :to="'/user/' + comment.UserInfo.ID" class="user-info">
            {{ comment.UserInfo.NickName }}</router-link
          >
        </div>
        <div class="comment-time">
          <time>{{ comment.CreateTime }}</time>
        </div>
      </div>

      <div class="comment-content">
        <div v-html="comment.Content"></div>
      </div>
      <div class="comment-actions">
        <span>
          <el-icon><CaretTop /></el-icon> &nbsp;{{ comment.LikeCount }}
        </span>
        &nbsp;
        <span
          ><el-icon @click="switchReply(comment)"
            ><ChatDotRound /></el-icon></span
        >&nbsp;{{ comment.CommentCount }}
      </div>
      <div v-if="reply.commentId == comment.CommentId">
        <Comment placeholder=""></Comment>
      </div>
    </el-card>
    <!-- </div> -->
  </div>
</template>

<script lang="ts" setup>
import { CommentResponse } from "@/model/response";
import { RouterLink } from "vue-router";
import { CaretTop, ChatDotRound } from "@element-plus/icons-vue";
import Comment from "@/components/Comment.vue";
import { ref } from "vue";

const reply = ref({
  commentId: 0,
  content: "",
});
// const a: CommentResponse =
const switchReply = (comment: CommentResponse) => {
  if (reply.value.commentId == comment.CommentId) {
    (reply.value.commentId = 0), (reply.value.content = "");
  } else {
    reply.value.commentId = comment.CommentId;
  }
};

const props = withDefaults(
  defineProps<{
    comments: CommentResponse[];
  }>(),
  {
    comments: () => [
      {
        CommentId: 5,
        Content: "可以参考<a href='https://whiteyan.top'>小白的个人博客</a>",
        EntityType: "article",
        EntityId: 3,
        LikeCount: 66,
        CommentCount: 4,
        CreateTime: 60,
        Liked: false,
        UserInfo: {
          ID: 1,
          NickName: "小白",
          Avatar:
            "https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png",
          CreateTime: 2,
        },
      },
      {
        CommentId: 4,
        Content: "可以参考<a href='https://whiteyan.top'>小白的个人博客</a>",
        EntityType: "article",
        EntityId: 3,
        LikeCount: 66,
        CommentCount: 4,
        CreateTime: 60,
        Liked: false,
        UserInfo: {
          ID: 1,
          NickName: "小白",
          Avatar:
            "https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png",
          CreateTime: 2,
        },
      },
    ],
  }
);
</script>

<style>
.comment-container {
  padding: 5px 0px;
  border-radius: 3px;
  margin-bottom: 5px;
  margin-top: 5px;
  /* background-color: white; */
}

.comment-title {
  text-align: left;
  display: flex;
  margin-left: 6px;
}

.comment-item-left {
  text-align: left;
  /* margin-left: 6px; */
  /* display: inline-block; */
}

.comment-item-avatar {
  display: flex;
  justify-content: space-between;
  margin-left: 6px;
  /* margin-right: 0; */
  /* margin: auto; */
  text-align: left;
}

.user-info {
  color: black;
  font-weight: bold;
  font-weight: 600;
  font-size: 16px;
}

.comment-time {
  font-size: 13px;
  margin-right: 5px;
  margin-left: auto;
}

.comment-content {
  /* display: flex; */
  margin-left: 45px;
  text-align: left;
  /* vertical-align: ; */
}

.comment-actions {
  /* display: flex; */
  text-align: left;
  margin-left: 45px;
  margin-top: 8px;
  margin-bottom: 0px;
}
</style>

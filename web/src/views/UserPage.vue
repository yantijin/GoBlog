<template>
  <div class="userpage-container">
    <div class="top-bar-container">
      <UserProfileTopBarVue :local-user="userinfo"></UserProfileTopBarVue>
    </div>
    <div class="main-item">
      <div class="left-side-bar">
        <SideBarUserProfileVue :userinfo="userinfo"></SideBarUserProfileVue>
      </div>
      <div class="right-side-container">
        <el-menu
          default-active="article"
          class="el-menu"
          mode="horizontal"
          @select="handleSelect"
        >
          <el-menu-item index="article">文章</el-menu-item>
          <el-menu-item index="comment">评论</el-menu-item>
        </el-menu>
        <div v-if="isArticle === true">
          <ArticlePreviewCardVue
            :articles="articleInfo"
          ></ArticlePreviewCardVue>
        </div>
        <div v-if="isArticle === false">
          <CommentListVue :comments="commentInfo"></CommentListVue>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import UserProfileTopBarVue from "@/components/UserProfileTopBar.vue";
import SideBarUserProfileVue from "@/components/SideBarUserProfile.vue";
import ArticlePreviewCardVue from "@/components/ArticlePreviewCard.vue";
import CommentListVue from "@/components/CommentList.vue";
import { onMounted, ref } from "vue";
import { useRoute } from "vue-router";
import { useUserStore } from "@/pinia/modules/user";
import {
  UserInfoData,
  ArticleResponse,
  CommentResponse,
} from "@/model/response";
import { getUserInfo } from "@/api/api_user";
import { getUserArticles } from "@/api/api_article";
import { getUserComments } from "@/api/api_comments";

// interface UserPageStruct {
//   userId: number,
// }

// const props = defineProps({
//   userId:
// })
const userStore = useUserStore();
const route = useRoute();
const isArticle = ref(true);
const userinfo = ref<UserInfoData>();
const articleInfo = ref<Array<ArticleResponse>>();
const commentInfo = ref<Array<CommentResponse>>();

const obtainUserInfo = async () => {
  if (route.params.id != userStore.userInfo.id) {
    console.log("当前用户浏览个人信息");
    return userStore.userInfo;
  } else {
    const { data } = await getUserInfo(route.params.id as string);
    if (data.code == 0) {
      // console.log("user信息");
      // console.log(data.data);
      return data.data;
    }
  }
};

const obtainUserArticles = async () => {
  const { data } = await getUserArticles(route.params.id as string);
  if (data.code == 0) {
    console.log("article信息");
    console.log(data.data);
    return data.data;
  }
};

const obtainUserComments = async () => {
  const { data } = await getUserComments(route.params.id as string);
  if (data.code == 0) {
    console.log("comments 信息");
    console.log(data.data);
    return data.data;
  }
};

onMounted(async () => {
  const uInfo = await obtainUserInfo();
  const aInfo = await obtainUserArticles();
  const cInfo = await obtainUserComments();
  userinfo.value = uInfo as UserInfoData;
  articleInfo.value = aInfo as Array<ArticleResponse>;
  commentInfo.value = cInfo as Array<CommentResponse>;
  console.log("mounted结果");
  // console.log(userinfo.value);
  console.log(articleInfo.value);
});

const handleSelect = (
  index: string,
  indexPath: string,
  routeResult: string
) => {
  console.log(index, indexPath, routeResult);
  if (index === "article") {
    isArticle.value = true;
  } else {
    isArticle.value = false;
  }
};
</script>

<style lang="scss" scoped>
$sidebar-width: 260px;
$sidebar-margin: 10px; // sidebar之间的缝隙
$content-width: calc(100% - #{$sidebar-width} - #{$sidebar-margin});
.userpage-container {
  position: absolute;
  margin: auto;
  top: 65px;
  left: 15%;
  // top: 65px;
}
.top-bar-container {
  // position: fixed;
  // top: 65px;
  // left: 15%;
  display: flex;
  // margin-top: 35px;
}
.main-item {
  display: flex;
  .left-side-bar {
    width: $sidebar-width;
  }
  .right-side-container {
    width: $content-width;
    margin-left: 10px;
  }
}
</style>

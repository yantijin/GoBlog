<template>
  <div class="userpage-container">
    <div class="top-bar-container">
      <UserProfileTopBarVue></UserProfileTopBarVue>
    </div>
    <div class="main-item">
      <div class="left-side-bar">
        <SideBarUserProfileVue></SideBarUserProfileVue>
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
          <ArticlePreviewCardVue></ArticlePreviewCardVue>
        </div>
        <div v-if="isArticle === false">
          <CommentListVue></CommentListVue>
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
import { ref } from "vue";

// interface UserPageStruct {
//   userId: number,
// }

// const props = defineProps({
//   userId:
// })

const isArticle = ref(true);

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

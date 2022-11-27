import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    redirect: "/home",
  },
  {
    path: "/home",
    name: "home",
    component: () => import("@/views/Home.vue"),
  },
  {
    path: "/article",
    name: "article",
    children: [
      {
        path: "create",
        // component: () => import("@/components/CreateArticle.vue"),
        // component: () => import("@/components/ArticlePreviewCard.vue"),
        // component: () => import("@/components/ArticleViewCard.vue"),
        // component: () => import("@/components/Comment.vue"),
        // component: () => import("@/components/CommentList.vue"),
        // component: () => import("@/components/EditUserInfoCard.vue"),
        // component: () => import("@/components/SideBarUserProfile.vue"),
        component: () => import("@/views/EditArticle.vue"),
      },
      {
        path: ":id",
        component: () => import("@/views/ViewArticles.vue"),
      },
    ],
  },
  {
    path: "/user/",
    children: [
      {
        path: ":id",
        name: "user",
        component: () => import("@/views/UserPage.vue"),
      },
      {
        path: "edit",
        name: "edit",
        component: () => import("@/views/EditUserInfo.vue"),
      },
    ],
  },
];

// 3
const router = createRouter({
  history: createWebHashHistory(),
  routes: routes,
});
// 4
export default router;

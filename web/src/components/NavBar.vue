<!-- 主要用于顶部导航栏的创建 -->
<template>
  <div class="nav">
    <el-menu
      :default-active="activeIndex"
      mode="horizontal"
      background-color="white"
      :ellipsis="false"
      @select="handleSelect"
    >
      <div id="logo">
        <el-menu-item index="home">
          <!-- <img src="../assets/logo.jpg" /> -->
          <el-avatar :src="logoImage" :size="55" shape="square"></el-avatar>
        </el-menu-item>
      </div>
      <div class="flex-grow" />

      <el-sub-menu index="publish">
        <template #title>发表</template>
        <el-menu-item index="publishArticle">发文章</el-menu-item>
      </el-sub-menu>
      <el-dropdown v-if="userStore.userInfo != ''" id="dp">
        <el-avatar shape="square" :src="squareUrl" />
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item @click="router.push('/user/1')">
              个人中心</el-dropdown-item
            >
          </el-dropdown-menu>
        </template>
      </el-dropdown>

      <div id="second">
        <el-menu-item index="login">
          <el-button
            @click="loginBtn"
            type="primary"
            v-if="userStore.userInfo == ''"
            >登录</el-button
          >
          <RegisterAndLogInVue
            :visible="states.visible"
            :flag="states.flag"
            @close="handleClose"
          ></RegisterAndLogInVue>
        </el-menu-item>
      </div>
    </el-menu>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, watch, computed } from "vue";
import RegisterAndLogInVue from "@/components/RegisterAndLogIn.vue";
import { useUserStore } from "@/pinia/modules/user";
import logoImg from "@/assets/logo.jpg";
import { useRouter } from "vue-router";

const activeIndex = ref("origin");
const logoImage = ref(logoImg);
const router = useRouter();

const states = reactive({
  visible: false,
  flag: false,
});

const handleSelect = (
  index: string,
  indexPath: string,
  routeResult: string
) => {
  console.log(index, indexPath, routeResult);
  if (index === "home") {
    router.push("/home");
  } else if (index === "publishArticle") {
    router.push("/article/create");
  }
};

const loginBtn = () => {
  states.visible = true;
};

const handleClose = (val: boolean) => {
  states.visible = val;
};

const userStore = useUserStore();
// const squareUrl = ref(
//   "https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png"
// );

// watch(
//   () => userStore.userInfo,
//   (val, oldval) => {
//     squareUrl.value = userStore.userInfo.avatar;
//   }
// );

const squareUrl = computed(() => userStore.userInfo.avatar);
</script>

<style>
.flex-grow {
  flex-grow: 1;
}
.nav {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
}
#logo {
  margin-left: 15%;
}
#second {
  margin-right: 15%;
  top: 15px;
}
#dp {
  margin-top: 10px;
}
</style>

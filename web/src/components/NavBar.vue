<!-- 主要用于顶部导航栏的创建 -->
<template>
  <div class="nav">
    <el-menu
      :default-active="activeIndex"
      mode="horizontal"
      background-color="grey"
      :ellipsis="false"
    >
      <div id="logo"><el-menu-item index="origin">LOGO</el-menu-item></div>
      <div class="flex-grow" />

      <el-sub-menu index="publish">
        <template #title>发表</template>
        <el-menu-item index="publishArticle">发文章</el-menu-item>
      </el-sub-menu>
      <el-dropdown v-if="userStore.isLogIn" id="dp">
        <el-avatar shape="square" :src="squareUrl" />
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item> 个人中心</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>

      <div id="second">
        <el-menu-item index="login">
          <el-button @click="loginBtn" type="primary" v-if="!userStore.isLogIn"
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
import { ref, reactive, watch } from "vue";
import RegisterAndLogInVue from "@/components/RegisterAndLogIn.vue";
import { useUserStore } from "@/pinia/modules/user";

const activeIndex = ref("origin");

const states = reactive({
  visible: false,
  flag: false,
});

const loginBtn = () => {
  states.visible = true;
};

const handleClose = (val: boolean) => {
  states.visible = val;
};

const userStore = useUserStore();
const squareUrl = ref(
  "https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png"
);

watch(
  () => userStore.userInfo,
  (val, oldval) => {
    squareUrl.value = userStore.userInfo.avatar;
  }
);
</script>

<style>
.flex-grow {
  flex-grow: 1;
}
.nav {
  position: fixed;
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

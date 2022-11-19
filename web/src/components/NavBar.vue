<!-- 主要用于顶部导航栏的创建 -->
<template>
  <div class="nav">
    <el-menu
      :default-active="activeIndex"
      mode="horizontal"
      background-color="grey"
      :ellipsis="false"
      @select="handleSelect"
    >
      <div id="logo"><el-menu-item index="origin">LOGO</el-menu-item></div>
      <div class="flex-grow" />

      <el-sub-menu index="publish">
        <template #title>发表</template>
        <el-menu-item index="publishArticle">发文章</el-menu-item>
      </el-sub-menu>
      <div id="second">
        <el-menu-item index="login">
          <el-button @click="loginBtn" type="primary">登录</el-button>
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
import { ref, reactive } from "vue";
import RegisterAndLogInVue from "@/components/RegisterAndLogIn.vue";

const activeIndex = ref("origin");
const handleSelect = (index: string, indexPath: string, item: string) => {
  console.log(index, indexPath, item);
};

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
</style>

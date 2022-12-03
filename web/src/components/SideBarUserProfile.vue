<template>
  <el-card class="card-container">
    <template #header>
      <div class="card-header">
        <span>个人资料</span>
        <el-button
          class="button"
          text
          @click="router.push('/user/edit')"
          v-if="route.params.id == userStore.userInfo.id"
          >编辑资料</el-button
        >
      </div>
    </template>
    <el-form label-width="60px">
      <el-form-item label="用户名">
        <span> {{ userinfo.username }}</span>
      </el-form-item>
      <hr />
      <el-form-item label="昵称">
        <span> {{ userinfo.nickname }}</span>
      </el-form-item>
      <hr />
      <el-form-item label="邮箱">
        <span> {{ userinfo.email }}</span>
      </el-form-item>
    </el-form>
  </el-card>
</template>

<script lang="ts" setup>
import { UserInfoData } from "@/model/response";
import { useUserStore } from "@/pinia/modules/user";
import { useRouter, useRoute } from "vue-router";

const props = withDefaults(defineProps<{ userinfo: UserInfoData }>(), {
  userinfo: () => {
    return {
      nickname: "小白",
      id: 3,
      avatar:
        "https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png",
      email: "yantijin@163.com",
      username: "yantijin",
      uuid: "",
    };
  },
});

const router = useRouter();
const route = useRoute();
const userStore = useUserStore();
</script>

<style>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

hr {
  border: none;
  height: 1px;
  background-color: grey;
}

.el-form-item {
  margin-bottom: 0px;
}
</style>

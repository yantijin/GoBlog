<template>
  <el-card class="card-container">
    <template #header>
      <div class="card-header">
        <span>个人资料</span>
        <el-button class="button" text @click="router.push('/user/edit')"
          >编辑资料</el-button
        >
      </div>
    </template>
    <el-form label-width="60px">
      <el-form-item label="用户名">
        <span> {{ props.UserName }}</span>
      </el-form-item>
      <hr />
      <el-form-item label="昵称">
        <span> {{ props.NickName }}</span>
      </el-form-item>
      <hr />
      <el-form-item label="邮箱">
        <span> {{ props.Email }}</span>
      </el-form-item>
    </el-form>
  </el-card>
</template>

<script lang="ts" setup>
import { UserResponse } from "@/model/response";
import { useRouter } from "vue-router";

const props: UserResponse = {
  UserName: "yantijin",
  NickName: "小白",
  Email: "yantijin@163.com",
  Avatar: "https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png",
};

const router = useRouter();
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

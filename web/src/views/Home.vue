<template>
  <div>
    <nav-bar></nav-bar>
    <el-button @click="clickBtn">点击按钮</el-button>
  </div>
</template>

<script lang="ts" setup>
import NavBar from "@/components/NavBar.vue";
const clickBtn = () => {
  console.log(import.meta.env);
};
</script>

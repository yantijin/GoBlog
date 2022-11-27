<template>
  <div class="create-article">
    <CreateArticleVue></CreateArticleVue>
  </div>
</template>

<script lang="ts" setup>
import CreateArticleVue from "@/components/CreateArticle.vue";
</script>

<style lang="scss" scoped>
.create-article {
  position: absolute;
  top: 65px;
  left: 15%;
}
</style>

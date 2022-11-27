<template>
  <div class="edit-userinfo-container">
    <EditUserInfoCardVue></EditUserInfoCardVue>
  </div>
</template>

<script lang="ts" setup>
import EditUserInfoCardVue from "@/components/EditUserInfoCard.vue";
</script>

<style lang="scss" scoped>
.edit-userinfo-container {
  position: absolute;
  top: 65px;
  left: 15%;
}
</style>

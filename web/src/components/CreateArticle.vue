<template>
  <el-card class="article-card">
    <template #header>
      <h2 class="submit">发文章</h2>
    </template>

    <div>
      <el-input placeholder="请输入标题" v-model="title" />
    </div>
    <hr />
    <div>
      <MdEditorVue mode="ir" theme="classic" ref="MdEditorMethod" height="65vh">
      </MdEditorVue>
    </div>
    <!-- <el-button @click="getValue">获取值</el-button> -->
    <el-button type="primary" id="submitBtn">发表</el-button>
  </el-card>
</template>

<script lang="ts" setup>
import MdEditorVue from "./MdEditor.vue";
import { ref } from "vue";

const MdEditorMethod = ref<null | InstanceType<typeof MdEditorVue>>(null);

const title = ref("");
// const getValue = () => {
//   console.log(MdEditorMethod.value?.getValue());
// };
</script>

<style>
.submit {
  display: flex;
  margin-bottom: 0px;
}

.article-card {
  width: 70vw;
  height: 90vh;
}

/* .input {
  height: 60px;
} */

#submitBtn {
  margin-top: 15px;
  display: flex;
  height: 30px;
}
</style>

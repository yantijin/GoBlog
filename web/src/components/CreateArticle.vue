<template>
  <el-card class="article-card">
    <template #header>
      <span class="title">发文章</span>
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
    <el-button type="primary" id="submitBtn" @click="handleSubmit"
      >发表</el-button
    >
  </el-card>
</template>

<script lang="ts" setup>
import MdEditorVue from "./MdEditor.vue";
import { ref } from "vue";
import { publishArticle } from "@/api/api_article";
import { useRouter } from "vue-router";
import { ElMessage } from "element-plus";

const router = useRouter();
const MdEditorMethod = ref<null | InstanceType<typeof MdEditorVue>>(null);

const title = ref("");
// const getValue = () => {
//   console.log(MdEditorMethod.value?.getValue());
// };

const handleSubmit = async () => {
  const value = MdEditorMethod.value?.getValue();
  const { data } = await publishArticle({
    title: title.value,
    content: value as string,
    contentType: "markdown",
  });
  if (data.code == 0) {
    console.log("发表文章成功");
    console.log(data.data);
    ElMessage({
      type: "success",
      message: "发表成功，即将跳转",
    });
    router.push("/article/" + data.data.articleId);
  }
};
</script>

<style>
.title {
  display: flex;
  margin-bottom: 0px;
  font-size: 40px;
  font-weight: bold;
  margin-top: 20px;
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

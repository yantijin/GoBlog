<!-- 这里假设comment也能够输入富文本来编辑 -->
<template>
  <div class="comment-container">
    <div>
      <MdEditorVue
        placeholder="请输入评论内容"
        theme="classic"
        height="auto"
      ></MdEditorVue>
    </div>
    <div class="btn">
      <el-button type="primary" @click="submitComment">提交</el-button>
    </div>
  </div>
</template>

<script setup lang="ts">
import MdEditorVue from "@/components/MdEditor.vue";

interface commentProps {
  entityType: string;
  entityId: number;
}

const props = withDefaults(defineProps<commentProps>(), {
  entityType: "article",
  entityId: 0,
});

const submitComment = () => {
  console.log("hhh");
};
</script>

<style>
.btn {
  text-align: right;
  padding-top: 6px;
}

.comment-container {
  border: 2px solid grey;
}
</style>

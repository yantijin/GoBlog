<!-- 这里假设comment也能够输入富文本来编辑 -->
<template>
  <div class="comment-container">
    <div>
      <MdEditorVue
        placeholder="请输入评论内容"
        theme="classic"
        height="auto"
        ref="MdEditorMethod"
      ></MdEditorVue>
    </div>
    <div class="btn">
      <el-button type="primary" @click="submitComment">提交</el-button>
    </div>
  </div>
</template>

<script setup lang="ts">
import MdEditorVue from "@/components/MdEditor.vue";
import { createComment } from "@/api/api_comments";
import { ref } from "vue";
import { useRouter, useRoute } from "vue-router";

const router = useRouter();
const route = useRoute();
const MdEditorMethod = ref<null | InstanceType<typeof MdEditorVue>>(null);
interface commentProps {
  entityType: string;
  entityId: number;
}

const props = withDefaults(defineProps<commentProps>(), {
  entityType: "article",
  entityId: 0,
});

const submitComment = async () => {
  const { data } = await createComment({
    entityId: props.entityId,
    entityType: props.entityType,
    contentType: "markdown",
    content: MdEditorMethod.value?.getValue() as string,
  });
  if (data.code == 0) {
    console.log("提交评论成功");
    console.log(data.data);
    // router.replace("/article/" + route.params.id);
    router.go(0);
  }
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

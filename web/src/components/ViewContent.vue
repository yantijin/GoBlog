<!-- 主要利用vditor预览用户的文档 -->
<template>
  <div id="preview" class="md-preview"></div>
</template>

<script lang="ts" setup>
import Vditor from "vditor";
import "vditor/dist/index.css";
import { onMounted } from "vue";

type viewContentProp = {
  md: string; // 要渲染的md stirng
  st: "github" | "monokai"; // 代码样式，https://xyproto.github.io/splash/docs/longer/all.html?utm_source=ld246.com
  mode: "dark" | "light"; // 主题样式
};

const props = withDefaults(defineProps<viewContentProp>(), {
  md: "# 这是标题1\n ## 这是标题2",
  st: "github",
  mode: "light",
});

// 注意： 这里不能用onMounted，因为如果拉数据慢了，那么这里渲染的时候，props数据仍为空
// 等数据拉完之后，这里已经渲染完了；有以下两种解决方案
// watch: 当props数据拉取结束后，再进行渲染
// 将渲染定义为方法，并在父组件中拉取数据完成后，通过ref来调用子组件的方法
// watch(props, () => {
//   Vditor.preview(
//     document.getElementById("preview") as HTMLDivElement,
//     props.md,
//     {
//       hljs: { style: props.st },
//       mode: props.mode,
//     }
//   );
// });

onMounted(() => {
  Vditor.preview(
    document.getElementById("preview") as HTMLDivElement,
    props.md,
    {
      hljs: { style: props.st },
      mode: props.mode,
    }
  );
});
</script>

<style lang="scss" scoped>
.md-preview {
  text-align: left;
  // color: white;
}
</style>

<template>
  <div id="vditor" class="md" />
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import Vditor from "vditor";
import "vditor/dist/index.css";

// 可以参考 https://blog.csdn.net/lmy_loveF/article/details/125317648 进行配置

interface MdProps {
  value: string;
  height: number | string;
  width: string | number;
  mode: "ir" | "wysiwyg" | "sv" | undefined;
  isPin: boolean;
  isHideTools: boolean;
  theme: "dark" | "classic";
  placeholder: string;
}

const props = withDefaults(defineProps<MdProps>(), {
  value: "", // 默认填充的值
  height: 0.6 * window.innerHeight, // 编辑器的高度
  width: "auto", // 编辑器的宽度
  mode: "ir", // Vditor模式, wysiwyg/ir/sv
  isPin: false, // 是否固定工具栏
  isHideTools: false, // 是否隐藏工具栏
  theme: "dark", // 主题：dark, classic
  placeholder: "",
});

const vditor = ref<Vditor | null>(null);

onMounted(() => {
  vditor.value = new Vditor("vditor", {
    after: () => {
      // vditor.value is a instance of Vditor now and thus can be safely used here
      vditor.value!.setValue(props.value);
    },
    theme: props.theme,
    height: props.height,
    width: props.width,
    toolbarConfig: {
      pin: props.isPin,
      hide: props.isHideTools,
    },
    cache: {
      enable: false,
    },
    mode: props.mode,
    placeholder: props.placeholder,
    // TODO: 后期配置MdEditor的功能：菜单等等，参考 https://ld246.com/article/1549638745630#API
  });
});
// 获取vditor中获取
const getValue = () => {
  return vditor.value!.getValue();
};
// 获取vditor中渲染之后的html
const getHtml = () => {
  return vditor.value!.getHTML();
};
// 设置vditor中的值
const setValue = (val: string) => {
  vditor.value!.setValue(val);
};
// 设置vditor为只读模式
const disabled = () => {
  vditor.value!.disabled();
};

// 将上述方法向外暴露
defineExpose({
  getValue,
  getHtml,
  setValue,
  disabled,
});
</script>

<style>
.md {
  text-align: left;
}
</style>

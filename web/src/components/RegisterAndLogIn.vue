<!-- 完成注册，登录弹出框的构建 -->
<!-- 注册信息：username, pwd, nickname, email, avatar -->
<template>
  <el-dialog
    v-model="state.dialogVisible"
    title="登录+注册"
    width="30%"
    align-center
    @close="handleClose"
  >
    <el-form label-position="top" :model="params" style="max-width: 460px">
      <el-form-item label="用户名" :label-width="state.formLabelWidth">
        <el-input v-model="params.username" />
      </el-form-item>
      <el-form-item
        label="头像"
        :label-width="state.formLabelWidth"
        v-if="state.registerFlag"
      >
        <el-input v-model="params.password" />
      </el-form-item>
      <el-form-item
        label="昵称"
        :label-width="state.formLabelWidth"
        v-if="state.registerFlag"
      >
        <el-input v-model="params.nickname" />
      </el-form-item>
      <el-form-item label="密码" :label-width="state.formLabelWidth">
        <el-input v-model="params.password" />
      </el-form-item>
      <el-form-item
        label="邮箱"
        :label-width="state.formLabelWidth"
        v-if="state.registerFlag"
      >
        <el-input v-model="params.password" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" v-if="!state.registerFlag">登录</el-button>
        <el-button key="plain" @click="handleFlag" v-if="!state.registerFlag"
          >尚未注册，去注册>></el-button
        >
        <el-button type="primary" v-if="state.registerFlag">注册</el-button>
        <el-button key="plain" @click="handleFlag" v-if="state.registerFlag"
          >已注册，去登录>></el-button
        >
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<script setup lang="ts">
import { login } from "@/api/api_user";
import { reactive, watch } from "vue";

export interface Props {
  visible: boolean; // dialog组件是否可见
  flag?: boolean; // 是否是注册，控制nickname，email和avatar是否显示
}

const props = withDefaults(defineProps<Props>(), {
  visible: false,
  flag: false,
});

// 相关状态的定义
const state = reactive({
  dialogVisible: props.visible, // 通过控制父组件传值来确定dialog是否可见
  formLabelWidth: "60px",
  registerFlag: props.flag, // 是否是注册的dialog，默认是否
});

const params = reactive({
  username: "",
  password: "",
  nickname: "",
  email: "",
  avatar: "",
});

const handleFlag = () => {
  if (state.registerFlag === false) {
    state.registerFlag = true;
  } else {
    state.registerFlag = false;
  }
};

const userlogin = () => {
  // TODO: 需要对信息进行校验，比如是否是真邮箱，有没有特殊字符等等
  // 登录后，需要把对应的状态 userinfo, token等保存到userStore中
  // 登录后，应当跳转到对应的界面
  login({ UserName: params.username, Password: params.password });
};

const emits = defineEmits(["close"]);
const handleClose = () => {
  emits("close", false);
};

watch(props, () => {
  // state.dialogVisible = val
  console.log("watch数据");
  console.log(props);
  state.dialogVisible = props.visible;
  console.log(state);
});
</script>

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
      <!-- <el-form-item
        label="头像"
        :label-width="state.formLabelWidth"
        v-if="state.registerFlag"
      >
        <el-input v-model="params.avatar" />
      </el-form-item> -->
      <el-form-item
        label="昵称"
        :label-width="state.formLabelWidth"
        v-if="state.registerFlag"
      >
        <el-input v-model="params.nickname" />
      </el-form-item>
      <el-form-item label="密码" :label-width="state.formLabelWidth">
        <el-input v-model="params.password" show-password />
      </el-form-item>
      <el-form-item
        label="邮箱"
        :label-width="state.formLabelWidth"
        v-if="state.registerFlag"
      >
        <el-input v-model="params.email" />
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          v-if="!state.registerFlag"
          @click="handleLogIn"
          >登录</el-button
        >
        <el-button key="plain" @click="handleFlag" v-if="!state.registerFlag"
          >尚未注册，去注册>></el-button
        >
        <el-button
          type="primary"
          v-if="state.registerFlag"
          @click="handleSignUp"
          >注册</el-button
        >
        <el-button key="plain" @click="handleFlag" v-if="state.registerFlag"
          >已注册，去登录>></el-button
        >
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<script setup lang="ts">
import { login, register } from "@/api/api_user";
import { LogInUserData, RegisterUserData } from "@/model/request";
import { ElLoading, ElMessage } from "element-plus";
import { reactive, ref, watch } from "vue";
import { useUserStore } from "@/pinia/modules/user";
import router from "@/router";

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

const emits = defineEmits(["close"]);
const handleClose = () => {
  emits("close", false);
};

watch(props, () => {
  // state.dialogVisible = val
  state.dialogVisible = props.visible;
});

// 对输入内容进行校验，
const userStore = useUserStore();

const loadingIns = ref<any>(null);
const handleLogIn = async () => {
  if (!params.username) {
    ElMessage({
      message: "用户名不能为空",
      type: "warning",
    });
    return;
  } else if (!params.password) {
    ElMessage({
      message: "密码不能为空",
      type: "warning",
    });
    return;
  }
  // 调用登录api
  const loginData: LogInUserData = {
    UserName: params.username,
    Password: params.password,
  };
  loadingIns.value = ElLoading.service({
    fullscreen: true,
    text: "正在登录，请稍等...",
  });
  try {
    const { data } = await login(loginData);

    if (data.code === 0) {
      // 请求成功，用户登录,保存用户信息和token
      userStore.saveUserInfo(data.data.user);
      userStore.setToken(data.data.token);
      // userStore.setStatus(true);
      ElMessage({
        message: "登录成功，即将跳转",
        type: "success",
      });
      emits("close", false);
    }
  } catch (e) {
    loadingIns.value.close();
  }
  loadingIns.value.close();
};

const handleSignUp = async () => {
  const reg = new RegExp(
    "^[a-z0-9]+([._\\-]*[a-z0-9])*@([a-z0-9]+[-a-z0-9]*[a-z0-9]+.){1,63}[a-z0-9]+$"
  ); //正则表达式
  if (!params.username) {
    ElMessage({
      message: "用户名不能为空",
      type: "warning",
    });
    return;
  } else if (!params.password) {
    ElMessage({
      message: "密码不能为空",
      type: "warning",
    });
    return;
  } else if (!reg.test(params.email)) {
    ElMessage({
      message: "邮箱格式错误，请检查",
      type: "warning",
    });
    return; // TODO: 发邮件并输入验证码
  } else if (!params.nickname) {
    ElMessage({
      message: "昵称不能为空",
      type: "warning",
    });
    return;
  }
  const signupData: RegisterUserData = {
    UserName: params.username,
    Password: params.password,
    NickName: params.nickname,
    Avatar: params.avatar,
    Email: params.email,
  };
  loadingIns.value = ElLoading.service({
    fullscreen: true,
    text: "正在注册，请稍等...",
  });
  try {
    const { data } = await register(signupData);
    if (data.code === 0) {
      // TODO: 请求成功，注册用户成功，应当给出信息，并跳转到登录后的界面
      const { data } = await login({
        UserName: params.username,
        Password: params.password,
      });
      ElMessage({
        message: "注册成功，即将跳转",
      });
      if (data.code === 0) {
        userStore.setToken(data.data.token);
        userStore.saveUserInfo(data.data.user);
        // userStore.setStatus(true);
        emits("close", false);
      }
    }
  } catch (e) {
    loadingIns.value.close();
    ElMessage({
      message: "注册失败，请稍后再试",
      type: "error",
    });
  }
  loadingIns.value.close();
};
</script>

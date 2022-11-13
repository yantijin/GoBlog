import axios, { AxiosRequestConfig, AxiosResponse } from "axios";
import { emitter } from "@/utils/bus";
import { useUserStore } from "@/pinia/modules/user";
import { ElMessage, ElMessageBox } from "element-plus";

const service = axios.create({
  baseURL: import.meta.url,
  timeout: 99999,
});

let activeAxios = 0;
let timer: number;
const showLoading = () => {
  activeAxios++;
  if (timer) {
    clearTimeout(timer);
  }
  timer = setTimeout(() => {
    if (activeAxios > 0) {
      emitter.emit("showLoading");
    }
  });
};
// 关闭的loading界面
const closeLoading = () => {
  activeAxios--;
  if (activeAxios <= 0) {
    emitter.emit("closeLoading");
  }
};

// 请求拦截器
service.interceptors.request.use(
  (config: AxiosRequestConfig) => {
    // if (!config.doNotShowLoading) {
    showLoading();
    // }
    const userStore = useUserStore();
    config.headers = {
      "Content-Type": "application/json",
      "x-token": userStore.token,
      "x-user-id": userStore.userInfo.id,
      ...config.headers,
    };
    return config;
  },
  (error: any) => {
    closeLoading();
    ElMessage({
      showClose: true,
      message: error,
      type: "error",
    });
    return error;
  }
);

// 响应拦截器
service.interceptors.response.use(
  (response: AxiosResponse) => {
    const userStore = useUserStore();
    if (response.data.data === 0 || response.headers.success === "true") {
      if (response.headers.msg) {
        response.data.msg = decodeURI(response.headers.msg);
      }
    } else {
      ElMessage({
        showClose: true,
        message: response.data.msg || decodeURI(response.headers.msg!),
        type: "error",
      });
      if (response.data.data && response.data.data.reload) {
        userStore.token = "";
        localStorage.clear();
        // TODO: 这里需要使用router跳转到登录界面
      }
    }
    return response;
  },
  (error: any) => {
    closeLoading();
    if (!error.response) {
      ElMessageBox.confirm(
        `
      <p>检测到请求错误</p>
      <p>${error}</p>
      `,
        "请求错误",
        {
          dangerouslyUseHTMLString: true,
          distinguishCancelAndClose: true,
          confirmButtonText: "请稍后重试",
          cancelButtonText: "取消",
        }
      );
    }
    return error;
  }
);
export default service;

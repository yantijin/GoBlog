/*
  主要是用户相关的状态操作；
  token设置，接收，管理；
  用户信息设置接收管理；
  用户登录，登出
*/

import { UserInfo } from "@/model/user";
import { defineStore } from "pinia";
import { ref } from "vue";

export const useUserStore = defineStore("user", () => {
  const userInfo = ref<UserInfo>({
    uuid: "",
    nickname: "",
    avatar: "",
    id: "",
  });

  const token = ref(window.localStorage.getItem("token") || "");
  const saveUserInfo = (val: UserInfo) => {
    userInfo.value = val;
  };
  const setToken = (val: string) => {
    token.value = val;
  };
  const isLogIn = ref(false);
  const setStatus = (val: boolean) => {
    isLogIn.value = val;
  };

  return {
    userInfo,
    token,
    saveUserInfo,
    setToken,
    isLogIn,
    setStatus,
  };
});

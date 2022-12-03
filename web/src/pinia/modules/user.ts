/*
  主要是用户相关的状态操作；
  token设置，接收，管理；
  用户信息设置接收管理；
  用户登录，登出
*/

import { UserResponse } from "@/model/response";
import { defineStore } from "pinia";
import { ref, watch } from "vue";

export const useUserStore = defineStore("user", () => {
  // const userInfo = ref<UserInfo>({
  //   uuid: "",
  //   nickname: "",
  //   avatar: "",
  //   id: "",
  // });

  const userInfo = ref(
    JSON.parse(window.localStorage.getItem("userInfo")) || ""
  );

  const token = ref(window.localStorage.getItem("token") || "");
  const saveUserInfo = (val: UserResponse) => {
    userInfo.value = val;
    window.localStorage.setItem("userInfo", JSON.stringify(val));
  };
  const setToken = (val: string) => {
    token.value = val;
  };
  // const isLogIn = ref(false);
  // const setStatus = (val: boolean) => {
  //   isLogIn.value = val;
  // };

  const resetInfo = () => {
    userInfo.value = "";
    token.value = "";
  };

  watch(
    () => token.value,
    () => {
      window.localStorage.setItem("token", token.value);
      console.log(window.localStorage);
    }
  );

  return {
    userInfo,
    token,
    saveUserInfo,
    setToken,
    resetInfo,
    // isLogIn,
    // setStatus,
  };
});

import service from "@/utils/request";
import {
  ChangePwdData,
  LogInUserData,
  RegisterUserData,
  SetUserInfoData,
} from "@/model/request";

// @Summary 登录
// @Produce: application/json
// @Param data body: {username: "string", password: "string"}
// @Router /user/signin [post]
export const login = (data: LogInUserData) => {
  return service({
    url: "/user/signin",
    method: "post",
    data: data,
  });
};

// @Summary 注册
// @Produce application/json
// @Param data body
// @Router /user/signup [post]
export const register = (data: RegisterUserData) => {
  return service({
    url: "/user/signup",
    method: "post",
    data: data,
  });
};

// @Summary 修改密码
// @Produce application/json
// @Param data body:
// @Router /user/change_pwd [post]
export const changePwd = (data: ChangePwdData) => {
  return service({
    url: "/user/change_pwd",
    method: "post",
    data: data,
  });
};

// @Summary 获取当前登录用户的信息，用于后续编辑
// @Produce application/json
// @Router /user/getEditUser [get]
export const getUserInfo = () => {
  return service({
    url: "/user/getEditUser",
    method: "get",
  });
};

// @Summary 编辑更新当前用户信息
// @Accept application/json
// @Produce application/json
// Param data body: model.UserResponse
// @Router /user/postEditUser [post]
export const setUserInfo = (data: SetUserInfoData) => {
  return service({
    url: "user/postEditUser",
    method: "post",
    data: data,
  });
};

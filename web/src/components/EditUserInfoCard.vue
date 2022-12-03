<template>
  <el-card class="edit-container">
    <div class="form-container">
      <el-form class="el-form" label-width="120px">
        <el-form-item label="头像" style="margin-bottom: 0px">
          <el-upload
            class="avatar-uploader"
            :action="uploadUrl"
            accept=".jpg,.png,.jpeg"
            show-file-list="false"
            :on-success="handleAvatarSuccess"
            :before-upload="beforeAvatarUpload"
          >
            <el-avatar
              v-if="userInfo.avatar != ''"
              :src="userInfo.avatar"
              size="100"
            />
            <el-icon v-else class="avatar-uploader-icon"><Plus /></el-icon>
          </el-upload>
        </el-form-item>
        <el-form-item label="用户名">
          <el-input v-model="userInfo.username" />
        </el-form-item>
        <el-form-item label="昵称">
          <el-input v-model="userInfo.nickname" />
        </el-form-item>
        <el-form-item label="邮箱地址">
          <el-input v-model="userInfo.email" />
        </el-form-item>
      </el-form>
    </div>
    <el-button type="primary" id="btn" @click="handleChangeInfo"
      >保存</el-button
    >
  </el-card>
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue";
import { UploadProps } from "element-plus";
import { ElMessage } from "element-plus";
import { editUserInfo, setUserInfo } from "@/api/api_user";
import { UserInfoData } from "@/model/response";
import { useUserStore } from "@/pinia/modules/user";
import { useRouter } from "vue-router";

const userStore = useUserStore();
const router = useRouter();
const userInfo = ref<UserInfoData>({
  username: "",
  nickname: "",
  id: 1,
  email: "",
  avatar: "",
  uuid: "",
});
const avator = ref(
  "https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png"
); //图片显示
//可以将userid换成自己的参数
const uploadUrl = ref("");

const handleAvatarSuccess: UploadProps["onSuccess"] = (
  response,
  uploadFile
) => {
  avator.value = URL.createObjectURL(uploadFile.raw!);
};
const beforeAvatarUpload: UploadProps["beforeUpload"] = rawFile => {
  if (rawFile.size / 1024 / 1024 > 2) {
    ElMessage.error("图片不能超过2MB!");
    return false;
  }
  return true;
};

const obtainSelfUserInfo = async () => {
  const { data } = await editUserInfo();
  if (data.code == 0) {
    console.log("获取资料成功");
    console.log(data.data);
    return data.data;
  }
};

const handleChangeInfo = async () => {
  const { data } = await setUserInfo(userInfo.value);
  if (data.code == 0) {
    // 更新localStorage和userStore
    userStore.saveUserInfo(userInfo.value);
    router.push("/user/" + userStore.userInfo.id);
  }
};

onMounted(async () => {
  const uInfo = await obtainSelfUserInfo();
  userInfo.value = uInfo;
});
</script>

<style>
.edit-container {
  width: 70vw;
}
.avatar-uploader {
  height: 50px;
}

.el-form .el-form-item__label {
  font-size: 16px;
}

/* .el-form {
  margin: auto;
  text-align: center;
  border: 3px solid rgb(194, 189, 189);
} */

.el-form-item {
  margin-bottom: 20px;
  margin-top: 5px;
  margin-right: 10px;
}

.form-container {
  width: 50%;
  margin: auto;
  border: 3px solid rgb(194, 189, 189);
}

#btn {
  /* margin-left: auto; */
  /* flex-direction: column; */
  margin-top: 20px;
}
</style>

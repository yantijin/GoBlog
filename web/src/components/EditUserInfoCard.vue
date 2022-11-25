<template>
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
        <el-avatar v-if="avator" :src="avator" size="100" />
        <el-icon v-else class="avatar-uploader-icon"><Plus /></el-icon>
      </el-upload>
    </el-form-item>
    <el-form-item label="用户名">
      <el-input />
    </el-form-item>
    <el-form-item label="昵称">
      <el-input />
    </el-form-item>
    <el-form-item label="邮箱地址">
      <el-input />
    </el-form-item>
  </el-form>
  <el-button type="primary" id="btn">保存</el-button>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { UploadProps } from "element-plus";
import { ElMessage } from "element-plus";

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
</script>

<style>
.el-form {
  width: 70vw;
}
.avatar-uploader {
  height: 50px;
}

.el-form .el-form-item__label {
  font-size: 16px;
}

#btn {
  /* margin-left: auto; */
  /* flex-direction: column; */
  margin-top: 20px;
}
</style>

<template>
  <div
    class="profile"
    :style="{ backgroundImage: 'url(' + backgroundImage + ')' }"
  >
    <div v-if="isOwner" class="file is-light is-small change-bg">
      <el-upload accept=".jpg,.png,.jpeg" :show-file-list="false">
        <el-button text :icon="Upload" bg>设置背景</el-button>
      </el-upload>
    </div>
    <el-avatar
      :src="props.localUser.avatar"
      :round="true"
      size="large"
      class="profile-avatar"
    >
      <img
        src="https://cube.elemecdn.com/e/fd/0fc7d20532fdaf769a25683617711png.png"
      />
    </el-avatar>
    <div class="profile-info">
      <div class="metas">
        <h1 class="nickname">
          <router-link
            :to="'/user/' + props.localUser.id"
            :key="props.localUser.id"
            >{{ localUser.nickname }}</router-link
          >
        </h1>
        <div v-if="props.localUser.nickname" class="description">
          <p>{{ localUser.nickname }}</p>
        </div>
      </div>
      <div class="action-btns" v-if="userStore.userInfo.id != route.params.id">
        <el-button type="primary" @click="handleFollow">follow</el-button>
        <!-- <follow-btn
          v-if="!currentUser || currentUser.ID !== localUser.ID"
          :user-id="props.localUser.ID"
          :followed="followed"
          @onFollowed="onFollowed"
        /> -->
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { UserInfoData } from "@/model/response";
import { ref } from "vue";
import defaultImg from "@/assets/default-user-bg.jpg";
import { Upload } from "@element-plus/icons-vue";
import { useRoute } from "vue-router";
import { useUserStore } from "@/pinia/modules/user";

const route = useRoute();
const userStore = useUserStore();
const backgroundImage = ref(defaultImg);

const isOwner = ref(route.params.id == userStore.userInfo.id);

// const props = defineProps<{
//   localUser: UserInfoData
// }>();
const props = withDefaults(defineProps<{ localUser: UserInfoData }>(), {
  localUser: () => {
    return {
      nickname: "小白",
      id: 3,
      avatar:
        "https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png",
      email: "yantijin@163.com",
      username: "yantijin",
      uuid: "",
    };
  },
});

// const localUser: UserInfoData = {
//   nickname: "小白",
//   id: 3,
//   avatar: "https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png",
//   email: "yantijin@163.com",
//   username: "yantijin",
//   uuid: "",
// };

// const currentUser: UserInfoData = {
//   nickname: "小白",
//   id: 3,
//   avatar: "https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png",
//   email: "yantijin@163.com",
//   username: "yantijin",
//   uuid: "",
// };

const handleFollow = () => {
  console.log("hello");
  console.log(userStore.userInfo);
  console.log("hello");
  console.log(route.params.id);
  console.log(userStore.userInfo.id == route.params.id);
  // const res = obtainUserInfo;
  // console.log(res);
};
// const followed = ref(false)
</script>

<style lang="scss" scoped>
.profile {
  display: flex;
  margin-bottom: 10px;
  border-top-left-radius: 6px;
  border-top-right-radius: 6px;
  background-size: cover;
  background-position: 50%;
  height: 220px;
  width: 70vw;
  // filter: blur(2px) contrast(0.8);
  position: relative;
  .profile-avatar {
    position: absolute;
    top: 110px;
    left: 10px;
  }
  input[type="file"] {
    color: transparent;
  }
  .change-bg {
    position: absolute;
    top: 10px;
    right: 10px;
    opacity: 0.5;
    &:hover {
      opacity: 1;
    }
  }
  .profile-info {
    display: flex;
    justify-content: space-between;
    align-items: center;
    width: 100%;
    margin-top: 140px;
    padding: 10px 10px 10px 120px;
    background-image: linear-gradient(
      90deg,
      #ffffffff,
      rgba(255, 255, 255, 0.5),
      #dce9f200
    );
    .metas {
      display: flex;
      align-items: flex-start;
      flex-direction: column;
      width: 100%;
      .nickname {
        font-size: 18px;
        font-weight: 700;
        a {
          color: black;
          &:hover {
            color: black;
            text-decoration: underline;
          }
        }
      }
      .description {
        font-size: 14px;
        color: black;
      }
      .homepage {
        font-size: 14px;
        a {
          color: var(--color2);
          &:hover {
            color: var(--text-link-color);
            text-decoration: underline;
          }
        }
      }
    }
    .action-btns {
      margin-left: 10px;
    }
  }
}
.dark-mode {
  .profile-info {
    background-image: linear-gradient(
      90deg,
      #00000088,
      rgba(255, 255, 255, 0.5),
      #dce9f200
    );
  }
}
</style>

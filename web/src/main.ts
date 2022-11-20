import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";
import { store } from "@/pinia";
import ElementPlus from "element-plus";
import "element-plus/dist/index.css";
import router from "./router";

const app = createApp(App);
app.use(ElementPlus).use(store).use(router);
app.mount("#app");

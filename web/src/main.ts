import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";
import { store } from "@/pinia";
import ElementPlus from "element-plus";
import "element-plus/dist/index.css";

const app = createApp(App);
app.use(ElementPlus).use(store);
app.mount("#app");

import { createApp } from "vue";

import "./style.css";
import App from "./App.vue";
import { setupRouter } from "./routers";
import ElementPlusPlugins from "./plugins/element-plus";

import { InitWasm } from "../wasm";
InitWasm().then(() => {
  const app = createApp(App);
  setupRouter(app);
  ElementPlusPlugins.install(app);
  app.mount("#app");
});

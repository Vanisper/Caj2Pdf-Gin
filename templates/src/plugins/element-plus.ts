// src\plugins\element-plus.ts
import { App } from "vue";
// import * as ElIconModules from "@element-plus/icons-vue";
import { ElMessage, ElMessageBox } from "element-plus";
import "element-plus/theme-chalk/index.css";

export default {
  install(app: App) {
    // // 批量注册 Element Plus 图标组件
    // // 或者自定义 ElIconModules 列表
    // for (const iconName in ElIconModules) {
    //   app.component(iconName, (ElIconModules as any)[iconName]);
    // }

    // 将消息提示组件注册为全局方法
    app.config.globalProperties.$message = ElMessage;
    app.config.globalProperties.$msgBox = ElMessageBox;
  },
};

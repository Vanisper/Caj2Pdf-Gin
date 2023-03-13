import { ElMessage, ElMessageBox } from "element-plus";

declare module "vue" {
  // vue 全局属性
  export interface ComponentCustomProperties {
    $message: typeof ElMessage;
    $msgBox: typeof ElMessageBox;
  }
}

<script lang="ts" setup>
import { UploadFilled } from '@element-plus/icons-vue';
import { UploadRawFile, } from "element-plus";
import { Awaitable } from 'element-plus/es/utils';
import {getCurrentInstance, onMounted, ref, watch} from 'vue';
import { FileMd5 } from "@/utils/FileMd5";

const instance = getCurrentInstance();
const proxy = instance?.proxy;
const $message = proxy?.$message;

interface IFileListItem {
    name: string;
    url: string;
    md5: string;
    size: number;
    pdf?: {
      name: string,
      url: string
    }
}
function sleep(time: number): Promise<void> {
    return new Promise((resolve) => {
        const t = setTimeout(() => {
            clearTimeout(t)
            return resolve;
        }, time);
    });
}
function createWorker(workerContent: string) {
    let blob = new Blob([workerContent]);
    let url = window.URL.createObjectURL(blob);
    return new Worker(url, { name: "md5-worker" })

}

const fileList = ref<IFileListItem[]>([]);
// 返回false 并不传递给某个接口  而是前端处理
const uploadBefore = async (rawFile: UploadRawFile): Promise<Awaitable<void | undefined | null | boolean | File | Blob>> => {
    const typeArry = [".caj", ".kdh"];
    const type = rawFile.name.substring(rawFile.name.lastIndexOf('.')).toLowerCase();
    const isRightType = typeArry.indexOf(type) > -1;
    let temp: IFileListItem;
    let md5 = "";
    if (isRightType) {
        const arrayBuffer = await rawFile.arrayBuffer();
        const chunkSize = 1024000 * 5; // 1MB*5
        const bytes = new Uint8Array(arrayBuffer.slice(0, arrayBuffer.byteLength));
        md5 = FileMd5(bytes, chunkSize);
        const blobUrl = URL.createObjectURL(new Blob([arrayBuffer]));
        temp = {
            name: rawFile.name,
            url: blobUrl,
            md5,
            size: rawFile.size
        }
        const oldIndex = fileList.value.map(e => e.md5).indexOf(md5);
        if (oldIndex !== -1) {
            fileList.value[oldIndex] = temp;
        } else {
            fileList.value.push(temp);
        }
    } else {
        $message?.warning("请选择正确的文件格式");
        return false;
    }
    const formData = new FormData();
    formData.append('file', rawFile);
    fetch('/api/v1/upload', {
      method: 'POST',
      body: formData
    })
      .then(response => response.json())
      .then(result => {
        const oldIndex = fileList.value.map(e => e.md5).indexOf(md5);
        fileList.value[oldIndex].pdf = {
          name: result.name,
          url: result.url
        }
        $message?.success(result);
      })
      .catch(error => {
        $message?.error(error);
      });
    return false;
}

watch(fileList, (value) => {
    console.log(value);
}, {
    deep: true
})

const run = (index: number) => {
    console.log(fileList.value[index]);
}

onMounted(()=>{
  fetch("/api/ping").then(response => response.json())
      .then(data => console.log(data))
      .catch(error => console.error(error));
})
</script>

<template>
    <div v-for="(item, index) in fileList">
        <span :key="item.md5" @click="run(index)">{{ item.name }}</span>
    </div>
    <el-upload class="upload-demo" drag multiple :before-upload="uploadBefore" accept=".caj,.kdh">
        <el-icon class="el-icon--upload"><upload-filled /></el-icon>
        <div class="el-upload__text">
            Drop file here or <em>click to upload</em>
        </div>
        <template #tip>
            <div class="el-upload__tip">
                jpg/png files with a size less than 500kb
            </div>
        </template>
    </el-upload>
</template>

<style lang="less" scoped></style>
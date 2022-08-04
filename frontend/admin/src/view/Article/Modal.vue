<template>
  <n-modal
    v-model:show="show"
    preset="dialog"
    :title="title"
    :mask-closable="false"
    positive-text="确认"
    negative-text="取消"
    @positive-click="onOK"
    @negative-click="onCancel"
  >
    <template #default>
      <div class="main">
        <n-form
          ref="formRef"
          :model="form"
          :rules="rules"
          label-placement="left"
          label-width="auto"
          require-mark-placement="right-hanging"
          size="small"
          :inline="false"
        >
          <n-form-item label="标题" path="title">
            <n-input v-model:value="form.title" placeholder="标题" />
          </n-form-item>

          <n-form-item label="描述" path="desc">
            <n-input
              type="textarea"
              v-model:value="form.title"
              placeholder="描述"
            />
          </n-form-item>

          <n-form-item label="状态" path="state">
            <n-switch v-model:value="form.state" />
          </n-form-item>

          <n-form-item label="内容" path="content">
            <n-input
              type="textarea"
              v-model:value="form.content"
              placeholder="内容"
            />
          </n-form-item>

          <n-form-item label="封面图片">
            <n-upload
              action=""
              :default-file-list="fileList"
              list-type="image-card"
              :custom-request="customUpload"
              :max="1"
              accept="image/png, image/jpeg"
            >
              上传文件
            </n-upload>
          </n-form-item>
        </n-form>
      </div>
    </template>
  </n-modal>
</template>

<script lang="ts" setup>
import apis from "@/apis/apis.js";
import {
  NModal,
  NInput,
  NForm,
  NFormItem,
  NSwitch,
  NUpload,
  UploadFileInfo,
  useMessage,
  UploadCustomRequestOptions,
} from "naive-ui";
import { reactive, ref, watch } from "vue";

const message = useMessage();

const title = ref("编辑文章");
const show = ref(true);

const form = reactive({
  title: "",
  desc: "",
  state: 0,
  cover_image_url: "",
  content: "",
  modified_by: "zz",
});

const rules = {
  title: {
    required: true,
    trigger: ["blur", "input"],
    message: "请输入标题",
  },
  desc: {
    required: true,
    trigger: ["blur", "input"],
    message: "请输入描述",
  },
};
const onOK = () => {};

const onCancel = () => {};

// 上传相关
const fileList = ref<UploadFileInfo[]>([]);

// 文件列表变化的时候，更新表单字段
watch(fileList, (val) => {
  if (val.length > 0) {
    form.cover_image_url = fileList.value[0].url as string;
  } else {
    form.cover_image_url = "";
  }
});

const customUpload = async ({
  file,
  data,
  headers,
  withCredentials,
  action,
  onFinish,
  onError,
  onProgress,
}: UploadCustomRequestOptions) => {
  const formData = new FormData();
  formData.append("file", file.file as File);
  formData.append("type", "1");

  console.log(formData);

  const res = await apis.uploadFile(formData);

  if (res.code === 0) {
    fileList.value = [
      {
        ...file,
        url: res.data.file_access_url,
      },
    ];

    message.success("上传成功");
    onFinish();
  } else {
    message.error(res.msg);
    onError();
  }
};
</script>

<style lang="less" scoped>
.main {
  width: 80vw;
  height: 80vh;
}
</style>

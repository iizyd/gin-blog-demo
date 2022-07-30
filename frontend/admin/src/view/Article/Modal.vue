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
        </n-form>
      </div>
    </template>
  </n-modal>
</template>

<script lang="ts" setup>
import { NModal, NInput, NForm, NFormItem, NSwitch } from "naive-ui";
import { reactive, ref } from "vue";

const title = ref("编辑文章");
const show = ref(true);

const form = reactive({
  title: "",
  desc: "",
  state: 0,
  cover_image_url: "",
  content: "",
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
</script>

<style lang="less" scoped>
.main {
  width: 80vw;
  height: 80vh;
}
</style>

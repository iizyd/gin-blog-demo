<template>
  <div class="article-wrap">
    <n-space vertical>
      <n-data-table
        remote
        bordered
        ref="table"
        :columns="columns"
        :data="data"
        :loading="loading"
        :pagination="pagination"
        :row-key="rowKey"
        @update:page="handlePageChange"
        flex-height
        :style="{ height: 'calc(100vh - 85px)' }"
      />

      <Modal />
    </n-space>
  </div>
</template>

<script lang="ts" setup>
import {
  NDataTable,
  NSpace,
  NButton,
  PaginationInfo,
  NImage,
  NTag,
} from "naive-ui";
import { h, reactive, ref, VNode } from "vue";
import apis from "@/apis/apis";
import Modal from "./Modal.vue";

interface ArticleItem {
  id: number;
  title: string;
  desc: string;
  cover_image_url: string;
  content: string;
  modified_on: string;
  state: number;
}

const onEdit = (row: ArticleItem) => {
  console.log(row);
};

const onDel = (row: ArticleItem) => {
  console.log(row);
};

const columns = ref([
  {
    title: "ID",
    key: "id",
    ellipsis: {
      tooltip: true,
    },
    width: 70,
  },
  {
    title: "标题",
    key: "title",
    ellipsis: {
      tooltip: true,
    },
  },
  {
    title: "描述",
    key: "desc",
    ellipsis: {
      tooltip: true,
    },
  },
  {
    title: "封面",
    key: "cover_image_url",
    ellipsis: {
      tooltip: true,
    },
    render(row: ArticleItem): VNode {
      return h(
        NImage,
        {
          width: 40,
          src: row.cover_image_url,
          fallbackSrc:
            "https://picx.zhimg.com/80/v2-9c5b94a0fd324ed303373c0c7e3b208d_1440w.jpg?source=1940ef5c",
        },
        {}
      );
    },
  },
  {
    title: "内容",
    key: "content",
    ellipsis: {
      tooltip: true,
    },
  },
  {
    title: "修改时间",
    key: "modified_on",
    ellipsis: {
      tooltip: true,
    },
  },
  {
    title: "状态",
    key: "state",
    width: 100,
    render(row: ArticleItem) {
      return h(
        NTag,
        {
          type: row.state === 0 ? "error" : "success",
          size: "medium",
        },
        { default: () => (row.state === 0 ? "下线" : "上线") }
      );
    },
  },
  {
    title: "操作",
    key: "opt",
    width: 100,
    render(row: ArticleItem) {
      return h("span", {}, [
        h(
          NButton,
          {
            text: true,
            type: "info",
            size: "medium",
            onClick: () => onEdit(row),
          },
          { default: () => "编辑" }
        ),
        h(
          NButton,
          {
            text: true,
            type: "error",
            size: "medium",
            onClick: () => onDel(row),
          },
          { default: () => "删除" }
        ),
      ]);
    },
  },
]);

const data = ref<ArticleItem[]>([]);

const pagination = reactive({
  page: 1,
  pageCount: 1,
  pageSize: 10,
  itemCount: 100,
  prefix({ itemCount }: PaginationInfo) {
    return `共 ${itemCount} 条`;
  },
});

const rowKey = (rowData: ArticleItem) => {
  return rowData.id;
};

const loading = ref(false);

const handlePageChange = (currentPage: number) => {
  pagination.page = currentPage;
  getData();
};

const getData = async () => {
  loading.value = true;
  const res = await apis.getArticles({
    page: pagination.page,
    page_size: pagination.pageSize,
  });
  console.log(res);

  if (res.code === 0) {
    data.value = res.data.list;
    pagination.itemCount = res.data.total;
  }

  loading.value = false;
};
getData();
</script>

<style lang="less" scoped>
.article-wrap {
  height: 100%;

  :deep(.n-button + .n-button) {
    margin-left: 10px;
  }
}
</style>

<template>
  <n-notification-provider>
    <n-message-provider>
      <Layout>
        <router-view />
      </Layout>
    </n-message-provider>
  </n-notification-provider>
</template>

<script setup lang="ts">
import { NNotificationProvider, NMessageProvider } from "naive-ui";
import Layout from "@/view/Layout/Layout.vue";
</script>

<style lang="less" scoped></style>

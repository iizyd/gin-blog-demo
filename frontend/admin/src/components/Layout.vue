<template>
  <div class="layout-wrap">
    <div class="side-bar"></div>
    <div class="right-container">
      <Header />
    </div>
  </div>
</template>

<script lang="ts" setup>
import Header from '@/components/Header/index.vue'
</script>

<style lang="less" scoped>
.layout-wrap {
  display: flex;
  flex-direction: row;
  height: 100%;
}
.side-bar {
  height: 100%;
  width: 200px;
  background: #000;
}

.right-container {
  height: 100%;
  width: calc(100% - 200px);
  background: deepskyblue;
}
</style>

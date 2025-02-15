<script setup lang="ts">
import SideNav from './side-nav/index.vue'
</script>

<template>
    <!-- 侧边导航 -->
    <SideNav class="w-230px shrink-0 bg-#fafafa border-r-(1px #e5e5e5) hover:(shadow-[2px_0_2px_#e5e5e5])" />
    <!-- 主面板区域 -->
    <div class="grow p-20px bg-#f4f4f5 main">
        <RouterView />
    </div>
</template>

<style scoped>
.main {
    overflow-y: auto;
}

.main::-webkit-scrollbar {
    width: 0;
}
</style>
<script lang="ts" setup>
import { useLocalStorage } from '@vueuse/core';
import {
    HomeIcon,
    PackageOpenIcon,
    GhostIcon,
    SettingsIcon,
    UserRoundIcon,
    TelescopeIcon
} from 'lucide-vue-next';
import { computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const route = useRoute()
const router = useRouter()

const activatedNames = computed(() => {
    return route.matched.map(matched => matched.name)
})

// 读取token
const accessToken = useLocalStorage("access-token", () => {
    return localStorage.getItem("access-token")
})
</script>

<template>
    <!-- 侧边导航 -->
    <div class="pl-16px pr-32px">
        <div class="py-16px flex items-center">
            <img class="w-48px h-48px border-2px rounded-1/2" src="https://picsum.photos/64/64" />
            <div class="ml-8px">AI Memos</div>
        </div>
        <div class="flex flex-col gap-8px">
            <div v-if="accessToken !== null" @click="router.push('/')"
                :class="{ 'bg-#ffffff border-1px shadow-[0_0_2px_0_#f0f0f0]': activatedNames.includes('home') }"
                class="flex items-center gap-4px rounded-16px px-16px py-8px cursor-pointer hover:(bg-#f0f0f0)">
                <HomeIcon :size="18" />
                <div>首页</div>
            </div>
            <div @click="router.push('/explore')"
                :class="{ 'bg-#ffffff border-1px shadow-[0_0_2px_0_#f0f0f0]': activatedNames.includes('explore') }"
                class="flex items-center gap-4px rounded-16px px-16px py-8px cursor-pointer hover:(bg-#f0f0f0)">
                <TelescopeIcon :size="18" />
                <div>探索</div>
            </div>
            <div v-if="accessToken !== null"
                class="flex items-center gap-4px rounded-16px px-16px py-8px cursor-pointer hover:(bg-#f0f0f0)">
                <PackageOpenIcon :size="18" />
                <div>资源库</div>
            </div>
            <div v-if="accessToken !== null"
                class="flex items-center gap-4px rounded-16px px-16px py-8px cursor-pointer hover:(bg-#f0f0f0)">
                <GhostIcon :size="18" />
                <div>收藏</div>
            </div>
            <div v-if="accessToken === null" @click="router.push('/signin')"
                :class="{ 'bg-#ffffff border-1px shadow-[0_0_2px_0_#f0f0f0]': activatedNames.includes('signin') }"
                class="flex items-center gap-4px rounded-16px px-16px py-8px cursor-pointer hover:(bg-#f0f0f0)">
                <UserRoundIcon :size="18" />
                <div>登录</div>
            </div>
            <div v-if="accessToken !== null" @click="router.push('/settings')"
                :class="{ 'bg-#ffffff border-1px shadow-[0_0_2px_0_#f0f0f0]': activatedNames.includes('settings') }"
                class="flex items-center gap-4px rounded-16px px-16px py-8px cursor-pointer hover:(bg-#f0f0f0)">
                <SettingsIcon :size="18" />
                <div>设置</div>
            </div>
        </div>
    </div>
</template>
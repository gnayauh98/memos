<script lang="ts" setup>
import { EllipsisVerticalIcon, ListStartIcon, PencilIcon, XIcon } from 'lucide-vue-next';

defineProps<{
    author?: string
    updateDate: string
    content: string
}>()
</script>

<template>
    <div class="p-8px border-(1px #ececec) bg-#ffffff rounded-8px shadow-[0_0_2px_#ececec]">
        <!-- 卡片头部区域 -->
        <div class="flex items-center">
            <div class="text-0.75em font-600 text-#00009f cursor-pointer">{{ author ? `@${author}` : '' }}</div>
            <div class="ml-8px text-0.75em text-#808080">{{ updateDate }}</div>
            <div class="ml-auto">
                <div class="relative group">
                    <EllipsisVerticalIcon :size="16"
                        class="cursor-pointer text-#808080 rounded-1/2 group-hover:bg-#f0f0f0 hover:(bg-#f0f0f0)" />
                    <div
                        class="absolute z-99 border-1px rounded-8px px-0.5em py-0.25em bg-#f4f4f5 hidden group-hover:block top-16px right-0px w-100px">
                        <div class="flex gap-16px items-center cursor-pointer">
                            <PencilIcon :size="16" />
                            <span>编辑</span>
                        </div>
                        <div class="flex gap-16px items-center cursor-pointer">
                            <ListStartIcon :size="16" />
                            <span>置顶</span>
                        </div>
                        <div class="w-90% mx-auto h-2px bg-#aaaaaa my-8px"></div>
                        <div class="flex gap-16px items-center cursor-pointer">
                            <XIcon :size="16" />
                            <span>删除</span>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <!-- 卡片内容区域 -->
        <div class="mt-4px whitespace-break-spaces" v-html="content" />
    </div>
</template>
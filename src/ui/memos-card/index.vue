<script lang="ts" setup>
import { EllipsisVerticalIcon, ListStartIcon, PencilIcon, XIcon } from 'lucide-vue-next';
import Editor from '../editor/index.vue'
import { ref } from 'vue';
import { getMemoById } from '../../api/memo';

const props = defineProps<{
    id: string
    author?: string
    updateDate: string
    content: string
}>()

const emits = defineEmits(["update"])

const isShowEditor = ref(false)
const rawContent = ref("")

const getRawContent = async () => {
    // 获取内容
    const data = await getMemoById(props.id)
    rawContent.value = data.data.content
    isShowEditor.value = true
}

const onUpdate = (memo: any) => {
    console.log(memo)
    emits("update", memo)
    isShowEditor.value = false
}
</script>

<template>
    <div v-if="!isShowEditor" class="p-8px border-(1px #ececec) bg-#ffffff rounded-8px shadow-[0_0_2px_#ececec]">
        <!-- 卡片头部区域 -->
        <div class="flex items-center">
            <div class="text-0.75em font-600 text-#00009f cursor-pointer">{{ author ? `@${author}` : '' }}</div>
            <div class="ml-8px text-0.75em text-#808080">{{ updateDate }}</div>
            <div class="ml-auto">
                <div class="relative group">
                    <EllipsisVerticalIcon :size="16"
                        class="cursor-pointer text-#808080 rounded-1/2 group-hover:bg-#f0f0f0 hover:(bg-#f0f0f0)" />
                    <div
                        class="absolute z-99 text-12px border-1px rounded-8px px-0.5em py-0.25em bg-#f4f4f5 hidden group-hover:block top-16px right-0px w-70px">
                        <div @click="getRawContent" class="flex gap-4px items-center cursor-pointer">
                            <PencilIcon :size="12" />
                            <span>编辑</span>
                        </div>
                        <div class="flex gap-4px items-center cursor-pointer">
                            <ListStartIcon :size="12" />
                            <span>置顶</span>
                        </div>
                        <div class="w-90% mx-auto h-1px bg-#aaaaaa my-8px"></div>
                        <div class="flex gap-4px items-center cursor-pointer">
                            <XIcon :size="12" />
                            <span>删除</span>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <!-- 卡片内容区域 -->
        <div class="mt-4px whitespace-break-spaces" v-html="content" />
    </div>
    <Editor @update="onUpdate" :raw-content="rawContent" mode="update" :memo-id="id" v-else />
</template>
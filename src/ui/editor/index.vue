<script setup lang="ts">
import {
    HashIcon,
    PackageOpenIcon,
    SquareMousePointerIcon,
    SendIcon,
    LoaderIcon,
    MicVocalIcon
} from 'lucide-vue-next';
import { ref, useTemplateRef } from 'vue';
import { createMemo, updateMemo } from '../../api/memo';
import { RequestCode } from '../../api';
import SelectList from '../resources/SelectList.vue';
import { useEventListener } from '@vueuse/core';

const { mode = 'create', memoId, rawContent } = defineProps<{
    mode?: 'create' | 'update'
    memoId?: string
    rawContent?: string
}>()

const emits = defineEmits(["create", "update"])
const loading = ref(false)

const textArea = useTemplateRef("textarea")

const onCreate = async () => {
    loading.value = true
    const value = textArea.value?.value

    if (!value) {
        loading.value = false
        return
    }

    const { code, data } = await createMemo(value)

    if (code === RequestCode.REQUEST_SUCCESS) {
        emits("create", data)
        textArea.value.value = ""
    }
    loading.value = false
}

const onUpdate = async () => {
    loading.value = true
    const value = textArea.value?.value

    if (!value) {
        loading.value = false
        return
    }

    const { code, data } = await updateMemo(value, memoId!)

    if (code === RequestCode.REQUEST_SUCCESS) {
        emits("update", data)
        textArea.value.value = ""
    }
    loading.value = false
}

const onSubmit = async () => {
    if (mode === 'create') {
        onCreate()
    }
    if (mode === 'update' && memoId) {
        onUpdate()
    }
}

const isShowPackage = ref(false)
const onPackageClicked = () => {
    isShowPackage.value = !isShowPackage.value
}

const onSelectResource = (id: string) => {
    // console.log(url, textArea.value)
    if (textArea.value) {
        textArea.value.value += `${textArea.value.value.length ? "\n" : ""}![资源库内容](@${id}?w=60&h=0.6)`
    }
}

useEventListener('click', (event) => {
    const element = event.target as HTMLElement

    const packageOpen = element.closest(".package-open")

    if (!packageOpen) {
        isShowPackage.value = false
    }
})

</script>

<template>
    <!-- 编辑器 -->
    <div class=" bg-#ffffff border-(1px #ececec) shadow-[0_0_4px_#ececec] rounded-8px p-8px">
        <textarea :value="rawContent" placeholder="此刻的想法..." ref="textarea"
            class="w-full outline-none min-h-[calc(6*1.5em)]" />
        <!-- 快捷键 -->
        <div class="flex gap-0px items-center">
            <HashIcon class="cursor-pointer p-4px hover:bg-#ececec rounded-8px" :size="24" />
            <SquareMousePointerIcon class="cursor-pointer p-4px hover:bg-#ececec rounded-8px" :size="24" />
            <div class="relative package-open">
                <PackageOpenIcon @click="onPackageClicked" class="cursor-pointer p-4px hover:bg-#ececec rounded-8px"
                    :size="24" />
                <SelectList @select="onSelectResource" v-if="isShowPackage"
                    class="absolute top-8px -right-404px shadow-[0_0_6px_#aaaaaa]" />
            </div>
            <MicVocalIcon class="cursor-pointer p-4px hover:bg-#ececec rounded-8px" :size="24" />
            <div v-if="!loading" @click="onSubmit"
                class="ml-auto flex items-center cursor-pointer text-white bg-#52aaa0 hover:bg-#52aab1 py-4px px-8px rounded-8px">
                保存
                <SendIcon class="ml-8px" :size="16" />
            </div>
            <div v-else
                class="ml-auto flex items-center cursor-pointer text-white bg-#52aaa0 hover:bg-#52aab1 py-4px px-8px rounded-8px">
                保存
                <LoaderIcon class="ml-8px animate-spin" :size="16" />
            </div>
        </div>
    </div>
</template>
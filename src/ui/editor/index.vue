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
import { createMemo } from '../../api/memo';
import { RequestCode } from '../../api';

const emits = defineEmits(["create"])
const loading = ref(false)

const textArea = useTemplateRef("textarea")

const onSubmit = async () => {
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

</script>

<template>
    <!-- 编辑器 -->
    <div class=" bg-#ffffff border-(1px #ececec) shadow-[0_0_4px_#ececec] rounded-8px p-8px">
        <textarea placeholder="此刻的想法..." ref="textarea" class="w-full outline-none min-h-[calc(6*1.5em)]" />
        <!-- 快捷键 -->
        <div class="flex gap-8px items-center">
            <HashIcon class="cursor-pointer" :size="16" />
            <SquareMousePointerIcon class="cursor-pointer" :size="16" />
            <PackageOpenIcon class="cursor-pointer" :size="16" />
            <MicVocalIcon class="cursor-pointer" :size="16" />
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
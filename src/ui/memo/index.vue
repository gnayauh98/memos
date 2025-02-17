<script lang="ts" setup>
import { useRoute } from 'vue-router';
import { getMemoById } from '../../api/memo';
import { reactive } from 'vue';
import dayjs from 'dayjs';

const route = useRoute()

const memoModal = reactive({
    memo: {
        content: "",
        createAt: ""
    },
    loading: false
})

function load() {
    memoModal.loading = true
    const id = route.params.id as string;
    getMemoById(id).then(({ data }) => {
        memoModal.memo = { ...data }
    }).finally(() => {
        memoModal.loading = false
    })
}

load()

</script>

<template>
    <div class="max-w-896px mx-auto">
        <!-- <div class="border-1 p-16px rounded-16px">
            <div class="i-lucide:arrowLeft cursor-pointer"></div>
        </div> -->
        <div class="mt-16px px-16px">
            <div class="grid grid-cols-[64px_1fr] gap-x-16px grid-rows-1">
                <img class="rounded-1/2 w-64px h-64px border-(1 solid #acacac) col-span-0"
                    src="https://picsum.photos/64/64?r=1234" />
                <div class="col-span-1">
                    <div class="mt-6px font-600">Anqzi</div>
                    <div class="mt-6px text-12px text-#aaaaaa">
                        <span class="">#我的日记#任务#算法</span>
                        <span class="ml-16px">{{ dayjs(memoModal.memo.createAt ?? Date.now()).format("YYYY-MM-DD HH:mm")
                            }}</span>
                    </div>
                </div>
            </div>
            <div class="mt-16px" v-html="memoModal.memo.content"></div>
            <div class="h-1px w-full my-32px bg-gray-200"></div>
            <div>
                <div class="text-#acacac">引用</div>
                <div class="mt-4px flex flex-col items-start">
                    <div class=" py-4px flex items-center cursor-pointer">
                        <span class="i-lucide:link text-blue inline-block">1</span>
                        <span class="ml-4px">2025年2月17日日记</span>
                    </div>
                    <div class=" py-4px flex items-center cursor-pointer">
                        <span class="i-lucide:link text-blue inline-block">1</span>
                        <span class="ml-4px">2025年2月17日日记</span>
                    </div>
                    <div class=" py-4px flex items-center cursor-pointer">
                        <span class="i-lucide:link text-blue inline-block">1</span>
                        <span class="ml-4px">2025年2月17日日记</span>
                    </div>
                </div>
            </div>
            <div class="mt-16px">
                <div class="text-#acacac">被引用</div>
                <div class="mt-4px flex flex-col items-start">
                    <div class=" py-4px flex items-center cursor-pointer">
                        <span class="i-lucide:link text-blue inline-block">1</span>
                        <span class="ml-4px">2025年2月17日日记</span>
                    </div>
                    <div class=" py-4px flex items-center cursor-pointer">
                        <span class="i-lucide:link text-blue inline-block">1</span>
                        <span class="ml-4px">2025年2月17日日记</span>
                    </div>
                    <div class=" py-4px flex items-center cursor-pointer">
                        <span class="i-lucide:link text-blue inline-block">1</span>
                        <span class="ml-4px">2025年2月17日日记</span>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
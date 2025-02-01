<script lang="ts" setup>
import { onMounted, ref } from 'vue';
import { queryMemos } from '../api/memo';
import dayjs from 'dayjs';

const memoList = ref<any[]>([])
const isMore = ref(true)
const loading = ref()
const pageNo = ref(1)

onMounted(async () => {
    const data = await queryMemos(pageNo.value)
    pageNo.value += 1

    console.log(data)
    memoList.value = data.data

    const observer = new IntersectionObserver(async (entries) => {
        const entry = entries?.[0]
        if (!entry) {
            return
        }
        if (entry.isIntersecting) {
            const data = await queryMemos(pageNo.value)
            memoList.value = [...memoList.value, ...data.data]
            if (data?.data?.length < 10) {
                isMore.value = false
                observer.unobserve(loading.value)
            }
        }
    })

    observer.observe(loading.value)
})
</script>

<template>
    <div class="max-w-768px mx-auto">
        <div v-for="memo in memoList" :key="memo.id"
            class="mt-32px first:mt-0 bg-white rounded-8px p-8px border-(1px solid #ececec)">
            <div class="flex items-center relative pl-36px">
                <img class="absolute -left-16px -top-16px object-cover w-48px h-48px rounded-1/2 border-(3px solid #f4f4f5)"
                    :src="`https://picsum.photos/200/200?r=${memo.id}`" />
                <div class="ml-4px">
                    <div>星光</div>
                    <div class="text-0.8em text-#808080">发布于 {{ dayjs(memo.createAt).format("YYYY/MM/DD hh:mm") }}</div>
                </div>
            </div>
            <div class="mt-8px ">
                <div v-html="memo.content"></div>
            </div>
        </div>
        <div v-if="isMore" ref="loading" class="flex flex-col items-center mt-16px">
            <div
                class="mx-auto relative animate-spin w-20px h-20px rounded-1/2 border-2px bg-gradient-conic bg-gradient-from-blue bg-gradient-to-lime before:(content-[''] absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 rounded-1/2 w-75% h-75% bg-white)">
            </div>
            <div class="text-#aaaaaa">加载中...</div>
        </div>
        <div v-else class="flex flex-col items-center mt-16px">
            <div class="text-#aaaaaa">没有更多了</div>
        </div>
    </div>
</template>
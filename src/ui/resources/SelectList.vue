<script lang="ts" setup>
import { computed, onMounted, ref } from 'vue';
import { getResources } from '../../api/resource';
import Placeholder from './Placeholder.vue';

const resourceList = ref<{
    url: string
    id: string
}[]>([])

const loading = ref(true)

const emits = defineEmits(["select"])

const onSelect = (id: string) => {
    emits('select', id)
}

const resourceModal = computed(() => ({
    rowsNum: Math.ceil(resourceList.value.length / 2),
    totalRowsNum: resourceList.value.length,
    isOddNumber: resourceList.value.length % 2 === 1
}))

onMounted(async () => {
    loading.value = true
    const data = await getResources()

    resourceList.value = data
    loading.value = false
})

</script>

<template>
    <div
        class="overflow-hidden flex flex-wrap gap-8px items-stretch overflow-auto bg-white border-(1px solid) rounded-8px">
        <div v-if="loading" class="flex p-8px flex-col items-center">
            <!-- <div
                class="mx-auto relative animate-spin w-20px h-20px rounded-1/2 border-2px bg-gradient-conic bg-gradient-from-blue bg-gradient-to-lime before:(content-[''] absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 rounded-1/2 w-75% h-75% bg-white)">
            </div>
            <div class="text-#aaaaaa">加载中...</div> -->
            <div class="grid grid-cols-2 gap-col-1 gap-row-1">
                <Placeholder />
                <Placeholder />
                <Placeholder />
                <Placeholder />
            </div>
        </div>
        <div v-else class="max-h-300px p-8px overflow-auto grid grid-cols-2 gap-col-1 gap-row-1">
            <img v-for="(resource, index) in resourceList" :key="resource.id" @click="onSelect(resource.id)"
                class="w-200px bg-white object-contain hover:(shadow-[0_0_1px_#888888] scale-150) flex-grow-1 aspect-16/9 border-1 rounded-8px cursor-pointer"
                :class="{
                    'origin-tl': index === 0,
                    'origin-tr': index === 1,
                    'origin-bl': ((Math.floor(index / 2) + 1) === resourceModal.rowsNum && index % 2 === 0),
                    'origin-rb': ((Math.floor(index / 2) + 1) === resourceModal.rowsNum && index % 2 === 1),
                    'origin-l': (![1, resourceModal.rowsNum].includes(Math.floor(index / 2) + 1) && index % 2 === 0),
                    'origin-r': (![1, resourceModal.rowsNum].includes(Math.floor(index / 2) + 1) && index % 2 === 1)
                }" loading="lazy" :src="resource.url" />
        </div>
    </div>
</template>
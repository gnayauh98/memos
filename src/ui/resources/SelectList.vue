<script lang="ts" setup>
import { onMounted, ref } from 'vue';
import { getResources } from '../../api/resource';

const resourceList = ref<{
    url: string
    id: string
}[]>([])

const loading = ref(true)

const emits = defineEmits(["select"])

const onSelect = (id: string) => {
    emits('select', id)
}

onMounted(async () => {
    loading.value = true
    const data = await getResources()

    resourceList.value = data
    loading.value = false
})

</script>

<template>
    <div
        class="w-400px max-h-500px flex flex-wrap gap-8px items-stretch overflow-auto bg-white border-(1px solid) rounded-8px p-8px">
        <div v-if="loading" class="flex flex-col items-center">
            <div
                class="mx-auto relative animate-spin w-20px h-20px rounded-1/2 border-2px bg-gradient-conic bg-gradient-from-blue bg-gradient-to-lime before:(content-[''] absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 rounded-1/2 w-75% h-75% bg-white)">
            </div>
            <div class="text-#aaaaaa">加载中...</div>
        </div>
        <template v-for="resource in resourceList" :key="resource.id">
            <img @click="onSelect(resource.id)" class="w-180px flex-grow-1 aspect-16/9 cursor-pointer" loading="lazy"
                :src="resource.url" />
        </template>
    </div>
</template>
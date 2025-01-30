<script setup lang="ts">
import { GridIcon, ImageUpIcon, ListIcon } from "lucide-vue-next"
import { onMounted, ref } from "vue";
import { createResource, getResources } from "../../api/resource";
import { RequestCode } from "../../api";

// const isOpened = ref(false)
const inputRef = ref<HTMLInputElement>()
const resourceList = ref<{
  url: string,
  id: string
}[]>([])

function onUploadClicked() {
  // if (isOpened.value) {
  //   return
  // }
  // isOpened.value = true
  inputRef.value?.click()
}

async function onUploadInput() {
  // isOpened.value = false
  const file = inputRef.value?.files?.[0]
  if (!file) {
    return
  }

  // 上传
  const data = await createResource(file)

  console.log(data)

  if (data.code != RequestCode.REQUEST_SUCCESS) {
    return
  }



  resourceList.value = [{
    url: data.data.url,
    id: data.data.id
  }, ...resourceList.value]
}

onMounted(async () => {
  const data = await getResources()

  resourceList.value = [...data, ...resourceList.value]
})
</script>

<template>
  <div class="max-w-768px mx-auto">
    <div>资源库</div>
    <!-- 上传区域 -->
    <div class="w-full h-200px p-16px relative mt-16px">
      <input ref="inputRef" @input="onUploadInput" class="w-full h-full border-1px rounded-8px" type="file"
        accept="image/*" />
      <div @click="onUploadClicked"
        class="w-full h-full cursor-pointer border-1px rounded-8px flex flex-col items-center justify-center absolute top-0 left-0 p-16px bg-white">
        <ImageUpIcon :size="64" class="text-#cacaca" />
        <div class="text-#cacaca">点击或拖拽上传资源</div>
      </div>
    </div>
    <div class="flex mt-16px items-center justify-start">
      <div class="ml-auto flex gap-8px border-1px p-4px rounded-4px">
        <GridIcon class="cursor-pointer" :size="16" />
        <ListIcon class="cursor-pointer" :size="16" />
      </div>
      <!-- <div class="ml-auto bg-green px-1em py-0.25em rounded-0.25em">上传</div> -->
    </div>
    <div class="mt-16px flex flex-wrap gap-x-8px gap-y-16px">
      <div class="w-200px h-120px" v-for="resource in resourceList" :key="resource.id">
        <img class="w-full h-full border-(1px solid #aaaaaa) rounded-8px object-cover" :src="resource.url" />
        <!-- <div class="text-center">图片{{ resource.id }}</div> -->
      </div>
    </div>
  </div>
</template>

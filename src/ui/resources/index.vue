<script setup lang="ts">
  import {GridIcon,ListIcon} from "lucide-vue-next"
</script>

<template>
  <div>
    <div>资源库</div>
    <div class="flex items-center justify-center">
      <div class="flex gap-8px border-1px p-4px rounded-4px">
        <GridIcon class="cursor-pointer" :size="16" />
        <ListIcon class="cursor-pointer" :size="16" />
      </div>
      <div class="ml-auto bg-green px-1em py-0.25em rounded-0.25em">上传</div>
    </div>
    <div class="mt-20px flex flex-wrap gap-x-8px gap-y-16px">
      <div class="grow-1" v-for="index in 30" :key="index">
        <img class="w-full rounded-8px" :src="`https://picsum.photos/200/120?r=${index}`" />
        <div class="text-center">图片{{ index }}</div>
      </div>
    </div>
  </div>
</template>


<script setup lang="ts">
import { computed } from 'vue'
import { NH3, NP, NImage, NListItem } from 'naive-ui'
import { config } from '@/config'
import type { News } from '@/lib'
const props = defineProps<{
  news: News
}>()

const coverUrl = computed(() => {
  return props.news.cover && props.news.cover !== ''
    ? `${config.apiBaseUrl}/api/files/news/${props.news.id}/${props.news.cover}`
    : ''
})
</script>

<template>
  <n-list-item>
    <template #prefix>
      <n-image width="100" :src="coverUrl" />
    </template>
    <n-h3>{{ props.news.title }}</n-h3>
    <n-p>
      <span v-html="props.news.body" />
    </n-p>
  </n-list-item>
</template>

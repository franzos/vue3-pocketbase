<script setup lang="ts">
import { useNewsStore } from '@/stores/news'
import { storeToRefs } from 'pinia'
import { NH1, NAlert, NList } from 'naive-ui'
import News from '../components/News.vue'
import { onMounted } from 'vue'
import { config } from '@/config'
const newsStore = useNewsStore()
const { news } = storeToRefs(newsStore)

onMounted(async () => {
  await newsStore.fetch()
})
</script>

<template>
  <main>
    <n-h1>News</n-h1>
    <n-list>
      <News v-for="n in news?.items" :key="n.id" :news="n" />
    </n-list>
    <n-alert
      v-if="!news || news.items.length === 0"
      :show-icon="false"
      title="No news found"
      type="default"
    >
      Add some on the <a :href="`${config.apiBaseUrl}/_/#/collection`">backend</a>.
    </n-alert>
  </main>
</template>

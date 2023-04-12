import { ref } from 'vue'
import { defineStore } from 'pinia'
import type { ListResult, RecordListQueryParams } from 'pocketbase'
import type { News } from '@/lib'
import { usePocketbaseStore } from './pb'

export const useNewsStore = defineStore('news', () => {
  const pb = usePocketbaseStore()
  const news = ref<ListResult<News>>()

  async function fetch(params?: { from?: number; to?: number; params?: RecordListQueryParams }) {
    const res = await pb.getCollection<News>('news', params)
    news.value = res
  }

  return {
    news,
    fetch,
    isBusy: pb.isBusy,
    apiErrors: pb.apiErrors
  }
})

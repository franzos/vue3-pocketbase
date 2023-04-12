import { computed, onMounted, ref } from 'vue'
import { defineStore } from 'pinia'
import PocketBase, {
  type ListResult,
  type RecordListQueryParams,
  type Record,
  ClientResponseError
} from 'pocketbase'
import type { Registration } from '@/lib'
import { config } from '@/config'

export const usePocketbaseStore = defineStore('pocketbase', () => {
  const api = new PocketBase(config.apiBaseUrl)
  const apiErrors = ref<ClientResponseError>()
  const isBusy = ref(false)

  const user = ref<Record | undefined>(undefined)

  const isLoggedIn = computed(() => {
    return user.value !== undefined
  })

  onMounted(async () => {
    const cookie = localStorage.getItem('pocketbase_auth')
    if (cookie) {
      api.authStore.loadFromCookie(cookie)
      try {
        const res = await api.collection('users').authRefresh()
        user.value = res.record
      } finally {
        isBusy.value = false
      }
    }
  })

  async function signup(data: Registration) {
    isBusy.value = true
    try {
      await api.collection('users').create(data)
      // await pb.collection('users').requestVerification('test@example.com');
      apiErrors.value = undefined
      return true
    } catch (error) {
      apiErrors.value = new ClientResponseError(error)
    } finally {
      isBusy.value = false
    }
  }

  async function login(identifier: string, password: string) {
    isBusy.value = true
    try {
      const res = await api.collection('users').authWithPassword(identifier, password)
      apiErrors.value = undefined
      user.value = res.record
    } catch (error) {
      apiErrors.value = new ClientResponseError(error)
    } finally {
      isBusy.value = false
    }
  }

  async function logout() {
    api.authStore.clear()
    user.value = undefined
  }

  async function getCollection<T>(
    collection: string,
    params?: {
      from?: number
      to?: number
      query?: RecordListQueryParams
    }
  ): Promise<ListResult<T> | undefined> {
    isBusy.value = true
    try {
      const res = await api
        .collection(collection)
        .getList<T>(params?.from || 1, params?.to || 10, params?.query || {})
      apiErrors.value = undefined
      return res
    } catch (error) {
      apiErrors.value = new ClientResponseError(error)
    } finally {
      isBusy.value = false
    }
  }

  return { api, apiErrors, isBusy, isLoggedIn, user, login, signup, logout, getCollection }
})

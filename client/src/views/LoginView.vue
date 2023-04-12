<script setup lang="ts">
import { usePocketbaseStore } from '@/stores/pb'
import { storeToRefs } from 'pinia'
import { computed } from 'vue'
import { NAlert } from 'naive-ui'
import Login from '../components/Login.vue'

const pb = usePocketbaseStore()
const { isLoggedIn, user } = storeToRefs(pb)

const username = computed(() => {
  return user.value?.name
})
</script>

<template>
  <main>
    <div v-if="isLoggedIn">
      <n-alert type="success" closable>
        <template #default> Welcome back, {{ username }}. </template>
      </n-alert>
    </div>
    <div v-else>
      <Login />
    </div>
  </main>
</template>

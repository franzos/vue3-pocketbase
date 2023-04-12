<script setup lang="ts">
import { computed } from 'vue'
import { NAlert, NSpace } from 'naive-ui'
import type { ClientResponseError } from 'pocketbase'

const props = defineProps<{
  errors: ClientResponseError | undefined
}>()

const errors = computed(() => (props.errors ? props.errors.data.data : undefined))
const errorMessage = computed(() => (props.errors ? props.errors.message : 'Unknown error'))
</script>

<template>
  <n-space :vertical="true" v-if="errors">
    <n-alert v-if="errorMessage" type="error" closable>
      <template #default>
        {{ errorMessage }}
      </template>
    </n-alert>
    <n-alert v-for="(value, key) in errors" :key="key" type="error" closable>
      <template #default> {{ key }}: {{ value.message }} </template>
    </n-alert>
  </n-space>
</template>

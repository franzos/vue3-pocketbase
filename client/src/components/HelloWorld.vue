<script setup lang="ts">
import { NH1, NH3, NP } from 'naive-ui'
defineProps<{
  msg: string
}>()
</script>

<template>
  <n-h1>{{ msg }}</n-h1>
  <n-h3>
    <n-p>
      You’ve successfully created a project with
      <a href="https://vitejs.dev/" target="_blank" rel="noopener">Vite</a> +
      <a href="https://vuejs.org/" target="_blank" rel="noopener">Vue 3</a>. What's next?
    </n-p>
  </n-h3>
</template>

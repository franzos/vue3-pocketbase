<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { NMenu, NImage, type MenuOption } from 'naive-ui'
import { computed, h, ref, type Ref } from 'vue'
import { usePocketbaseStore } from './stores/pb'
import { storeToRefs } from 'pinia'
import Button from 'naive-ui/es/button/src/Button'

const pb = usePocketbaseStore()
const { isLoggedIn, user } = storeToRefs(pb)
const username = computed(() => {
  return user.value?.name
})

const activeKey = ref('home')

const menuOptions: Ref<MenuOption[]> = computed(() => {
  return [
    {
      label: () =>
        h(
          NImage,
          {
            src: '/src/assets/logo.svg',
            width: 12,
            height: 12,
            fit: 'cover'
          },
          {
            default: () => 'Home'
          }
        )
    },
    {
      label: () =>
        h(
          RouterLink,
          {
            label: {
              name: 'home'
            },
            to: '/'
          },
          {
            default: () => 'Home'
          }
        ),
      key: 'home'
    },
    {
      label: () =>
        h(
          RouterLink,
          {
            label: {
              name: 'news'
            },
            to: '/news'
          },
          {
            default: () => 'News'
          }
        ),
      key: 'news'
    },
    {
      label: () =>
        h(
          RouterLink,
          {
            to: {
              name: 'about'
            }
          },
          {
            default: () => 'About'
          }
        ),
      key: 'about'
    },
    {
      label: () =>
        h(
          RouterLink,
          {
            to: {
              name: 'login'
            }
          },
          {
            default: () => 'Login'
          }
        ),
      key: 'login',
      show: !isLoggedIn.value
    },
    {
      label: () =>
        h(
          Button,
          {
            onClick: () => {
              pb.logout()
            }
          },
          {
            default: () => `Logout (${username.value})`
          }
        ),
      key: 'logout',
      show: isLoggedIn.value
    }
  ]
})
</script>

<template>
  <header>
    <n-menu v-model:value="activeKey" mode="horizontal" :options="menuOptions" />
  </header>

  <div class="container">
    <RouterView />
  </div>
</template>

<style>
.container {
  padding: 1rem;
}
</style>

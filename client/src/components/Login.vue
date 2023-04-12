<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { storeToRefs } from 'pinia'
import {
  NForm,
  NFormItem,
  NInput,
  NButton,
  NSpace,
  type FormInst,
  type FormValidationError
} from 'naive-ui'
import Errors from '@/components/Errors.vue'
import { usePocketbaseStore } from '@/stores/pb'

const pb = usePocketbaseStore()
const { apiErrors, isBusy } = storeToRefs(pb)
const router = useRouter()
const route = useRoute()

const formRef = ref<FormInst | null>(null)

const form = ref({
  // username or email
  identifier: '',
  password: ''
})

const formRules = ref({
  identifier: [
    {
      required: true,
      message: 'Username or Email is required'
    }
  ],
  password: [
    {
      required: true,
      message: 'Password is required'
    }
  ]
})

onMounted(() => {
  if (route.query) {
    form.value.identifier = (route.query.identifier as string) || ''
  }
})

async function login(event: MouseEvent) {
  isBusy.value = true
  event.preventDefault()
  formRef.value?.validate(async (errors: Array<FormValidationError> | undefined) => {
    if (!errors) {
      await pb.login(form.value.identifier, form.value.password)
    } else {
      alert(errors)
    }
  })
}

function signup() {
  apiErrors.value = undefined
  router.push('/sign-up')
}
</script>

<template>
  <n-form ref="formRef" :model="form" :rules="formRules">
    <n-form-item label="Username or Email:">
      <n-input v-model:value="form.identifier" />
    </n-form-item>

    <n-form-item label="Password:">
      <n-input type="password" v-model:value="form.password" />
    </n-form-item>

    <Errors :errors="apiErrors" />

    <n-space>
      <n-button type="primary" native-type="submit" :disabled="isBusy" @click="login">
        Login
      </n-button>
      <n-button quaternary :disabled="isBusy" @click="signup">Sign up</n-button>
    </n-space>
  </n-form>
</template>

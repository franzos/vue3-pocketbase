<script setup lang="ts">
import { ref, type Ref } from 'vue'
import { useRouter } from 'vue-router'
import { storeToRefs } from 'pinia'
import {
  NForm,
  NFormItem,
  NInput,
  NButton,
  NCheckbox,
  NSpace,
  type FormInst,
  type FormItemRule,
  type FormValidationError
} from 'naive-ui'
import Errors from '@/components/Errors.vue'
import { usePocketbaseStore } from '@/stores/pb'
import type { Registration } from '@/lib'

const pb = usePocketbaseStore()
const { apiErrors, isBusy } = storeToRefs(pb)
const router = useRouter()

const formRef = ref<FormInst | null>(null)

const form: Ref<Registration> = ref({
  name: 'test',
  username: 'test_username',
  email: 'test@example.com',
  emailVisibility: true,
  password: '12345678',
  passwordConfirm: '12345678'
})

const formRules = ref({
  name: [
    {
      required: true,
      message: 'Name is required'
    }
  ],
  username: [
    {
      required: true,
      message: 'Username is required'
    }
  ],
  email: [
    {
      message: 'Email is required'
    }
  ],
  password: [
    {
      required: true,
      message: 'Password is required'
    }
  ],
  passwordConfirm: [
    {
      required: true,
      message: 'Password confirmation is required',
      trigger: ['input', 'blur']
    },
    {
      validator: validatePasswordSame,
      message: 'Password is not same as re-entered password!',
      trigger: ['blur', 'password-input']
    }
  ]
})

function validatePasswordSame(rule: FormItemRule, value: string): boolean {
  return value === form.value.password
}

async function signup(event: MouseEvent) {
  isBusy.value = true
  event.preventDefault()
  formRef.value?.validate(async (errors: Array<FormValidationError> | undefined) => {
    if (!errors) {
      const success = await pb.signup(form.value)
      if (success) {
        router.push({ name: 'login', query: { identifier: form.value.username } })
      }
    } else {
      alert(errors)
    }
  })
}

function cancel() {
  apiErrors.value = undefined
  router.push('/login')
}
</script>

<template>
  <n-form ref="formRef" :model="form" :rules="formRules">
    <n-form-item label="Name:">
      <n-input v-model:value="form.name" />
    </n-form-item>

    <n-form-item label="Username:">
      <n-input v-model:value="form.username" />
    </n-form-item>

    <n-form-item label="Email:">
      <n-input type="email" v-model:value="form.email" />
    </n-form-item>

    <n-form-item label="Email Visibility:">
      <n-checkbox v-model="form.emailVisibility" />
    </n-form-item>

    <n-form-item label="Password:">
      <n-input type="password" v-model:value="form.password" />
    </n-form-item>

    <n-form-item label="Confirm Password:">
      <n-input type="password" v-model:value="form.passwordConfirm" />
    </n-form-item>

    <Errors :errors="apiErrors" />

    <n-space>
      <n-button type="primary" native-type="submit" :disabled="isBusy" @click="signup"
        >Sign Up</n-button
      >
      <n-button :disabled="isBusy" @click="cancel">Cancel</n-button>
    </n-space>
  </n-form>
</template>

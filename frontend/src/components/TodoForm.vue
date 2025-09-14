<template>
  <v-card class="mb-6" elevation="4">
    <v-card-title class="text-h5 text-center">
      <v-icon left>mdi-plus-circle</v-icon>
      添加新任务
    </v-card-title>

    <v-card-text>
      <v-form @submit.prevent="handleSubmit">
        <v-text-field
          v-model="formData.name"
          label="任务名称"
          placeholder="请输入任务名称"
          required
          outlined
          :rules="[rules.required]"
          prepend-inner-icon="mdi-clipboard-text"
          class="mb-4"
        />

        <v-text-field
          v-model="formData.description"
          label="任务描述"
          placeholder="请输入任务描述（可选）"
          outlined
          prepend-inner-icon="mdi-text"
          class="mb-4"
        />

        <v-btn
          type="submit"
          color="primary"
          size="large"
          block
          :loading="loading"
          :disabled="!formData.name.trim()"
        >
          <v-icon left>mdi-plus</v-icon>
          {{ loading ? '创建中...' : '创建任务' }}
        </v-btn>
      </v-form>
    </v-card-text>
  </v-card>
</template>

<script setup lang="ts">
  import { ref } from 'vue'
  import type { CreateTodoRequest } from '@/types/todo'

  const emit = defineEmits<{
    submit: [todoData: CreateTodoRequest]
  }>()

  defineProps<{
    loading?: boolean
  }>()

  const formData = ref<CreateTodoRequest>({
    name: '',
    description: ''
  })

  const rules = {
    required: (value: string) => !!value || '此字段为必填项'
  }

  const handleSubmit = (): void => {
    if (!formData.value.name.trim()) {
      return
    }

    emit('submit', { ...formData.value })

    // 清空表单
    formData.value = {
      name: '',
      description: ''
    }
  }
</script>

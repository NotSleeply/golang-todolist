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

<template>
  <t-card class="todo-form" header="添加新任务">
    <template #header>
      <div class="card-header">
        <t-icon name="add-circle" />
        <span class="header-title">添加新任务</span>
      </div>
    </template>

    <t-form @submit="handleSubmit" layout="vertical">
      <t-form-item label="任务名称" name="name" :rules="[rules.required]">
        <t-input v-model="formData.name" placeholder="请输入任务名称" clearable />
      </t-form-item>

      <t-form-item label="任务描述" name="description">
        <t-input v-model="formData.description" placeholder="请输入任务描述（可选）" clearable />
      </t-form-item>

      <t-form-item>
        <t-button
          type="submit"
          theme="primary"
          size="large"
          block
          :loading="loading"
          :disabled="!formData.name.trim()"
        >
          <template #icon>
            <t-icon name="add" />
          </template>
          {{ loading ? '创建中...' : '创建任务' }}
        </t-button>
      </t-form-item>
    </t-form>
  </t-card>
</template>

<style scoped>
  .todo-form {
    margin-bottom: 24px;
  }

  .card-header {
    display: flex;
    align-items: center;
    gap: 8px;
  }

  .header-title {
    font-size: 18px;
    font-weight: 600;
  }
</style>

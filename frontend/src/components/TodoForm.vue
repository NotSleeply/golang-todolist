<template>
  <div class="todo-form">
    <h2>添加新任务</h2>
    <form @submit.prevent="handleSubmit">
      <div class="form-group">
        <input
          v-model="formData.name"
          type="text"
          placeholder="任务名称"
          required
          class="form-input"
        />
      </div>
      <div class="form-group">
        <input
          v-model="formData.description"
          type="text"
          placeholder="任务描述"
          class="form-input"
        />
      </div>
      <button type="submit" class="btn btn-primary" :disabled="loading">
        <span v-if="loading">创建中...</span>
        <span v-else>创建任务</span>
      </button>
    </form>
  </div>
</template>

<script setup lang="ts">
  import { ref } from 'vue'
  import type { CreateTodoRequest } from '@/types/todo'

  const emit = defineEmits<{
    submit: [todoData: CreateTodoRequest]
  }>()

  const props = defineProps<{
    loading?: boolean
  }>()

  const formData = ref<CreateTodoRequest>({
    name: '',
    description: ''
  })

  const handleSubmit = (): void => {
    if (!formData.value.name.trim()) {
      alert('请输入任务名称')
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

<style scoped>
  .todo-form {
    background: rgba(255, 255, 255, 0.1);
    backdrop-filter: blur(10px);
    border-radius: 16px;
    padding: 2rem;
    margin-bottom: 2rem;
  }

  .todo-form h2 {
    text-align: center;
    color: white;
    margin-bottom: 1.5rem;
    font-size: 1.5rem;
  }

  .form-group {
    margin-bottom: 1rem;
  }

  .form-input {
    width: 100%;
    padding: 0.75rem 1rem;
    border: none;
    border-radius: 8px;
    background: rgba(255, 255, 255, 0.9);
    font-size: 1rem;
    transition: all 0.3s ease;
  }

  .form-input:focus {
    outline: none;
    background: white;
    transform: translateY(-2px);
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  }

  .btn {
    padding: 0.75rem 2rem;
    border: none;
    border-radius: 8px;
    font-size: 1rem;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.3s ease;
    width: 100%;
  }

  .btn-primary {
    background: linear-gradient(45deg, #667eea, #764ba2);
    color: white;
  }

  .btn-primary:hover:not(:disabled) {
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
  }

  .btn:disabled {
    opacity: 0.6;
    cursor: not-allowed;
  }

  @media (max-width: 768px) {
    .todo-form {
      padding: 1.5rem;
    }
  }
</style>

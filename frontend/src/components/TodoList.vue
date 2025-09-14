<template>
  <div class="todo-list">
    <h2>任务列表</h2>

    <div v-if="loading" class="loading">
      <div class="spinner"></div>
      <span>加载中...</span>
    </div>

    <div v-else-if="todos.length === 0" class="no-todos">
      <div class="empty-icon">📝</div>
      <p>暂无任务，开始创建你的第一个任务吧！</p>
    </div>

    <div v-else class="todos-container">
      <TodoItem
        v-for="todo in todos"
        :key="todo.id"
        :todo="todo"
        @update="$emit('update', $event)"
        @delete="$emit('delete', $event)"
        @complete="$emit('complete', $event)"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
  import TodoItem from './TodoItem.vue'
  import type { Todo } from '@/types/todo'

  defineEmits<{
    update: [todo: Todo]
    delete: [id: string]
    complete: [todo: Todo]
  }>()

  defineProps<{
    todos: Todo[]
    loading?: boolean
  }>()
</script>

<style scoped>
  .todo-list {
    background: rgba(255, 255, 255, 0.1);
    backdrop-filter: blur(10px);
    border-radius: 16px;
    padding: 2rem;
  }

  .todo-list h2 {
    text-align: center;
    color: white;
    margin-bottom: 1.5rem;
    font-size: 1.5rem;
  }

  .loading {
    text-align: center;
    color: white;
    padding: 2rem;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 1rem;
  }

  .spinner {
    width: 40px;
    height: 40px;
    border: 4px solid rgba(255, 255, 255, 0.3);
    border-top: 4px solid white;
    border-radius: 50%;
    animation: spin 1s linear infinite;
  }

  @keyframes spin {
    0% {
      transform: rotate(0deg);
    }

    100% {
      transform: rotate(360deg);
    }
  }

  .no-todos {
    text-align: center;
    color: white;
    padding: 3rem 2rem;
  }

  .empty-icon {
    font-size: 4rem;
    margin-bottom: 1rem;
    opacity: 0.7;
  }

  .no-todos p {
    font-size: 1.1rem;
    opacity: 0.8;
  }

  .todos-container {
    display: flex;
    flex-direction: column;
    gap: 1rem;
  }

  @media (max-width: 768px) {
    .todo-list {
      padding: 1.5rem;
    }
  }
</style>

<script setup lang="ts">
  import { onMounted } from 'vue'
  import { useTodoStore } from '@/stores/todo'
  import TodoForm from '@/components/TodoForm.vue'
  import TodoList from '@/components/TodoList.vue'
  import type { CreateTodoRequest, UpdateTodoRequest, Todo } from '@/types/todo'

  const todoStore = useTodoStore()

  const handleCreateTodo = async (todoData: CreateTodoRequest): Promise<void> => {
    try {
      await todoStore.createTodo(todoData)
    } catch (error) {
      console.error('创建任务失败:', error)
    }
  }

  const handleUpdateTodo = async (todoData: UpdateTodoRequest): Promise<void> => {
    try {
      await todoStore.updateTodo(todoData)
    } catch (error) {
      console.error('更新任务失败:', error)
    }
  }

  const handleDeleteTodo = async (id: string): Promise<void> => {
    if (confirm('确定要删除这个任务吗？')) {
      try {
        await todoStore.deleteTodo(id)
      } catch (error) {
        console.error('删除任务失败:', error)
      }
    }
  }

  const handleCompleteTodo = async (todo: Todo): Promise<void> => {
    try {
      await todoStore.updateTodo({
        ...todo,
        completed: true
      })
    } catch (error) {
      console.error('完成任务失败:', error)
    }
  }

  onMounted(() => {
    todoStore.fetchTodos()
  })
</script>

<template>
  <div class="home">
    <div class="container">
      <h1 class="title">{{ $route.meta.title || 'Vue 3 TodoList' }}</h1>

      <!-- 统计信息 -->
      <div class="stats">
        <div class="stat-item">
          <span class="stat-number">{{ todoStore.totalTodos }}</span>
          <span class="stat-label">总任务</span>
        </div>
        <div class="stat-item">
          <span class="stat-number">{{ todoStore.pendingTodos.length }}</span>
          <span class="stat-label">待完成</span>
        </div>
        <div class="stat-item">
          <span class="stat-number">{{ todoStore.completedTodos.length }}</span>
          <span class="stat-label">已完成</span>
        </div>
      </div>

      <!-- 添加任务组件 -->
      <TodoForm @submit="handleCreateTodo" />

      <!-- 错误提示 -->
      <div v-if="todoStore.error" class="error-message">
        {{ todoStore.error }}
        <button @click="todoStore.clearError" class="error-close">&times;</button>
      </div>

      <!-- 任务列表组件 -->
      <TodoList
        :todos="todoStore.todos"
        :loading="todoStore.loading"
        @update="handleUpdateTodo"
        @delete="handleDeleteTodo"
        @complete="handleCompleteTodo"
      />
    </div>
  </div>
</template>

<style scoped>
  .home {
    min-height: 100vh;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    padding: 2rem 0;
  }

  .container {
    max-width: 800px;
    margin: 0 auto;
    padding: 0 1rem;
  }

  .title {
    text-align: center;
    color: white;
    font-size: 2.5rem;
    margin-bottom: 2rem;
    text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
  }

  .stats {
    display: flex;
    justify-content: center;
    gap: 2rem;
    margin-bottom: 2rem;
  }

  .stat-item {
    background: rgba(255, 255, 255, 0.1);
    backdrop-filter: blur(10px);
    border-radius: 12px;
    padding: 1.5rem;
    text-align: center;
    color: white;
    min-width: 120px;
  }

  .stat-number {
    display: block;
    font-size: 2rem;
    font-weight: bold;
    margin-bottom: 0.5rem;
  }

  .stat-label {
    font-size: 0.9rem;
    opacity: 0.8;
  }

  .error-message {
    background: #ff6b6b;
    color: white;
    padding: 1rem;
    border-radius: 8px;
    margin-bottom: 1rem;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .error-close {
    background: none;
    border: none;
    color: white;
    font-size: 1.5rem;
    cursor: pointer;
    padding: 0;
    width: 24px;
    height: 24px;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  @media (max-width: 768px) {
    .stats {
      flex-direction: column;
      align-items: center;
    }

    .title {
      font-size: 2rem;
    }
  }
</style>

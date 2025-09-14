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
  <div class="home-page">
    <t-space direction="vertical" size="24px" style="width: 100%">
      <!-- 页面标题 -->
      <t-card class="title-card">
        <div class="title-content">
          <t-icon name="format-list-checks" size="32px" />
          <h1 class="page-title">Vue 3 TodoList</h1>
        </div>
      </t-card>

      <!-- 统计信息 -->
      <t-row :gutter="16">
        <t-col :span="4">
          <t-card class="stat-card stat-primary">
            <div class="stat-content">
              <t-icon name="format-list-bulleted" size="32px" class="stat-icon" />
              <div class="stat-number">{{ todoStore.totalTodos }}</div>
              <div class="stat-label">总任务</div>
            </div>
          </t-card>
        </t-col>
        <t-col :span="4">
          <t-card class="stat-card stat-warning">
            <div class="stat-content">
              <t-icon name="time" size="32px" class="stat-icon" />
              <div class="stat-number">{{ todoStore.pendingTodos.length }}</div>
              <div class="stat-label">待完成</div>
            </div>
          </t-card>
        </t-col>
        <t-col :span="4">
          <t-card class="stat-card stat-success">
            <div class="stat-content">
              <t-icon name="check-circle" size="32px" class="stat-icon" />
              <div class="stat-number">{{ todoStore.completedTodos.length }}</div>
              <div class="stat-label">已完成</div>
            </div>
          </t-card>
        </t-col>
      </t-row>

      <!-- 添加任务组件 -->
      <TodoForm @submit="handleCreateTodo" :loading="todoStore.loading" />

      <!-- 错误提示 -->
      <t-alert
        v-if="todoStore.error"
        theme="error"
        message="操作失败"
        :description="todoStore.error"
        close
        @close="todoStore.clearError"
      />

      <!-- 任务列表组件 -->
      <TodoList
        :todos="todoStore.todos"
        :loading="todoStore.loading"
        @update="handleUpdateTodo"
        @delete="handleDeleteTodo"
        @complete="handleCompleteTodo"
      />
    </t-space>
  </div>
</template>

<style scoped>
  .home-page {
    max-width: 1200px;
    margin: 0 auto;
    padding: 24px;
  }

  .title-card {
    text-align: center;
  }

  .title-content {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 12px;
    padding: 24px;
  }

  .page-title {
    margin: 0;
    font-size: 28px;
    font-weight: 600;
    color: var(--td-text-color-primary);
  }

  .stat-card {
    text-align: center;
    cursor: pointer;
    transition: all 0.3s ease;
  }

  .stat-card:hover {
    transform: translateY(-2px);
    box-shadow: 0 4px 16px rgba(0, 0, 0, 0.1);
  }

  .stat-content {
    padding: 20px;
  }

  .stat-icon {
    margin-bottom: 8px;
  }

  .stat-number {
    font-size: 24px;
    font-weight: 600;
    margin-bottom: 4px;
  }

  .stat-label {
    font-size: 14px;
    color: var(--td-text-color-secondary);
  }

  .stat-primary .stat-icon,
  .stat-primary .stat-number {
    color: var(--td-brand-color);
  }

  .stat-warning .stat-icon,
  .stat-warning .stat-number {
    color: var(--td-warning-color);
  }

  .stat-success .stat-icon,
  .stat-success .stat-number {
    color: var(--td-success-color);
  }
</style>

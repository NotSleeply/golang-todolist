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
  <v-container class="py-8">
    <v-row justify="center">
      <v-col cols="12" md="10" lg="8">
        <!-- 页面标题 -->
        <v-card class="mb-6 text-center" elevation="8">
          <v-card-title class="text-h3 pa-6">
            <v-icon left size="large" color="primary">mdi-format-list-checks</v-icon>
            {{ $route.meta.title || 'Vue 3 TodoList' }}
          </v-card-title>
        </v-card>

        <!-- 统计信息 -->
        <v-row class="mb-6">
          <v-col cols="12" sm="4">
            <v-card class="text-center" color="primary" dark elevation="4">
              <v-card-text>
                <v-icon size="48" class="mb-2">mdi-format-list-bulleted</v-icon>
                <div class="text-h4 font-weight-bold">{{ todoStore.totalTodos }}</div>
                <div class="text-subtitle1">总任务</div>
              </v-card-text>
            </v-card>
          </v-col>
          <v-col cols="12" sm="4">
            <v-card class="text-center" color="warning" dark elevation="4">
              <v-card-text>
                <v-icon size="48" class="mb-2">mdi-clock-outline</v-icon>
                <div class="text-h4 font-weight-bold">{{ todoStore.pendingTodos.length }}</div>
                <div class="text-subtitle1">待完成</div>
              </v-card-text>
            </v-card>
          </v-col>
          <v-col cols="12" sm="4">
            <v-card class="text-center" color="success" dark elevation="4">
              <v-card-text>
                <v-icon size="48" class="mb-2">mdi-check-circle</v-icon>
                <div class="text-h4 font-weight-bold">{{ todoStore.completedTodos.length }}</div>
                <div class="text-subtitle1">已完成</div>
              </v-card-text>
            </v-card>
          </v-col>
        </v-row>

        <!-- 添加任务组件 -->
        <TodoForm @submit="handleCreateTodo" :loading="todoStore.loading" />

        <!-- 错误提示 -->
        <v-alert
          v-if="todoStore.error"
          type="error"
          dismissible
          @click:close="todoStore.clearError"
          class="mb-4"
        >
          {{ todoStore.error }}
        </v-alert>

        <!-- 任务列表组件 -->
        <TodoList
          :todos="todoStore.todos"
          :loading="todoStore.loading"
          @update="handleUpdateTodo"
          @delete="handleDeleteTodo"
          @complete="handleCompleteTodo"
        />
      </v-col>
    </v-row>
  </v-container>
</template>

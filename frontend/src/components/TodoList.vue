<template>
  <t-card class="todo-list">
    <template #header>
      <div class="card-header">
        <t-icon name="format-list-bulleted" />
        <span class="header-title">任务列表</span>
      </div>
    </template>

    <div v-if="loading" class="loading-container">
      <t-loading size="large" />
      <div class="loading-text">加载中...</div>
    </div>

    <div v-else-if="todos.length === 0" class="empty-container">
      <t-icon name="clipboard" size="64px" class="empty-icon" />
      <div class="empty-title">暂无任务</div>
      <div class="empty-desc">开始创建你的第一个任务吧！</div>
    </div>

    <div v-else class="todo-items">
      <TodoItem
        v-for="todo in todos"
        :key="todo.id"
        :todo="todo"
        @update="$emit('update', $event)"
        @delete="$emit('delete', $event)"
        @complete="$emit('complete', $event)"
      />
    </div>
  </t-card>
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
.card-header {
  display: flex;
  align-items: center;
  gap: 8px;
}

.header-title {
  font-size: 18px;
  font-weight: 600;
}

.loading-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 48px;
  gap: 16px;
}

.loading-text {
  font-size: 16px;
  color: var(--td-text-color-secondary);
}

.empty-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 48px;
  gap: 16px;
}

.empty-icon {
  color: var(--td-text-color-disabled);
}

.empty-title {
  font-size: 18px;
  font-weight: 500;
  color: var(--td-text-color-secondary);
}

.empty-desc {
  font-size: 14px;
  color: var(--td-text-color-placeholder);
}

.todo-items {
  padding: 16px;
}
</style>

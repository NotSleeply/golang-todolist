<template>
  <v-card elevation="4">
    <v-card-title class="text-h5 text-center">
      <v-icon left>mdi-format-list-checks</v-icon>
      任务列表
    </v-card-title>

    <v-card-text>
      <div v-if="loading" class="text-center py-8">
        <v-progress-circular indeterminate color="primary" size="64" />
        <div class="mt-4 text-h6">加载中...</div>
      </div>

      <div v-else-if="todos.length === 0" class="text-center py-8">
        <v-icon size="80" color="grey-lighten-1" class="mb-4"> mdi-clipboard-text-outline </v-icon>
        <div class="text-h6 grey--text">暂无任务</div>
        <div class="text-body-1 grey--text">开始创建你的第一个任务吧！</div>
      </div>

      <div v-else>
        <TodoItem
          v-for="todo in todos"
          :key="todo.id"
          :todo="todo"
          @update="$emit('update', $event)"
          @delete="$emit('delete', $event)"
          @complete="$emit('complete', $event)"
        />
      </div>
    </v-card-text>
  </v-card>
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

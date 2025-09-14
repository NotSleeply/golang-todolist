<script setup lang="ts">
  import { ref, reactive } from 'vue'
  import type { Todo, UpdateTodoRequest } from '@/types/todo'

  const emit = defineEmits<{
    update: [todo: UpdateTodoRequest]
    delete: [id: string]
    complete: [todo: Todo]
  }>()

  const props = defineProps<{
    todo: Todo
  }>()

  const isEditing = ref<boolean>(false)
  const editData = reactive<{
    name: string
    description: string
  }>({
    name: '',
    description: ''
  })

  const rules = {
    required: (value: string) => !!value || '此字段为必填项'
  }

  const toggleEdit = (): void => {
    if (isEditing.value) {
      cancelEdit()
    } else {
      startEdit()
    }
  }

  const startEdit = (): void => {
    isEditing.value = true
    editData.name = props.todo.name
    editData.description = props.todo.description || ''
  }

  const cancelEdit = (): void => {
    isEditing.value = false
    editData.name = ''
    editData.description = ''
  }

  const handleUpdate = (): void => {
    if (!editData.name.trim()) {
      return
    }

    const updateData: UpdateTodoRequest = {
      id: props.todo.id,
      name: editData.name.trim(),
      description: editData.description.trim(),
      completed: props.todo.completed
    }

    emit('update', updateData)
    cancelEdit()
  }
</script>

<template>
  <t-card :class="{ 'completed-todo': todo.completed }" class="todo-item" hover>
    <div class="todo-content">
      <div class="todo-info">
        <h3 :class="{ 'completed-text': todo.completed }" class="todo-title">
          {{ todo.name }}
        </h3>
        <p v-if="todo.description" :class="{ 'completed-text': todo.completed }" class="todo-desc">
          {{ todo.description }}
        </p>
      </div>

      <div class="todo-actions">
        <t-button
          v-if="!todo.completed"
          @click="$emit('complete', todo)"
          theme="success"
          variant="outline"
          size="small"
          shape="circle"
        >
          <template #icon>
            <t-icon name="check" />
          </template>
        </t-button>

        <t-button
          @click="toggleEdit"
          :theme="isEditing ? 'warning' : 'primary'"
          variant="outline"
          size="small"
          shape="circle"
        >
          <template #icon>
            <t-icon :name="isEditing ? 'close' : 'edit'" />
          </template>
        </t-button>

        <t-button
          @click="$emit('delete', todo.id)"
          theme="danger"
          variant="outline"
          size="small"
          shape="circle"
        >
          <template #icon>
            <t-icon name="delete" />
          </template>
        </t-button>
      </div>
    </div>

    <!-- 编辑表单 -->
    <div v-if="isEditing" class="edit-form">
      <t-form @submit.prevent="handleUpdate" layout="vertical">
        <t-form-item label="任务名称" name="name" :rules="[rules.required]">
          <t-input v-model="editData.name" placeholder="请输入任务名称" clearable />
        </t-form-item>

        <t-form-item label="任务描述" name="description">
          <t-input v-model="editData.description" placeholder="请输入任务描述（可选）" clearable />
        </t-form-item>

        <div class="form-actions">
          <t-button type="submit" theme="primary" size="small" :disabled="!editData.name.trim()">
            <template #icon>
              <t-icon name="check" />
            </template>
            保存
          </t-button>

          <t-button @click="cancelEdit" theme="default" variant="outline" size="small">
            <template #icon>
              <t-icon name="close" />
            </template>
            取消
          </t-button>
        </div>
      </t-form>
    </div>
  </t-card>
</template>

<style scoped>
  .todo-actions {
    display: flex;
    justify-content: flex-end; /* 让按钮靠右 */
    align-items: center;
    gap: 16px; /* 按钮之间的距离 */
    margin-top: 12px;
  }

  .todo-actions .t-button {
    font-size: 24px; /* 图标和按钮变大 */
    width: 40px;
    height: 40px;
    padding: 0;
  }
</style>

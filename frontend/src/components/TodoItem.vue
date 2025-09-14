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
  <v-card
    :class="{ 'bg-grey-lighten-4': todo.completed }"
    class="mb-3"
    variant="outlined"
    :elevation="todo.completed ? 1 : 2"
  >
    <v-card-text>
      <div class="d-flex align-center">
        <div class="flex-grow-1">
          <h3
            :class="{ 'text-decoration-line-through text-grey': todo.completed }"
            class="text-h6 mb-1"
          >
            {{ todo.name }}
          </h3>
          <p
            v-if="todo.description"
            :class="{ 'text-decoration-line-through text-grey': todo.completed }"
            class="text-body-2 ma-0"
          >
            {{ todo.description }}
          </p>
        </div>

        <div class="d-flex gap-2 ml-4">
          <v-btn
            v-if="!todo.completed"
            @click="$emit('complete', todo)"
            color="success"
            icon="mdi-check"
            size="small"
            variant="outlined"
          ></v-btn>

          <v-btn
            @click="toggleEdit"
            :color="isEditing ? 'warning' : 'primary'"
            :icon="isEditing ? 'mdi-close' : 'mdi-pencil'"
            size="small"
            variant="outlined"
          ></v-btn>

          <v-btn
            @click="$emit('delete', todo.id)"
            color="error"
            icon="mdi-delete"
            size="small"
            variant="outlined"
          ></v-btn>
        </div>
      </div>

      <!-- 编辑表单 -->
      <v-expand-transition>
        <div v-if="isEditing" class="mt-4 pt-4" style="border-top: 1px solid #e0e0e0">
          <v-form @submit.prevent="handleUpdate">
            <v-text-field
              v-model="editData.name"
              label="任务名称"
              placeholder="请输入任务名称"
              required
              variant="outlined"
              density="compact"
              :rules="[rules.required]"
              class="mb-3"
            />

            <v-text-field
              v-model="editData.description"
              label="任务描述"
              placeholder="请输入任务描述（可选）"
              variant="outlined"
              density="compact"
              class="mb-3"
            />

            <div class="d-flex gap-2">
              <v-btn type="submit" color="primary" variant="flat" :disabled="!editData.name.trim()">
                <v-icon left>mdi-content-save</v-icon>
                保存
              </v-btn>

              <v-btn @click="cancelEdit" color="grey" variant="outlined">
                <v-icon left>mdi-cancel</v-icon>
                取消
              </v-btn>
            </div>
          </v-form>
        </div>
      </v-expand-transition>
    </v-card-text>
  </v-card>
</template>

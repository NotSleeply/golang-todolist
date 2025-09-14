<template>
  <div class="todo-item" :class="{ completed: todo.completed }">
    <div class="todo-content">
      <h3 :class="{ completed: todo.completed }">{{ todo.name }}</h3>
      <p v-if="todo.description" :class="{ completed: todo.completed }">
        {{ todo.description }}
      </p>
    </div>
    
    <div class="todo-actions">
      <button 
        v-if="!todo.completed"
        @click="$emit('complete', todo)" 
        class="btn btn-success"
        title="标记为完成"
      >
        ✓
      </button>
      <button 
        @click="toggleEdit" 
        class="btn btn-info"
        :title="isEditing ? '取消编辑' : '编辑任务'"
      >
        {{ isEditing ? '✕' : '✏️' }}
      </button>
      <button 
        @click="$emit('delete', todo.id)" 
        class="btn btn-danger"
        title="删除任务"
      >
        🗑️
      </button>
    </div>
    
    <!-- 编辑表单 -->
    <div v-if="isEditing" class="edit-form">
      <form @submit.prevent="handleUpdate">
        <div class="form-group">
          <input 
            v-model="editData.name" 
            type="text" 
            placeholder="任务名称"
            required
            class="form-input"
          >
        </div>
        <div class="form-group">
          <input 
            v-model="editData.description" 
            type="text" 
            placeholder="任务描述"
            class="form-input"
          >
        </div>
        <div class="form-actions">
          <button type="submit" class="btn btn-primary">保存</button>
          <button type="button" @click="cancelEdit" class="btn btn-secondary">取消</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'

const emit = defineEmits(['update', 'delete', 'complete'])

const props = defineProps({
  todo: {
    type: Object,
    required: true
  }
})

const isEditing = ref(false)
const editData = reactive({
  name: '',
  description: ''
})

const toggleEdit = () => {
  if (isEditing.value) {
    cancelEdit()
  } else {
    startEdit()
  }
}

const startEdit = () => {
  isEditing.value = true
  editData.name = props.todo.name
  editData.description = props.todo.description || ''
}

const cancelEdit = () => {
  isEditing.value = false
  editData.name = ''
  editData.description = ''
}

const handleUpdate = () => {
  if (!editData.name.trim()) {
    alert('请输入任务名称')
    return
  }

  emit('update', {
    id: props.todo.id,
    name: editData.name,
    description: editData.description,
    completed: props.todo.completed
  })

  cancelEdit()
}
</script>

<style scoped>
.todo-item {
  background: rgba(255, 255, 255, 0.9);
  border-radius: 12px;
  padding: 1.5rem;
  transition: all 0.3s ease;
}

.todo-item:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0,0,0,0.2);
}

.todo-item.completed {
  opacity: 0.7;
  background: rgba(255, 255, 255, 0.6);
}

.todo-content {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  margin-bottom: 1rem;
}

.todo-content h3 {
  margin: 0;
  color: #333;
  font-size: 1.2rem;
  transition: all 0.3s ease;
}

.todo-content h3.completed {
  text-decoration: line-through;
  color: #999;
}

.todo-content p {
  margin: 0;
  color: #666;
  line-height: 1.4;
  transition: all 0.3s ease;
}

.todo-content p.completed {
  text-decoration: line-through;
  color: #aaa;
}

.todo-actions {
  display: flex;
  gap: 0.5rem;
  justify-content: flex-end;
}

.btn {
  padding: 0.5rem 1rem;
  border: none;
  border-radius: 6px;
  font-size: 0.9rem;
  cursor: pointer;
  transition: all 0.2s ease;
  min-width: 40px;
}

.btn-success {
  background: #10b981;
  color: white;
}

.btn-success:hover {
  background: #059669;
}

.btn-info {
  background: #3b82f6;
  color: white;
}

.btn-info:hover {
  background: #2563eb;
}

.btn-danger {
  background: #ef4444;
  color: white;
}

.btn-danger:hover {
  background: #dc2626;
}

.btn-primary {
  background: #8b5cf6;
  color: white;
}

.btn-primary:hover {
  background: #7c3aed;
}

.btn-secondary {
  background: #6b7280;
  color: white;
}

.btn-secondary:hover {
  background: #4b5563;
}

.edit-form {
  margin-top: 1rem;
  padding-top: 1rem;
  border-top: 1px solid #e5e7eb;
}

.form-group {
  margin-bottom: 1rem;
}

.form-input {
  width: 100%;
  padding: 0.5rem;
  border: 1px solid #d1d5db;
  border-radius: 6px;
  font-size: 0.9rem;
}

.form-input:focus {
  outline: none;
  border-color: #8b5cf6;
  box-shadow: 0 0 0 2px rgba(139, 92, 246, 0.2);
}

.form-actions {
  display: flex;
  gap: 0.5rem;
  justify-content: flex-end;
}

@media (max-width: 768px) {
  .todo-actions {
    justify-content: center;
  }
  
  .form-actions {
    justify-content: center;
  }
  
  .btn {
    font-size: 0.8rem;
    padding: 0.4rem 0.8rem;
  }
}
</style>

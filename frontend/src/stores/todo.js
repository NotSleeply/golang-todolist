import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { todoApi } from '@/api/todoApi'

export const useTodoStore = defineStore('todo', () => {
  // 状态
  const todos = ref([])
  const loading = ref(false)
  const error = ref(null)

  // 计算属性
  const completedTodos = computed(() => 
    todos.value.filter(todo => todo.completed)
  )
  
  const pendingTodos = computed(() => 
    todos.value.filter(todo => !todo.completed)
  )
  
  const totalTodos = computed(() => todos.value.length)

  // 操作方法
  const fetchTodos = async () => {
    try {
      loading.value = true
      error.value = null
      const response = await todoApi.getAll()
      todos.value = response.data || []
    } catch (err) {
      error.value = '获取任务失败'
      console.error('获取任务失败:', err)
    } finally {
      loading.value = false
    }
  }

  const createTodo = async (todoData) => {
    try {
      loading.value = true
      await todoApi.create(todoData)
      await fetchTodos() // 重新获取列表
    } catch (err) {
      error.value = '创建任务失败'
      throw err
    } finally {
      loading.value = false
    }
  }

  const updateTodo = async (todoData) => {
    try {
      loading.value = true
      await todoApi.update(todoData)
      await fetchTodos() // 重新获取列表
    } catch (err) {
      error.value = '更新任务失败'
      throw err
    } finally {
      loading.value = false
    }
  }

  const deleteTodo = async (id) => {
    try {
      loading.value = true
      await todoApi.delete(id)
      await fetchTodos() // 重新获取列表
    } catch (err) {
      error.value = '删除任务失败'
      throw err
    } finally {
      loading.value = false
    }
  }

  const clearError = () => {
    error.value = null
  }

  return {
    // 状态
    todos,
    loading,
    error,
    // 计算属性
    completedTodos,
    pendingTodos,
    totalTodos,
    // 方法
    fetchTodos,
    createTodo,
    updateTodo,
    deleteTodo,
    clearError
  }
})

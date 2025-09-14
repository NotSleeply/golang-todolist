import axios, { type AxiosResponse } from 'axios'
import type { Todo, CreateTodoRequest, UpdateTodoRequest, ApiResponse } from '@/types/todo'

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8001'

// 创建 axios 实例
const apiClient = axios.create({
  baseURL: `${API_BASE_URL}/api/v1`,
  headers: {
    'Content-Type': 'application/json'
  },
  timeout: 10000
})

// 请求拦截器
apiClient.interceptors.request.use(
  (config) => {
    console.log('发送请求:', config.method?.toUpperCase(), config.url)
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// 响应拦截器
apiClient.interceptors.response.use(
  (response: AxiosResponse) => {
    console.log('收到响应:', response.status, response.data)
    return response
  },
  (error) => {
    console.error('请求失败:', error.response?.data || error.message)
    return Promise.reject(error)
  }
)

export const todoApi = {
  // 获取所有任务
  getAll: (): Promise<AxiosResponse<Todo[]>> => apiClient.get('/todos'),
  
  // 创建任务
  create: (todo: CreateTodoRequest): Promise<AxiosResponse<ApiResponse<Todo>>> => 
    apiClient.post('/todos', todo),
  
  // 更新任务
  update: (todo: UpdateTodoRequest): Promise<AxiosResponse<ApiResponse<string>>> => 
    apiClient.put('/todos', todo),
  
  // 删除任务
  delete: (id: string): Promise<AxiosResponse<ApiResponse<string>>> => 
    apiClient.delete('/todos', { data: { id } })
}

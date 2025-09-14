export interface Todo {
  id: string
  name: string
  description: string
  completed: boolean
}

export interface CreateTodoRequest {
  name: string
  description: string
}

export interface UpdateTodoRequest {
  id: string
  name: string
  description: string
  completed: boolean
}

export interface ApiResponse<T> {
  data: T
  message?: string
}

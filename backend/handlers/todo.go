package handlers

import (
	"net/http"
	"todo/db"
	"todo/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TodoHandler struct{}

func NewTodoHandler() *TodoHandler {
	return &TodoHandler{}
}

/*
* CreateTodo 处理创建新的 Todo 项目的请求。
 */
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var params struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTodo := &model.Todo{
		ID:          uuid.New().String(),
		Name:        &params.Name,
		Description: &params.Description,
		Completed:   new(bool), // 默认 false
	}

	if err := db.CreateTodo(newTodo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo created successfully", "todo": newTodo})
}

/*
* GetAllTodos 处理获取所有 Todo 项目的请求。
 */
func (h *TodoHandler) GetAllTodos(c *gin.Context) {
	todos, err := db.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}

/*
* UpdateTodo 处理更新现有 Todo 项目的请求。
 */
func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	var params struct {
		ID          string `json:"id" binding:"required"`
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
		Completed   bool   `json:"completed"`
	}

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := &model.Todo{
		ID:          params.ID,
		Name:        &params.Name,
		Description: &params.Description,
		Completed:   &params.Completed,
	}

	if err := db.UpdateTodo(todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully"})
}

/*
* DeleteTodo 处理删除指定 Todo 项目的请求。
 */
func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	var params struct {
		ID string `json:"id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DeleteTodo(params.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}

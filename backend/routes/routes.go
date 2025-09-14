package routes

import (
	"net/http"
	"todo/handlers"
	"todo/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// 使用中间件
	r.Use(middleware.CORS())

	// 创建处理器实例
	todoHandler := handlers.NewTodoHandler()

	// 根路径
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "欢迎访问 TodoList 服务")
	})

	// API 路由组
	api := r.Group("/api/v1")
	{
		api.POST("/todos", todoHandler.CreateTodo)
		api.GET("/todos", todoHandler.GetAllTodos)
		api.PUT("/todos", todoHandler.UpdateTodo)
		api.DELETE("/todos", todoHandler.DeleteTodo)
	}

	// 保持兼容旧的路由
	r.POST("/create", todoHandler.CreateTodo)
	r.GET("/get-all-todos", todoHandler.GetAllTodos)
	r.PUT("/update", todoHandler.UpdateTodo)
	r.DELETE("/delete", todoHandler.DeleteTodo)

	return r
}

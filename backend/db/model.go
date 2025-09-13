package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Todo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func Init() {
	// 修改为你的 MySQL 连接信息
	dsn := "root:mYdatabase@tcp(127.0.0.1:3305)/my_database?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&Todo{})
}

func CreateTodo(todo Todo) error {
	result := db.Create(&todo)
	return result.Error
}

func GetAllTodos() ([]Todo, error) {
	var todos []Todo
	result := db.Find(&todos)
	return todos, result.Error
}

func UpdateTodo(todo Todo) error {
	result := db.Save(&todo)
	return result.Error
}

func DeleteTodo(id string) error {
	result := db.Delete(&Todo{}, "id = ?", id)
	return result.Error
}

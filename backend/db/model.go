package db

import (
	"todo/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Todo struct {
	ID          string `json:"id" gorm:"primarykey"`
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description"`
	Completed   bool   `json:"completed" gorm:"default:false"`
}

func Init() {
	cfg := config.Load()

	var err error
	db, err = gorm.Open(mysql.Open(cfg.GetDSN()), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败: " + err.Error())
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

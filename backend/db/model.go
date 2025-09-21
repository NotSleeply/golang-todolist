package db

import (
	"todo/config"
	"todo/model"
	"todo/query"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var Q *query.Query

/*
* Init 初始化数据库连接并进行自动迁移。
 */
func Init() {
	cfg := config.Load()

	var err error
	db, err = gorm.Open(mysql.Open(cfg.GetDSN()), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败: " + err.Error())
	}

	// 迁移 schema
	db.AutoMigrate(&model.Todo{})

	// 初始化 Gen 查询器
	Q = query.Use(db)
}

// 创建
func CreateTodo(todo *model.Todo) error {
	return Q.Todo.Create(todo)
}

// 查询全部
func GetAllTodos() ([]*model.Todo, error) {
	return Q.Todo.Find()
}

// 根据ID查询
func GetTodoByID(id string) (*model.Todo, error) {
	return Q.Todo.Where(Q.Todo.ID.Eq(id)).First()
}

// 更新
func UpdateTodo(todo *model.Todo) error {
	_, err := Q.Todo.Where(Q.Todo.ID.Eq(todo.ID)).Updates(todo)
	return err
}

// 删除
func DeleteTodo(id string) error {
	_, err := Q.Todo.Where(Q.Todo.ID.Eq(id)).Delete()
	return err
}

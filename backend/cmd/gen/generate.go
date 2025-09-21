package main

import (
	"todo/config"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	// 连接数据库
	cfg := config.Load()
	gormDB, err := gorm.Open(mysql.Open(cfg.GetDSN()), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败: " + err.Error())
	}

	// 创建 gen 实例
	g := gen.NewGenerator(gen.Config{
		OutPath:       "query",
		Mode:          gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	// 设置数据库连接
	g.UseDB(gormDB)

	// 生成所有表的代码
	tables := g.GenerateAllTable()
	g.ApplyBasic(tables...)

	// 生成代码
	g.Execute()
}

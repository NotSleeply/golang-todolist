package config

import "os"

type Config struct {
	Port       string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

/*
* Load 读取环境变量并返回项目配置。
* 用于初始化数据库和服务端口等参数。
 */
func Load() *Config {
	return &Config{
		Port:       getEnv("PORT", "8001"),
		DBHost:     getEnv("DB_HOST", "127.0.0.1"),
		DBPort:     getEnv("DB_PORT", "3305"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "mYdatabase"),
		DBName:     getEnv("DB_NAME", "my_database"),
	}
}

/*
* getEnv 获取环境变量的值，如果未设置则返回默认值。
 */
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

/*
* GetDSN 构建并返回数据库连接字符串。
 */
func (c *Config) GetDSN() string {
	return c.DBUser + ":" + c.DBPassword + "@tcp(" + c.DBHost + ":" + c.DBPort + ")/" + c.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
}

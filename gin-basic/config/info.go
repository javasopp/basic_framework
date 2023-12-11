package config

type ServerConfig struct {
	Port        int    `json:"port"`
	ContextPath string `json:"context-path"`
}

// MySQLConfig 结构体用于存储MySQL配置
type MySQLConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

// RedisConfig 结构体用于存储Redis配置
type RedisConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	Database int    `json:"database"`
}

// AppConfig 结构体用于存储整体应用程序配置
type AppConfig struct {
	Server ServerConfig `json:"server"`
	MySQL  MySQLConfig  `json:"mysql"`
	Redis  RedisConfig  `json:"redis"`
}

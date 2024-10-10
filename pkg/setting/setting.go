package setting

type Config struct {
	LogConfig LogConfig `json:"logger"`
}

type LogConfig struct {
	Level string `json:"level" binding:"required"`
	FileLogName string `json:"file_log_name" binding:"required"`
	MaxSize int `json:"max_size" binding:"required"`
	MaxBackups int `json:"max_backups" binding:"required"`
	MaxAge int `json:"max_age" binding:"required"`
	Compress bool `json:"compress" binding:"required"`
}
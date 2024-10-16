package setting

import "time"

type Config struct {
	LogConfig LogConfig `mapstructure:"logger"`
	DatabaseConfig DataBaseConfig `mapstructure:"database"`
	ServerConfig ServerConfig `mapstructure:"server"`
	TokenConfig TokenConfig `mapstructure:"token"`
}

type TokenConfig struct {
	TokenSymmetricKey string `mapstructure:"token_symmetric_key"`
	AccessTokenDuration time.Duration `mapstructure:"access_token_duration"`
}

type ServerConfig struct {
	Port int `mapstructure:"port"`
	Address string `mapstructure:"address"`
	Mode string `mapstructure:"mode"`
}

type LogConfig struct {
	Level string `mapstructure:"log_level"`
	FileLogName string `mapstructure:"file_log_name"`
	MaxSize int `mapstructure:"max_size"`
	MaxBackups int `mapstructure:"max_backups"`
	MaxAge int `mapstructure:"max_age"`
	Compress bool `mapstructure:"compress"`

}

type DataBaseConfig struct {
	Driver string `mapstructure:"driver"`
	UserName string `mapstructure:"user_name"`
	Password string `mapstructure:"password"`
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	DatabaseName string `mapstructure:"database_name"`
	MaxIdleConns int `mapstructure:"max_idle_conns"`
	MaxOpenConns int `mapstructure:"max_open_conns"`
	ConnMaxLifetime int `mapstructure:"conn_max_lifetime"`
}
